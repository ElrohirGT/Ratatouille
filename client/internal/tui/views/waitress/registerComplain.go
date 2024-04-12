package waitress

import (
	"context"
	"database/sql"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type RegisterComplain struct {
	forms    components.FormsModel
	name     string
	nit      string
	address  string
	errorMsg string
}

func (m RegisterComplain) Init() tea.Cmd {
	return nil
}

func (m RegisterComplain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateWaitressView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			m.name = m.forms.FormInputs["Name"].Value
			m.nit = m.forms.FormInputs["NIT"].Value
			m.address = m.forms.FormInputs["Address"].Value

			return m, handleCreateClient(m.name, m.nit, m.address)
		}

		newForm, cmds := m.forms.Update(msg)
		m.forms = newForm.(components.FormsModel)
		return m, cmds

	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		return CreateWaitressView(), nil
	}

	return m, nil
}

func (m RegisterComplain) View() string {
	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleRegisterComplain(name, NIT, address string) tea.Cmd {
	if name == "" || NIT == "" || address == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	return func() tea.Msg {
		_, err := global.Driver.CreateClient(context.Background(),
			db.CreateClientParams{Nombre: name, Nit: NIT, Direccion: sql.NullString{String: address, Valid: true}})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "Client created correctly"}
		}
	}

}
