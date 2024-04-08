package analyst

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	MostOrdersView sessionState = iota + 1
	RushHourView
	AverageTimeToEat
)

type AnalystModel struct {
	state                   sessionState
	mostFamousDishesModel          MostFamousDishesModel
	getRushHourView         GetRushHourModel
	getAverageTimeToEatView GetAverageTimeToEatModel

	menu components.MenuModel
}

func InitialModel() tea.Model {

	items := []components.MenuItem{
		{ItemTitle: "Most ordered dishes", ItemDescription: "Between a date range"},
		{ItemTitle: "Schedule in which there are more orders", ItemDescription: "Betweenn a date range"},
		{ItemTitle: "Average eat time", ItemDescription: "Depending group of people"},
		{ItemTitle: "Complains grouped by person", ItemDescription: "Between a data range"},
		{ItemTitle: "Complains grouped by dish", ItemDescription: "Between a data range"},
		{ItemTitle: "Waitress eficiency", ItemDescription: "Grouped by person, in the last 6 months"},
	}

	menu := components.InitialModel("Menu Principal", items)

	return AnalystModel{0, MostFamousDishesModel{}, GetRushHourModel{}, GetAverageTimeToEatModel{}, menu}
}

func (m AnalystModel) Init() tea.Cmd {
	return nil
}

func (m AnalystModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch m.state {
	case 0:
		newMenu, cmd := m.menu.Update(msg)
		return newMenu.(components.MenuModel), cmd
	case 1: 
		// return m.mostFamousDishesModel.Update(msg)
		return m, nil
	}

	return m, nil
}

func (m AnalystModel) View() string {
	return m.menu.View()
}
