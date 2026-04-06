package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/12awoodward/gator/internal/config"
	"github.com/12awoodward/gator/internal/database"
	"github.com/google/uuid"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	allCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.allCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	runCmd, ok := c.allCommands[cmd.name]
	if !ok {
		return errors.New("command doesn't exist")
	}

	if err := runCmd(s, cmd); err != nil {
		return err
	}
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no username given")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.UserName); err != nil {
		return err
	}

	fmt.Printf("User set to: %s\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("name must be provided")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserName: cmd.args[0],
	})
	if err != nil {
		return err
	}

	s.cfg.SetUser(cmd.args[0])

	fmt.Printf("user was created.\n%v\n", user)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		userText := "* " + user.UserName

		if user.UserName == s.cfg.CurrentUserName {
			userText += " (current)"
		}
		fmt.Println(userText)
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("users reset")
	return nil
}