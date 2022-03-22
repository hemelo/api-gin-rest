package main

import (
	"api-go-gin/config"
	"api-go-gin/database"
	"api-go-gin/routes"
)

func main() {
	config.Settings.TestMode = false
	database.Connect()
	routes.HandleRequests()
}
