package main

import (
	"pokedex/connections"
	"pokedex/repository"
)

func main() {
	// connect DB
	repository.Db = connections.DbConnection()
	// connect GIN
	connections.GinSetup()
}
