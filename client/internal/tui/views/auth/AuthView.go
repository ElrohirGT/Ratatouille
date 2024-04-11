package auth

import (
	"fmt"

	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type AuthModel struct {
	Menu components.MenuModel
}

func (m AuthModel) Init() tea.Cmd {
	return nil
}

func (m AuthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newMenu, cmd := m.Menu.Update(msg)
	m.Menu = newMenu.(components.MenuModel)

	switch newMsg := msg.(type) {
	case tea.KeyMsg:

		if newMsg.Type == tea.KeyEscape {
			return m, tea.Quit
		}

		if newMsg.Type == tea.KeyEnter {
			switch m.Menu.FocusItem.Route {
			case "Login":
				fmt.Println("YOU PRESS ME")
				return m, cmd
			case "SignUp":
				fmt.Println("YOU PRESS ME 2")
				return m, cmd
			}
		}
		return m, cmd
	}
	return m, nil
}
func (m AuthModel) View() string {
	return m.Menu.View()
}
