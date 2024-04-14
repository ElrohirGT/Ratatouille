package manager

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type complaintsByPersonModel struct {
	forms     components.FormsModel
	data      []db.Queja
	startDate string
	endDate   string
	employee  string
	errorMsg  string
}

func (m complaintsByPersonModel) Init() tea.Cmd {
	return nil
}

func (m complaintsByPersonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
			m.employee = m.forms.FormInputs["Employee"].Value
			return m, handleFetchComplainsByPerson(m.employee, m.startDate, m.endDate)
		}
	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		response := newMsg.Value
		m.data = response.([]db.Queja)
		return m, nil
	}

	return m, cmds
}

func (m complaintsByPersonModel) View() string {

	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if len(m.data) > 0 {
		b.WriteString(printComplaintsByPersonData(m.data))
	}

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleFetchComplainsByPerson(employee, startDate, endDate string) tea.Cmd {

	if startDate == "" || endDate == "" || employee == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	date1, err := global.ParseDate(startDate)
	date2, err2 := global.ParseDate(endDate)
	empleado, err3 := strconv.Atoi(employee)

	if err != nil || err2 != nil || err3 != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Not valid date format"}
		}
	}

	return func() tea.Msg {
		v, err := global.Driver.GetComplaintsForEmployeeBetween(context.Background(),
			db.GetComplaintsForEmployeeBetweenParams{
				Empleado: sql.NullInt32{Int32: int32(empleado), Valid: true},
				Fecha:    date1, Fecha_2: date2})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "User created correctly", Value: v}
		}
	}
}

func printComplaintsByPersonData(data []db.Queja) string {
	var b strings.Builder

	b.WriteString(
		fmt.Sprintf("%-10s| %-10s| %-10s| %-10s| %s|\n",
			"Client",
			"Severity",
			"Employee",
			"Item",
			"Reason"))
	for _, v := range data {
		b.WriteString(
			fmt.Sprintf("%-10d| %-10d| %-10d| %-10d| %s\n",
				v.Cliente,
				v.Gravedad,
				v.Empleado.Int32,
				v.Item.Int32,
				v.Motivo))
	}

	return b.String()
}
