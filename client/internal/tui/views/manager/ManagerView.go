package manager

import (
	"fmt"

	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	tea "github.com/charmbracelet/bubbletea"
)

type ManagerModel struct {
	Menu components.MenuModel
}

func (m ManagerModel) Init() tea.Cmd {
	return nil
}

func (m ManagerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch newMsg := msg.(type) {
	case tea.KeyMsg:

		if newMsg.Type == tea.KeyEscape {
			return m, tea.Quit
		}

		if newMsg.Type == tea.KeyEnter {
			switch m.Menu.SelectedItem.Route {
			case "MostFamousDish":
				return CreateFamousDishView(), nil
			case "PeekHour":
				return m, nil
			case "AverageEatTime":
				return m, nil
			case "ComplaintsByPerson":
				return m, nil
			case "ComplaintsByDish":
				return m, nil
			case "WaitressEfficiency":
				return m, nil
			}
		}
		newMenu, cmd := m.Menu.Update(msg)
		m.Menu = newMenu.(components.MenuModel)

		return m, cmd
	}
	return m, nil
}
func (m ManagerModel) View() string {
	return m.Menu.View() + "\n" + fmt.Sprintf("%d %s", global.Id, global.Role)
}
