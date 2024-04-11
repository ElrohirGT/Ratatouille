package auth

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type LoginModel struct {
	forms    components.FormsModel
	username string
	password string
}

func (m LoginModel) Init() tea.Cmd {
	return nil
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newForm, cmds := m.forms.Update(msg)
	m.forms = newForm.(components.FormsModel)

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateAuthView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			m.username = m.forms.FormInputs["Username"].Value
			m.password = m.forms.FormInputs["Password"].Value
			return m, cmds
		}
	}

	return m, cmds
}

func (m LoginModel) View() string {

	return m.forms.View()
}
