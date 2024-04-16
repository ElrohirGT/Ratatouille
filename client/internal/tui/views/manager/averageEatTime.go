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

type averageEatTimeModel struct {
	forms     components.FormsModel
	data      []db.GetAverageTimeToEatPerClientQuantityRow
	startDate string
	endDate   string
	errorMsg  string
}

func (m averageEatTimeModel) Init() tea.Cmd {
	return nil
}

func (m averageEatTimeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
			return m, handleFethAverageEatTime(m.startDate, m.endDate)
		}
	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		response := newMsg.Value
		m.data = response.([]db.GetAverageTimeToEatPerClientQuantityRow)
		return m, nil
	}

	return m, cmds
}

func (m averageEatTimeModel) View() string {

	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if len(m.data) > 0 {
		b.WriteString(printAverageEatTimeData(m.data))
	}

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleFethAverageEatTime(startDate, endDate string) tea.Cmd {

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
		v, err := global.Driver.GetAverageTimeToEatPerClientQuantity(context.Background(),
			db.GetAverageTimeToEatPerClientQuantityParams{Fecha: date1, Fecha_2: date2})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "User created correctly", Value: v}
		}
	}
}

func printAverageEatTimeData(data []db.GetAverageTimeToEatPerClientQuantityRow) string {
	var b strings.Builder

	b.WriteString(
		fmt.Sprintf("%-15s| %-15s|\n", "Num People", "Avg Eat Time"))
	for _, v := range data {
		b.WriteString(
			fmt.Sprintf("%-15d| %-15f|\n", v.Numpersonas, v.Timetoeat))
	}

	return b.String()
}
