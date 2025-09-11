package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/d-darac/gator/internal/config"
	"github.com/d-darac/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("error opening database connection: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	programState := &state{
		cfg: &cfg,
		db:  database.New(db),
	}

	cmds := commands{
		registry: make(map[string]func(*state, command) error),
	}

	cmds.register("addfeed", handlerAddFeed)
	cmds.register("agg", handlerAggregate)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 {
		fmt.Println("usage: cli <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	if err := cmds.run(programState, command{Name: cmdName, Args: cmdArgs}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
