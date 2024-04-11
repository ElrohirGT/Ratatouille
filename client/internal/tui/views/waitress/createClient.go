package waitress

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type CreateClientView struct {
	forms   components.FormsModel
	name    string
	nit     string
	address string
}

func (m CreateClientView) Init() tea.Cmd {
	return nil
}

func (m CreateClientView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newForm, cmds := m.forms.Update(msg)
	m.forms = newForm.(components.FormsModel)

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateWaitressView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			m.name = m.forms.FormInputs["Username"].Value
			m.nit = m.forms.FormInputs["Password"].Value
			m.address = m.forms.FormInputs["Address"].Value
			return m, cmds
		}
	}

	return m, cmds
}

func (m CreateClientView) View() string {

	return m.forms.View()
}
