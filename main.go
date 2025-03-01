package main

import (
	"fmt"
	"os"

	"github.com/cygran/gator/internal/cli"
	"github.com/cygran/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Failed to read config file: %v\n", err)
		os.Exit(1)
	}
	state := &cli.State{
		Config: &cfg,
	}
	cmds := &cli.Commands{
		Handlers: make(map[string]func(*cli.State, cli.Command) error),
	}
	cmds.Register("login", cli.HandlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arguments")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := cli.Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	if err := cmds.Run(state, cmd); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}
