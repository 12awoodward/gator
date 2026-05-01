package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/12awoodward/gator/internal/config"
	"github.com/12awoodward/gator/internal/database"
	_ "github.com/lib/pq"
)



func main() {
	s, err := initialStateSetup()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cmds := commands{allCommands: map[string]func(*state, command) error{}}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", handlerFollowing)
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	cmds.register("agg", handlerAgg)

	args := []string{}

	if len(os.Args) < 2 {
		fmt.Println("no command given")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	cmd := command{name: os.Args[1], args: args}

	if err := cmds.run(s, cmd); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func initialStateSetup() (*state, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		return nil, err
	}
	
	s := state{
		cfg: &cfg,
		db: database.New(db),
	}
	return &s, nil
}