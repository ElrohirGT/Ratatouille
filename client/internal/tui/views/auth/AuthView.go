package auth

import (
	"fmt"

	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	tea "github.com/charmbracelet/bubbletea"
)

type AuthModel struct {
	Menu components.MenuModel
}

func (m AuthModel) Init() tea.Cmd {
	return nil
}

func (m AuthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch newMsg := msg.(type) {
	case tea.KeyMsg:

		if newMsg.Type == tea.KeyEscape {
			return m, tea.Quit
		}

		if newMsg.Type == tea.KeyEnter {
			switch m.Menu.SelectedItem.Route {
			case "Login":
				return CreateLoginView(), nil
			case "SignUp":
				return CreateSignUpView(), nil
			}
		}

		newMenu, cmd := m.Menu.Update(msg)
		m.Menu = newMenu.(components.MenuModel)

		return m, cmd
	}

	return m, nil
}
func (m AuthModel) View() string {
	return m.Menu.View() + "\n" + fmt.Sprintf(global.Role)
}
