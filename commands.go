package main

import (
	"fmt"

	"gator/internal/config"
	"gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	Options map[string]func(*state, command) error
}


// Run a given command with the state
func (c *commands) run(s *state, cmd command) error {
	if (c.Options[cmd.Name] == nil) {
		return fmt.Errorf("Given command [%s] has not been registered", cmd.Name)
	}

	err := c.Options[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

// Register a new command
func (c *commands) register(name string, f func(*state, command) error) {
	c.Options[name] = f
}


