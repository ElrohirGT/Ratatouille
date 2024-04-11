package waitress

import "github.com/ElrohirGT/Ratatouille/internal/tui/components"

func CreateWaitressView() WaitressModel {
	menuItems := []components.MenuItem{
		{Route: "createClient", ItemTitle: "Create Client", ItemDescription: "Registers a new client to the records."},
	}

	menu := components.CreateMenu("Waitress Management", menuItems)
	return WaitressModel{Menu: menu}
}

func CreateCreateClientView() CreateClientView {
	newForm := components.CreateForms("Create Client", map[string]components.FormsInput{
		"Name":    {Placeholder: "Hector Hurtarte"},
		"NIT":     {Placeholder: "XXXXXXXXX", InputType: "password"},
		"Address": {Placeholder: "Guatemala city"},
	})
	return CreateClientView{forms: newForm}
}
