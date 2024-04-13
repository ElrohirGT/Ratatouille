package auth

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
)

func CreateAuthView() AuthModel {
	menuItems := []components.MenuItem{
		{Route: "Login", ItemTitle: "Login", ItemDescription: "Enter the app"},
		{Route: "SignUp", ItemTitle: "SignUp", ItemDescription: "Create a new user"},
	}

	menu := components.CreateMenu("Ratatouille", menuItems)
	return AuthModel{Menu: menu}
}

func CreateSignUpView() SignUpModel {
	newForm := components.CreateForms("Sign Up", map[string]components.FormsInput{
		"Username":   {Placeholder: "Hector Hurtarte"},
		"Password":   {Placeholder: "password", InputType: "password"},
		"Role":       {Placeholder: "1 (Mesero), 2 (Chef), 3 (Bartender), 4 (Encargado)"},
		"EmployeeID": {Placeholder: "XXX"},
	})
	return SignUpModel{forms: newForm}
}

func CreateLoginView() LoginModel {
	newForm := components.CreateForms("Log In", map[string]components.FormsInput{
		"Username": {Placeholder: "Hector Hurtarte"},
		"Password": {Placeholder: "password", InputType: "password"},
	})
	return LoginModel{forms: newForm}
}
