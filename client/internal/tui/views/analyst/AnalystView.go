package analyst

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/golang-collections/collections/stack"
)

var ViewStack *stack.Stack = stack.New()

type sessionState string

// SUBVIEWS
const (
	Main                 sessionState = "/"
	MostOrdersView       sessionState = "/most_orders"
	RushHourView         sessionState = "/rush_hour"
	AverageTimeToEatView sessionState = "/average_time_to_eat"
)

type AnalystModel struct {
	state                sessionState
	FamousDishesView     MostFamousDishesModel
	RushHourView         GetRushHourModel
	AverageTimeToEatView AverageTimeToEatModel

	menu components.MenuModel
}

func InitialModel() tea.Model {

	ViewStack.Push(string(Main))

	items := []components.MenuItem{
		{Route: string(MostOrdersView), ItemTitle: "Most ordered dishes", ItemDescription: "Between a date range"},
		{Route: string(RushHourView), ItemTitle: "Schedule in which there are more orders", ItemDescription: "Betweenn a date range"},
		{Route: string(AverageTimeToEatView), ItemTitle: "Average eat time", ItemDescription: "Depending group of people"},
		{Route: "/", ItemTitle: "Complains grouped by person", ItemDescription: "Between a data range"},
		{Route: "/", ItemTitle: "Complains grouped by dish", ItemDescription: "Between a data range"},
		{Route: "/", ItemTitle: "Waitress eficiency", ItemDescription: "Grouped by person, in the last 6 months"},
	}

	menu := components.CreateMenu("Menu Principal", items)

	return AnalystModel{Main, MostFamousDishesModel{}, GetRushHourModel{}, AverageTimeToEatModel{}, menu}
}

func (m AnalystModel) Init() tea.Cmd {
	return nil
}

func (m AnalystModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch currentState := sessionState(ViewStack.Peek().(string)); currentState {
	case Main:
		var cmd tea.Cmd
		newMenu, cmd := m.menu.Update(msg)
		m.menu = newMenu.(components.MenuModel)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				ViewStack.Push(m.menu.FocusItem.Route)
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
				ViewStack.Pop()
			}
		}
		return m, nil
	case RushHourView:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEsc {
				ViewStack.Pop()
			}
		}
		return m, nil
	case AverageTimeToEatView:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEsc {
				ViewStack.Pop()
			}
		
		}
		return m, nil
	}

	return m, nil
}

func (m AnalystModel) View() string {

	switch currentState := sessionState(ViewStack.Peek().(string)); currentState {
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
