package main

import (
	"fmt"

	"anon-chat/config"
	"anon-chat/internal/app"
)

func main() {
	cfg, err := config.NewConfig(".")
	if err != nil {
		fmt.Printf("Config error: %s", err)
	}
	app.Run(cfg)
}
