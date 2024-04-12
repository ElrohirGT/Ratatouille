package waitress

import (
	"fmt"

	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	tea "github.com/charmbracelet/bubbletea"
)

type WaitressModel struct {
	Menu components.MenuModel
}

func (m WaitressModel) Init() tea.Cmd {
	return nil
}

func (m WaitressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch newMsg := msg.(type) {
	case tea.KeyMsg:

		if newMsg.Type == tea.KeyEscape {
			return m, tea.Quit
		}

		if newMsg.Type == tea.KeyEnter {
			switch m.Menu.SelectedItem.Route {
			case "createClient":
				return CreateCreateClientView(), nil
			case "openAccount":
				return CreateOpenAccountView(), nil
			case "takeOrder":
				return CreateTakeOrder(), nil
			case "getActiveAccounts":
				return CreateGetActiveAccounts(), handleGetActiveAccounts()
			case "getClients":
				return CreateGetClientsView(), handleGetClients()
			case "getMenuItems":
				return CreateGetMenuItems(), handleGetMenuItems()
			case "generateBills":
				return m, nil
			case "registerComplain":
				return CreateRegisterComplain(), nil
			}
		}
		newMenu, cmd := m.Menu.Update(msg)
		m.Menu = newMenu.(components.MenuModel)

		return m, cmd
	}
	return m, nil
}
func (m WaitressModel) View() string {
	return m.Menu.View() + "\n" + fmt.Sprintf("%d %s", global.Id, global.Role)
}
