package analyst

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/constants"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	Main sessionState = iota
	MostOrdersView
	RushHourView
	AverageTimeToEatView
)

type AnalystModel struct {
	state                   sessionState
	mostFamousDishesModel   MostFamousDishesModel
	getRushHourView         GetRushHourModel
	getAverageTimeToEatView GetAverageTimeToEatModel

	menu components.MenuModel
}

func InitialModel() tea.Model {

	constants.GetViewStack().Push(0)

	items := []components.MenuItem{
		{Index: 1, ItemTitle: "Most ordered dishes", ItemDescription: "Between a date range"},
		{Index: 2, ItemTitle: "Schedule in which there are more orders", ItemDescription: "Betweenn a date range"},
		{Index: 3, ItemTitle: "Average eat time", ItemDescription: "Depending group of people"},
		{Index: 4, ItemTitle: "Complains grouped by person", ItemDescription: "Between a data range"},
		{Index: 5, ItemTitle: "Complains grouped by dish", ItemDescription: "Between a data range"},
		{Index: 6, ItemTitle: "Waitress eficiency", ItemDescription: "Grouped by person, in the last 6 months"},
	}

	menu := components.InitialModel("Menu Principal", items)

	return AnalystModel{0, MostFamousDishesModel{}, GetRushHourModel{}, GetAverageTimeToEatModel{}, menu}
}

func (m AnalystModel) Init() tea.Cmd {
	return nil
}

func (m AnalystModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch currentState := sessionState(constants.ViewStack.Peek().(int)); currentState  {
	case Main:
		var cmd tea.Cmd
		newMenu, cmd := m.menu.Update(msg)
		m.menu = newMenu.(components.MenuModel)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				constants.GetViewStack().Push(m.menu.FocusItem.Index)
			}
			if msg.Type == tea.KeyEsc {
				return m, tea.Quit
			}
		}
		return m, cmd
	case MostOrdersView:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEsc {
				constants.GetViewStack().Pop()
			}
		}
		return m, nil
	case RushHourView:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEsc {
				constants.GetViewStack().Pop()
			}
		}
		return m, nil
	case AverageTimeToEatView:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEsc {
				constants.GetViewStack().Pop()
			}
		}
		return m, nil
	}

	return m, nil
}

func (m AnalystModel) View() string {

	switch currentState := sessionState(constants.ViewStack.Peek().(int)); currentState {
	case Main:
		return m.menu.View()
	case MostOrdersView:
		// return m.mostFamousDishesModel.Update(msg)
		return "Hello from Most Orders View"
	case RushHourView:
		// return m.mostFamousDishesModel.Update(msg)
		return "Hello from Rush Hour View"
	case AverageTimeToEatView:
		// return m.mostFamousDishesModel.Update(msg)
		return "Hello from Averate Time To Eat View"
	}

	return m.menu.View()
}
