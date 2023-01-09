package main

import (
	"golang_API/config"
	"golang_API/routes"
)

func main() {
	db := config.ConnectToDB()
	defer db.Close()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
