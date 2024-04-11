package waitress

import (
	"context"
	"database/sql"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
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
			m.name = m.forms.FormInputs["Name"].Value
			m.nit = m.forms.FormInputs["NIT"].Value
			m.address = m.forms.FormInputs["Address"].Value
			
			return m, handleCreateClient(m.name, m.nit, m.address)
		}
	}

	return m, cmds
}

func (m CreateClientView) View() string {

	return m.forms.View()
}

func handleCreateClient(name, NIT, address string) tea.Cmd {
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
