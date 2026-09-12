package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"

	"gator/internal/database"
)

// Login the user 
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Username required when logging in")
	}

	user, err := s.db.GetOneUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error fetching user: %s", err)
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("User '%s' has been logged in\n", user.Name)

	return nil
}

// Register a new user then login
func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Provide a username to register")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	})

	if err != nil {
		return fmt.Errorf("Some error... username probably already exists: %s", err)
	}

	fmt.Printf("User %s has been registered\n", user.Name)
	handlerLogin(s, cmd)

	return nil
}

// Delete all users from the database
func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Did not delete all users: %s", err)
	}
	fmt.Printf("Dropped all users from table\n")
	return nil
}

// Get all user names
func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetManyUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not fetch users: %s", err)
	}
	
	for _, user := range users {
		if (user.Name == s.cfg.CurrentUserName) {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
