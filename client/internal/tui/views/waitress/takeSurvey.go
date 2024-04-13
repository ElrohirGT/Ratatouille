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

type takeSurveyModel struct {
	forms    components.FormsModel
	employee string
	client   string
	kindness string
	speed    string
	errorMsg string
}

func (m takeSurveyModel) Init() tea.Cmd {
	return nil
}

func (m takeSurveyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newForm, cmds := m.forms.Update(msg)
	m.forms = newForm.(components.FormsModel)

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateWaitressView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			m.employee = m.forms.FormInputs["Employee"].Value
			m.client = m.forms.FormInputs["Client"].Value
			m.kindness = m.forms.FormInputs["Kindness"].Value
			m.speed = m.forms.FormInputs["Speed"].Value
			return m, handleTakeSurvey(m.employee, m.client, m.kindness, m.speed)
		}
	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		return m, nil
	}

	return m, cmds
}

func (m takeSurveyModel) View() string {

	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleTakeSurvey(employee, client, kindness, speed string) tea.Cmd {

	if employee == "" || client == "" || kindness == "" || speed == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	if validIntegers := global.CanConvertToInt32(employee, client, kindness, speed); !validIntegers {
		return func() tea.Msg {
			return global.ErrorDB{Description: "All values must be Integers"}
		}
	}

	employeeConverted, _ := strconv.Atoi(employee)
	clientConverted, _ := strconv.Atoi(client)
	kindnessConverted, _ := strconv.Atoi(kindness)
	speedConverted, _ := strconv.Atoi(speed)

	return func() tea.Msg {
		err := global.Driver.TakeSurvey(context.Background(),
			db.TakeSurveyParams{
				Empleado:        int32(employeeConverted),
				Cliente:         int32(clientConverted),
				Gradoamabilidad: int32(kindnessConverted),
				Gradoexactitud:  int32(speedConverted)})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "Survey registered correctly"}
		}
	}
}
