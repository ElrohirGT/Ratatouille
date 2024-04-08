package main

import (
	"fmt"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui"
)

var DB_Driver *db.Queries

func main() {

	// Connection Initialization
	dbConnection, err := db.NewDriver("backend", "backend", "ratatouille", "127.0.0.1", "5566")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Driver Initialization
	DB_Driver = db.New(dbConnection)

	tui.StartApp("encargado")
}
