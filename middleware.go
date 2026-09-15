package main

import (
	"fmt"
	"context"

	"gator/internal/database"
)

type LoggedInHandler func(s *state, cmd command, user database.User) error
type NormalHandler func(*state, command) error

func middlewareLoggedIn(handler LoggedInHandler) NormalHandler {
	return func(s *state, cmd command) error {
		user, err := s.db.GetOneUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user not found for name %s", s.cfg.CurrentUserName)
		}
		return handler(s, cmd, user)
	}
}

