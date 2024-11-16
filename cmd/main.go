package main

import (
	"fmt"

	"echo-box/config"
	"echo-box/internal/app"
)

func main() {
	cfg, err := config.NewConfig(".")
	if err != nil {
		fmt.Printf("Config error: %s", err)
	}
	app.Run(cfg)
}
