package main

import (
	"activities-backend/app"
	activityDatabase "activities-backend/database"
)

func main() {
	activityDatabase.StartDbEngine()
	app.StartRoute()
}
