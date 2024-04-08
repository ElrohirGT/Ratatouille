package main

import (
	"fmt"
	"os"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/views/analyst"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	// Driver initialization

	dbConnection, err := db.NewDriver("root", "root", "ratatouille")

	if err != nil {
		println("Something wrong with the Database...")
		return
	}
	driver := db.New(dbConnection)
	
	println(driver)

	m := analyst.InitialModel()

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program", err)
		os.Exit(1)
	}
}
