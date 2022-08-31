package main

import (
	"github.com/vpris/test-jwt/initializers"
	routes "github.com/vpris/test-jwt/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := routes.RoutesInit()
	router.Run()
}
