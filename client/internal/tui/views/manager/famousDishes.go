package manager

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type famousDishesModel struct {
	forms     components.FormsModel
	data      []db.GetMostFamousDishesBetweenRow
	startDate string
	endDate   string
	errorMsg  string
}

func (m famousDishesModel) Init() tea.Cmd {
	return nil
}

func (m famousDishesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newForm, cmds := m.forms.Update(msg)
	m.forms = newForm.(components.FormsModel)

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateManagerView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			m.startDate = m.forms.FormInputs["StartDate"].Value
			m.endDate = m.forms.FormInputs["EndDate"].Value
			return m, handleFetchFamousDishes(m.startDate, m.endDate)
		}
	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		response := newMsg.Value
		m.data = response.([]db.GetMostFamousDishesBetweenRow)
		return m, nil
	}

	return m, cmds
}

func (m famousDishesModel) View() string {

	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if len(m.data) > 0 {
		b.WriteString(printData(m.data))
	}

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleFetchFamousDishes(startDate, endDate string) tea.Cmd {

	if startDate == "" || endDate == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	date1, err := global.ParseDate(startDate)
	date2, err2 := global.ParseDate(endDate)

	if err != nil || err2 != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Not valid date format"}
		}
	}

	return func() tea.Msg {
		v, err := global.Driver.GetMostFamousDishesBetween(context.Background(),
			db.GetMostFamousDishesBetweenParams{Fecha: date1, Fecha_2: date2})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "User created correctly", Value: v}
		}
	}
}

func printData(data []db.GetMostFamousDishesBetweenRow) string {
	var b strings.Builder

	b.WriteString(
		fmt.Sprintf("%-25s| %-8s| %s\n", "Name", "Count", "Description"))
	for _, v := range data {
		b.WriteString(
			fmt.Sprintf("%-25s| %-8d| %s\n", v.Nombre, v.Count, v.Descripcion))
	}

	return b.String()
}
