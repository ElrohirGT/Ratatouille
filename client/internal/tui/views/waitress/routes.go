package waitress

import "github.com/ElrohirGT/Ratatouille/internal/tui/components"

func CreateWaitressView() WaitressModel {
	menuItems := []components.MenuItem{
		{Route: "createClient", ItemTitle: "Create Client", ItemDescription: "Registers a new client to the records."},
		{Route: "openAccount", ItemTitle: "Open Account", ItemDescription: "Create and account for a client to start order."},
		{Route: "takeOrder", ItemTitle: "Take Order", ItemDescription: "Ask the chef for a new food for a table"},
		{Route: "getActiveAccounts", ItemTitle: "Show Active accounts", ItemDescription: "Show accounts that are ordering dishes"},
		{Route: "getClients", ItemTitle: "Show Clients", ItemDescription: "Show all existing clients"},
		{Route: "generateBill", ItemTitle: "Generate Bill", ItemDescription: "Generate the Bill, close the account and pass the survey"},
		{Route: "registerComplain", ItemTitle: "Register Complain", ItemDescription: "Register a complain from a client"},
	}

	menu := components.CreateMenu("Waitress Management", menuItems)
	return WaitressModel{Menu: menu}
}

func CreateCreateClientView() CreateClientView {
	newForm := components.CreateForms("Create Client", map[string]components.FormsInput{
		"Name":    {Placeholder: "Hector Hurtarte"},
		"NIT":     {Placeholder: "XXXXXXXXX"},
		"Address": {Placeholder: "Guatemala city"},
	})
	return CreateClientView{forms: newForm}
}

func CreateOpenAccountView() OpenAccountView{
	newForm := components.CreateForms("Open Account", map[string]components.FormsInput{
		"Mesa":    {Placeholder: "1,2,3"},
		"numPersonas":     {Placeholder: "Between 6 an 14"},
	})
	return OpenAccountView{forms: newForm}
}