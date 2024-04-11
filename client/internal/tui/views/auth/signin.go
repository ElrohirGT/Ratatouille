package auth

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type SignInModel struct {
	forms    components.FormsModel
	username string
	password string
	role     string
}

func (m SignInModel) Init() tea.Cmd {
	return nil
}

func (m SignInModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newForm, cmds := m.forms.Update(msg)
	m.forms = newForm.(components.FormsModel)

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateAuthView(), cmds
		}
		if newMsg.Type == tea.KeyEnter {
			m.username = m.forms.FormInputs["Username"].Value
			m.password = m.forms.FormInputs["Password"].Value
			m.role = m.forms.FormInputs["Role"].Value
			return m, nil
		}
	}

	return m, cmds
}

func (m SignInModel) View() string {

	return m.forms.View()
}
