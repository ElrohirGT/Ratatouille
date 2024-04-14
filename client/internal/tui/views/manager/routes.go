package manager

import "github.com/ElrohirGT/Ratatouille/internal/tui/components"

func CreateManagerView() ManagerModel {
	menuItems := []components.MenuItem{
		{Route: "MostFamousDish", ItemTitle: "Most Famous Dishes", ItemDescription: "Between 2 dates"},
		{Route: "PeekHour", ItemTitle: "Peek Hour", ItemDescription: "Between 2 dates"},
		{Route: "AverageEatTime", ItemTitle: "Average Eat Time", ItemDescription: "Between 2 dates"},
		{Route: "ComplaintsByPerson", ItemTitle: "Complaints", ItemDescription: "By person"},
		{Route: "ComplaintsByDish", ItemTitle: "Complaints", ItemDescription: "By Dish"},
		{Route: "WaitressEfficiency", ItemTitle: "Waitress Eficiency", ItemDescription: "Based on surveys"},
	}

	menu := components.CreateMenu("Manager Management", menuItems)
	return ManagerModel{Menu: menu}
}

func CreateFamousDishView() famousDishesModel {
	newForm := components.CreateForms("Log In", map[string]components.FormsInput{
		"StartDate": {Placeholder: "YYYY-MM-DD"},
		"EndDate":   {Placeholder: "YYYY-MM-DD"},
	})
	return famousDishesModel{forms: newForm}
}
