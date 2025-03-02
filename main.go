package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cygran/gator/internal/cli"
	"github.com/cygran/gator/internal/config"
	"github.com/cygran/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Failed to read config file: %v\n", err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
	}
	dbQueries := database.New(db)
	state := &cli.State{
		Db:     dbQueries,
		Config: &cfg,
	}
	cmds := &cli.Commands{
		Handlers: make(map[string]func(*cli.State, cli.Command) error),
	}
	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)
	cmds.Register("reset", cli.HandlerResetUsers)
	cmds.Register("users", cli.HandlerUsers)
	cmds.Register("agg", cli.HandlerAgg)
	cmds.Register("addfeed", cli.MiddlewareLoggedIn(cli.HandlerAddFeed))
	cmds.Register("feeds", cli.HandlerFeeds)
	cmds.Register("follow", cli.MiddlewareLoggedIn(cli.HandlerFollow))
	cmds.Register("following", cli.MiddlewareLoggedIn(cli.HandlerFollowing))
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
