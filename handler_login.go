package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("couldn't set current user %w", err)
	}

	fmt.Printf("logged in as: %s\n", name)

	return nil
}
