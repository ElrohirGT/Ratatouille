package waitress

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/charmbracelet/bubbles/table"
)

func CreateWaitressView() WaitressModel {
	menuItems := []components.MenuItem{
		{Route: "createClient", ItemTitle: "Create Client", ItemDescription: "Registers a new client to the records."},
		{Route: "openAccount", ItemTitle: "Open Account", ItemDescription: "Create and account for a client to start order."},
		{Route: "takeOrder", ItemTitle: "Take Order", ItemDescription: "Ask the chef for a new food for a table"},
		{Route: "getActiveAccounts", ItemTitle: "Show Active accounts", ItemDescription: "Show accounts that are ordering dishes"},
		{Route: "getClients", ItemTitle: "Show Clients", ItemDescription: "Show all existing clients"},
		{Route: "getMenuItems", ItemTitle: "Show Menu", ItemDescription: "Show all existing dishes and drinks options."},
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

func CreateOpenAccountView() OpenAccountView {
	newForm := components.CreateForms("Open Account", map[string]components.FormsInput{
		"Mesa":        {Placeholder: "1,2,3"},
		"numPersonas": {Placeholder: "Less than 14"},
	})
	return OpenAccountView{forms: newForm}
}

func CreateTakeOrder() TakeOrderView {
	newForm := components.CreateForms("Take Order", map[string]components.FormsInput{
		"NO.account": {Placeholder: "XXX"},
		"ItemID":     {Placeholder: "XXX"},
		"Amount":     {Placeholder: "XXX"},
	})
	return TakeOrderView{forms: newForm}
}

func CreateRegisterComplain() RegisterComplain {
	newForm := components.CreateForms("Register Complain", map[string]components.FormsInput{
		"Client":   {Placeholder: "1,2,3"},
		"Severity": {Placeholder: "1 to 5"},
		"Reason":   {Placeholder: "Because...", InputType: "text"},
		"Employee": {Placeholder: "1,2,3..."},
		"Item":     {Placeholder: "1,2,3..."},
	})
	return RegisterComplain{forms: newForm}
}

func CreateGenerateBill() GenerateBillView {
	newForm := components.CreateForms("Generate Bill", map[string]components.FormsInput{
		"Account": {Placeholder: "1,2,3..."},
		"Client":  {Placeholder: "1,2,3..."},
	})
	return GenerateBillView{forms: newForm}
}

func CreateTakeSurvey() takeSurveyModel {
	newForm := components.CreateForms("Take Survey", map[string]components.FormsInput{
		"Employee": {Placeholder: "1,2,3..."},
		"Client":   {Placeholder: "1,2,3..."},
		"Kindness": {Placeholder: "1,2,3..."},
		"Speed":    {Placeholder: "1,2,3..."},
	})
	return takeSurveyModel{forms: newForm}
}

func CreateGetClientsView() getClientsViewModel {
	return getClientsViewModel{table: table.New()}
}

func CreateGetActiveAccounts() getActiveAccountsModel {
	return getActiveAccountsModel{table: table.New()}
}

func CreateGetMenuItems() getMenuItemsViewModel {
	return getMenuItemsViewModel{table: table.New()}
}
