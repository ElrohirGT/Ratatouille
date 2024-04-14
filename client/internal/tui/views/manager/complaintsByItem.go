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

type complaintsByItemModel struct {
	forms     components.FormsModel
	data      []db.Queja
	startDate string
	endDate   string
	item  string
	errorMsg  string
}

func (m complaintsByItemModel) Init() tea.Cmd {
	return nil
}

func (m complaintsByItemModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
			m.item = m.forms.FormInputs["Item"].Value
			return m, handleFetchComplainsByItem(m.item, m.startDate, m.endDate)
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

func (m complaintsByItemModel) View() string {

	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if len(m.data) > 0 {
		b.WriteString(printComplaintsByItemData(m.data))
	}

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleFetchComplainsByItem(item, startDate, endDate string) tea.Cmd {

	if startDate == "" || endDate == "" || item == ""{
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	date1, err := global.ParseDate(startDate)
	date2, err2 := global.ParseDate(endDate)
	itemMenu, err3 := strconv.Atoi(item)

	if err != nil || err2 != nil || err3 != nil{
		return func() tea.Msg {
			return global.ErrorDB{Description: "Not valid date format"}
		}
	}

	return func() tea.Msg {
		v, err := global.Driver.GetComplaintsForDishBetween(context.Background(),
			db.GetComplaintsForDishBetweenParams{
				Item: sql.NullInt32{Int32: int32(itemMenu), Valid: true},
				Fecha: date1, Fecha_2: date2})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "User created correctly", Value: v}
		}
	}
}

func printComplaintsByItemData(data []db.Queja) string {
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
