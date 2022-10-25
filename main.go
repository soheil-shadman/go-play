package main

import (
	"myapp/config"
	"myapp/db"
	"myapp/routes"
)

func main() {

	//db connection
	db.Init()

	// Routes
	echoServer := routes.Init()
	echoServer.Logger.Fatal(echoServer.Start(":" + config.PORT_STRING()))

}
