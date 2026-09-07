package main

import (
	"os"
	"fmt"

	"gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		os.Exit(1)
	}

	cfg.SetUser("matthew")

	cfg, err = config.Read()
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("%+v\n", cfg)
}
