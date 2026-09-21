package main

import (
	"fmt"
	"context"
	"strconv"

	"gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.Args) != 0 {
		parseLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("could not parse string to int: %v", err)
		}
		limit = int32(parseLimit)
	} 

	posts, err := s.db.GetPostsForUser(
		context.Background(), 
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit: limit,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to get posts: %v", err)
	}

	fmt.Println("Posts: ------------")
	for _, post := range posts {
		fmt.Printf("\t%s - %s\n", post.Title, post.Url)
		// if post.Description.Valid {
		// 	fmt.Printf("\t\t%s\n", post.Description.String)
		// }
	}

	return nil
}

