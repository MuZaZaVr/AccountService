package main

import "github.com/MuZaZaVr/account-service/internal/app"

const configPath = "config/main"

func main() {
	app.Run(configPath)
}