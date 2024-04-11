package auth

import (
	"context"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type LoginModel struct {
	forms    components.FormsModel
	username string
	password string
	errorMsg string
}

func (m LoginModel) Init() tea.Cmd {
	return nil
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newForm, cmds := m.forms.Update(msg)
	m.forms = newForm.(components.FormsModel)

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateAuthView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			m.username = m.forms.FormInputs["Username"].Value
			m.password = m.forms.FormInputs["Password"].Value
			return m, LoginUser(m.username, m.password)
		}
	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		return CreateAuthView(), nil
	}

	return m, cmds
}

func (m LoginModel) View() string {

	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")
	
	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func LoginUser(username, password string) tea.Cmd {
	
	if username == "" || password == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	return func() tea.Msg {
		encryptedPassword := global.EncryptSHA256(password)
		v, err := global.Driver.LogIn(context.Background(),
			db.LogInParams{Nombre: username, Contraseña: encryptedPassword})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			global.Username = v
			return global.SuccesDB{Description: "User created correctly"}
		}
	}
}

