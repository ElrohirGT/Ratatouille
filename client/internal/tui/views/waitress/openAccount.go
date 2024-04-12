package waitress

import (
	"context"
	"strconv"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type OpenAccountView struct {
	forms       components.FormsModel
	numTable        string
	numPeople string
	errorMsg    string
}

func (m OpenAccountView) Init() tea.Cmd {
	return nil
}

func (m OpenAccountView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateWaitressView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			m.numTable = m.forms.FormInputs["Mesa"].Value
			m.numPeople = m.forms.FormInputs["numPersonas"].Value

			return m, handleOpenAccount(m.numTable, m.numPeople)
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

func (m OpenAccountView) View() string {
	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleOpenAccount(table, numPeople string) tea.Cmd {
	if table == "" || numPeople == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}
	
	tableConverted, err := strconv.Atoi(table)
	numPeopleConverted, err2 := strconv.Atoi(numPeople)

	if err != nil || err2 != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Table and numPeople must be integers!"}
		}
	}

	return func() tea.Msg {
		_, err := global.Driver.OpenAccount(context.Background(),
			db.OpenAccountParams(db.OpenAccountParams{Mesa: int32(tableConverted), Numpersonas: int32(numPeopleConverted)}))
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "Open account correctly"}
		}
	}

}
