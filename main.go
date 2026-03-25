package main

import (
	"fmt"

	"github.com/12awoodward/gator/internal/config"
)



func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)

	cfg.SetUser("lane")

	newConfig, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newConfig)
}