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

			return m, SignUser(m.username, m.password, m.role)
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

func SignUser(username, password, role string) tea.Cmd {
	
	if username == "" || password == "" || role == ""{
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	v, err := strconv.Atoi(role)
	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Role must be an integer between 1 and 4"}
		}
	}

	return func() tea.Msg {
		encryptedPassword := global.EncryptSHA256(password)
		err := global.Driver.SignIn(context.Background(),
			db.SignInParams{Nombre: username, Contraseña: encryptedPassword, Tipo: int32(v)})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "User created correctly"}
		}
	}
}
