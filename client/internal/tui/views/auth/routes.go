package auth

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type IController interface {
	ShowAuthView() tea.Model
}

func CreateAuthView() AuthModel {
	menuItems := []components.MenuItem{
		{Route: "Login", ItemTitle: "Login", ItemDescription: "Enter the app"},
		{Route: "SignUp", ItemTitle: "SignUp", ItemDescription: "Create a new user"},
	}

	menu := components.CreateMenu("Ratatouille", menuItems)
	return AuthModel{Menu: menu}
}
