package auth

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

type SignUpModel struct {
	forms    components.FormsModel
	username string
	password string
	employee string
	role     string
	errorMsg string
}

func (m SignUpModel) Init() tea.Cmd {
	return nil
}

func (m SignUpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newForm, cmds := m.forms.Update(msg)
	m.forms = newForm.(components.FormsModel)

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateAuthView(), cmds
		}
		if newMsg.Type == tea.KeyEnter && m.forms.FocusIndex == len(m.forms.FormInputs) {

			m.username = m.forms.FormInputs["Username"].Value
			m.password = m.forms.FormInputs["Password"].Value
			m.role = m.forms.FormInputs["Role"].Value
			m.employee = m.forms.FormInputs["EmployeeID"].Value

			return m, handleSignUser(m.username, m.password, m.role, m.employee)
		}
	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		return CreateAuthView(), nil
	}

	return m, cmds
}

func (m SignUpModel) View() string {
	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleSignUser(username, password, role, employeeID string) tea.Cmd {

	if username == "" || password == "" || role == "" || employeeID == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	tipo, err := strconv.Atoi(role)
	empleado, err2 := strconv.Atoi(employeeID)

	if err != nil || err2 != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Role and Employee must be integers!"}
		}
	}

	return func() tea.Msg {
		encryptedPassword := global.EncryptSHA256(password)
		err := global.Driver.SignIn(context.Background(),
			db.SignInParams{Nombre: username, Contraseña: encryptedPassword, Tipo: int32(tipo), Empleado: int32(empleado)} )
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "User created correctly"}
		}
	}
}
