package main

import (
	"fmt"
	"os"

	"github.com/d-darac/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v", err)
		os.Exit(1)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registry: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

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
