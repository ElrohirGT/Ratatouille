package components

import (
	"fmt"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

var confirmationButton = styles.GetFocusedStyle().Render("[ Yes ]")
var confirmationButtonOff = fmt.Sprintf("[ %s ]", styles.GetDeactivateStyle().Render("Yes"))
var negationButton = styles.GetFocusedStyle().Render("[ No ]")
var negationButtonOff = fmt.Sprintf("[ %s ]", styles.GetDeactivateStyle().Render("No"))
var errorBtn = styles.GetFocusedStyle().Render("[ OK ]")

type ConfirmationModel struct {
	DisplayText    string
	focusButton    int
	onConfirmation func() tea.Cmd
	onNegation     func() (tea.Model, tea.Cmd)
	onError        func() (tea.Model, tea.Cmd)
	onSuccess      func() (tea.Model, tea.Cmd)
	errorMsg       string
}

func CreateConfirmation(
	displayText string,
	onConfirmation func() tea.Cmd,
	onNegation func() (tea.Model, tea.Cmd),
	onSuccess func() (tea.Model, tea.Cmd),
	onError func() (tea.Model, tea.Cmd)) ConfirmationModel {
	return ConfirmationModel{DisplayText: displayText,
		onConfirmation: onConfirmation,
		onNegation:     onNegation,
		onSuccess:      onSuccess,
		onError:        onError,
		errorMsg: "",
	}
}

func (m ConfirmationModel) Init() tea.Cmd {
	return nil
}

func (m ConfirmationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if m.errorMsg == "" {
			if newMsg.Type == tea.KeyRight {
				m.focusButton = 1
			}
			if newMsg.Type == tea.KeyLeft {
				m.focusButton = 0
			}
			if newMsg.Type == tea.KeyEnter && m.focusButton == 0 {
				return m, m.onConfirmation()
			}
			if newMsg.Type == tea.KeyEnter && m.focusButton == 1 {
				return m.onNegation()
			}
		} else {
			if newMsg.Type == tea.KeyEnter {
				return m.onError()
			}
		}
	case global.ErrorDB:
		m.errorMsg = newMsg.Description
		return m, nil
	case global.SuccesDB:
		return m.onSuccess()
	}
	return m, nil
}

func (m ConfirmationModel) View() string {
	var b strings.Builder

	if m.errorMsg == "" {
		b.WriteString(m.DisplayText)
		b.WriteString("\n\n")

		if m.focusButton == 0 {
			b.WriteString(confirmationButton + "   ")
			b.WriteString(negationButtonOff)
		} else {
			b.WriteString(confirmationButtonOff + "   ")
			b.WriteString(negationButton)
		}
	} else {
		b.WriteString("\n" + styles.GetErrorStyle().Render(m.errorMsg) + "\n\n")
		b.WriteString(errorBtn)
	}

	return b.String()
}
