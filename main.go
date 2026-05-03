package main

import "fraud-analytics/internal/app"

func main() {
	application := app.NewApp()
	application.Run()
}
