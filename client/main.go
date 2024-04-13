package main

import (
	"fmt"
	"os"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/views/waitress"
	tea "github.com/charmbracelet/bubbletea"
)

var DB_Driver *db.Queries

func main() {

	// Connection Initialization
	dbConnection, err := db.NewDriver("backend", "backend", "ratatouille", "127.0.0.1", "5566")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dbConnection.Close()

	// Driver Initialization
	global.Driver = db.New(dbConnection)

	p := tea.NewProgram(waitress.CreatePayBillView(1, 100))
	if _, err := p.Run(); err != nil {
		fmt.Printf("An Error happened: %v", err)
		os.Exit(1)
	}

}
