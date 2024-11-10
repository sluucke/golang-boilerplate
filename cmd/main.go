package main

import (
	"github.com/sluucke/golang-boilerplate/config"
	"github.com/sluucke/golang-boilerplate/internal/app"
)

func main() {
	config := config.NewEnv()
	app.Run(config)
}
