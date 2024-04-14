package manager

import (
	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
)

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
	newForm := components.CreateForms("Famous Dishes", map[string]components.FormsInput{
		"StartDate": {Placeholder: "YYYY-MM-DD"},
		"EndDate":   {Placeholder: "YYYY-MM-DD"},
	})
	return famousDishesModel{forms: newForm}
}

func CreatePeekHourView() peekHourModel {
	newForm := components.CreateForms("Peek Hour", map[string]components.FormsInput{
		"StartDate": {Placeholder: "YYYY-MM-DD"},
		"EndDate":   {Placeholder: "YYYY-MM-DD"},
	})
	return peekHourModel{forms: newForm, data: db.GetRushHourBetweenRow{}}
}

func CreateAverageTimeToEatView() averageEatTimeModel {
	newForm := components.CreateForms("Average Time To Eat", map[string]components.FormsInput{
		"StartDate": {Placeholder: "YYYY-MM-DD"},
		"EndDate":   {Placeholder: "YYYY-MM-DD"},
	})
	return averageEatTimeModel{forms: newForm}
}

func CreateComplaintsByPersonView() complaintsByPersonModel {
	newForm := components.CreateForms("Complaints", map[string]components.FormsInput{
		"StartDate": {Placeholder: "YYYY-MM-DD"},
		"EndDate":   {Placeholder: "YYYY-MM-DD"},
		"Employee":  {Placeholder: "1,2,3..."},
	})
	return complaintsByPersonModel{forms: newForm}
}

func CreateComplaintsByItemView() complaintsByItemModel {
	newForm := components.CreateForms("Complaints", map[string]components.FormsInput{
		"StartDate": {Placeholder: "YYYY-MM-DD"},
		"EndDate":   {Placeholder: "YYYY-MM-DD"},
		"Item":   {Placeholder: "1,2,3..."},
	})
	return complaintsByItemModel{forms: newForm}
}

func CreateWaitressEfficiencyView() waitressEfficiency{
	return waitressEfficiency{}
}
