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

func CreateSignInView() SignInModel {
	newForm := components.CreateForms("Sign In", map[string]components.FormsInput{
		"Username": {Placeholder: "Hector Hurtarte"},
		"Password": {Placeholder: "password", InputType: "password"},
		"Role":     {Placeholder: "Mesero, Chef..."},
	})
	return SignInModel{forms: newForm}
}

func CreateLoginView() SignInModel {
	newForm := components.CreateForms("Sign In", map[string]components.FormsInput{
		"Username": {Placeholder: "Hector Hurtarte"},
		"Password": {Placeholder: "password", InputType: "password"},
	})
	return SignInModel{forms: newForm}
}
