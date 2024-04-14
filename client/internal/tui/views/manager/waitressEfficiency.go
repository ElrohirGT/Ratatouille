package manager

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type waitressEfficiency struct {
	data      []db.GetEfficiencyReportRow
	errorMsg  string
}

func (m waitressEfficiency) Init() tea.Cmd {
	return nil
}

func (m waitressEfficiency) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateManagerView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			return m, handleFetchWaitressEfficienty()
		}
	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		response := newMsg.Value
		m.data = response.([]db.GetEfficiencyReportRow)
		return m, nil
	}

	return m, nil
}

func (m waitressEfficiency) View() string {

	var b strings.Builder

	if len(m.data) > 0 {
		b.WriteString(printWaitressEfficiency(m.data))
	}

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleFetchWaitressEfficienty() tea.Cmd {

	return func() tea.Msg {
		v, err := global.Driver.GetEfficiencyReport(context.Background())
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "User created correctly", Value: v}
		}
	}
}

func printWaitressEfficiency(data []db.GetEfficiencyReportRow) string {
	var b strings.Builder

	b.WriteString(
		fmt.Sprintf("%-10s| %-10s| %-10s| %-10s|\n",
			"Month",
			"Employee",
			"Avg",
			"Avg_2"))
	for _, v := range data {
		b.WriteString(
			fmt.Sprintf("%-10s| %-10d| %-10f| %-10f|\n",
				v.Mes,
				v.Empleado,
				v.Avg,
				v.Avg_2,
				))
	}

	return b.String()
}
