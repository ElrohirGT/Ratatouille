package tui

import (
	"fmt"
	"os"

	"github.com/ElrohirGT/Ratatouille/internal/tui/views/auth"
	tea "github.com/charmbracelet/bubbletea"
)

func StartAuthentication() (username string, password string){

	program := tea.NewProgram(auth.InitialModel(), tea.WithAltScreen())

	finalModel, err := program.Run()
	
	if err != nil {
		fmt.Println("Error running program", err)
		os.Exit(1)
	}

	response := finalModel.(auth.AuthModel)
	
	return response.Username, response.Password
}

func StartApp() {

}
