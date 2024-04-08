package tui

import (
	"fmt"
	"os"

	"github.com/ElrohirGT/Ratatouille/internal/tui/views/analyst"
	"github.com/ElrohirGT/Ratatouille/internal/tui/views/auth"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: Add SQL Driver to paramater, and add it to the model
func StartAuthentication() (username string, password string, role string) {

	program := tea.NewProgram(auth.InitialModel(), tea.WithAltScreen())

	finalModel, err := program.Run()

	if err != nil {
		fmt.Println("Error running program", err)
		os.Exit(1)
	}

	response := finalModel.(auth.AuthModel)

	return response.Username, response.Password, "role"
}

func StartApp(role string) {

	var m tea.Model
	switch role {
	case "encargado":
		m = analyst.InitialModel()
	default:
		m = analyst.InitialModel()
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program", err)
		os.Exit(1)
	}
}
