package main

import (
	"fmt"

	"github.com/cygran/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Failed to read config file: %v\n", err)
	}
	cfg.SetUser("Cygran")
	if err != nil {
		fmt.Printf("Failed to complete setting user: %v\n", err)
		return
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Failed to read updated config file: %v\n", err)
		return
	}
	fmt.Printf("Config successfully updated: %+v\n", cfg)
}
