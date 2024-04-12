package waitress

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type RegisterComplain struct {
	forms    components.FormsModel
	client   string
	severity string
	reason   string
	employee string
	item     string
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
			m.client = m.forms.FormInputs["Client"].Value
			m.severity = m.forms.FormInputs["Severity"].Value
			m.reason = m.forms.FormInputs["Reason"].Value
			m.employee = m.forms.FormInputs["Employee"].Value
			m.item = m.forms.FormInputs["Item"].Value

			return m, handleRegisterComplain(m.client, m.severity, m.reason, m.employee, m.item)
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

func handleRegisterComplain(client, severity, reason, employee, item string) tea.Cmd {
	if client == "" || severity == "" || reason == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Client, Severity, Reason MUST not be null!"}
		}
	}

	if validIntegers := global.CanConvertToInt32(client, severity); !validIntegers {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Client, Severity, Reason MUST must be integers"}
		}
	}
	cliente, _ := strconv.Atoi(client)
	gravedad, _ := strconv.Atoi(severity)

	value, err := strconv.Atoi(employee)
	value2, err2 := strconv.Atoi(item)

	empleado := sql.NullInt32{}
	comida := sql.NullInt32{}

	if err == nil {
		empleado.Int32 = int32(value)
	}
	if err2 == nil {
		comida.Int32 = int32(value2)
	}

	return func() tea.Msg {
		err := global.Driver.RegisterComplaint(context.Background(),
			db.RegisterComplaintParams{
				Cliente:  int32(cliente),
				Motivo:   reason,
				Gravedad: int32(gravedad),
				Empleado: empleado,
				Item:     comida})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "Complaint created correctly"}
		}
	}

}
