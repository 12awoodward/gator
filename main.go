package main

import (
	"fmt"
	"os"

	"github.com/12awoodward/gator/internal/config"
)



func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	s := state{cfg: &cfg}
	cmds := commands{allCommands: map[string]func(*state, command) error{}}

	cmds.register("login", handlerLogin)

	args := []string{}

	if len(os.Args) < 2 {
		fmt.Println("no command given")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	cmd := command{name: os.Args[1], args: args}

	if err := cmds.run(&s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}