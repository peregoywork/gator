package main

import _ "github.com/lib/pq"
import (
	"os"
	"fmt"

	"database/sql"

	"gator/internal/config"
	"gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		os.Exit(1)
	}
	
	db, err := sql.Open("postgres", cfg.DbUrl)
		if err != nil {
		os.Exit(1)
	}
	
	dbQueries := database.New(db)
	var projState = state{ 
		db: dbQueries,
		cfg: &cfg, 
	}

	var cmdMap = commands{ 
		Options: map[string]func(*state, command) error{}, 
	}
	cmdMap.register("login", handlerLogin)
	cmdMap.register("register", handlerRegister)
	cmdMap.register("reset", handlerReset)
	cmdMap.register("users", handlerGetUsers)
	cmdMap.register("agg", handlerAgg)
	cmdMap.register("addfeed", handlerAddFeed)
	cmdMap.register("feeds", handlerGetFeeds)
	cmdMap.register("follow", handlerFollow)
	cmdMap.register("following", handlerListFollowing)

	args := os.Args
	if (len(args) < 2) {
		fmt.Println("Command name not given")
		os.Exit(1)
	}

	var cmd = command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmdMap.run(&projState, cmd)
	if err != nil {
		fmt.Printf("Error running command [%s] - %s\n", cmd.Name, err)
		os.Exit(1)
	}
}
