package components

import (
	"fmt"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

var confirmationButton = styles.GetFocusedStyle().Render("[ Yes ]")
var confirmationButtonOff = fmt.Sprintf("[ %s ]", styles.GetDeactivateStyle().Render("Yes"))
var negationButton = styles.GetFocusedStyle().Render("[ No ]")
var negationButtonOff = fmt.Sprintf("[ %s ]", styles.GetDeactivateStyle().Render("No"))

type ConfirmationModel struct {
	DisplayText    string
	focusButton    int
	onConfirmation func() (tea.Model, tea.Cmd)
	onNegation     func() (tea.Model, tea.Cmd)
}

func CreateConfirmation(
	displayText string,
	onConfirmation,
	onNegation func() (tea.Model, tea.Cmd)) ConfirmationModel {
	return ConfirmationModel{DisplayText: displayText,
		onConfirmation: onConfirmation,
		onNegation:     onNegation}
}

func (m ConfirmationModel) Init() tea.Cmd {
	return nil
}

func (m ConfirmationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch newMsg := msg.(type) {
	case tea.KeyType:
		if newMsg == tea.KeyRight {
			m.focusButton = 1
		}
		if newMsg == tea.KeyLeft {
			m.focusButton = 0
		}
		if newMsg == tea.KeyEnter && m.focusButton == 0 {
			return m.onConfirmation()
		}
		if newMsg == tea.KeyEnter && m.focusButton == 1 {
			return m.onNegation()
		}
	}
	return m, nil
}

func (m ConfirmationModel) View() string {
	var b strings.Builder

	b.WriteString(m.DisplayText)
	if m.focusButton == 0 {
		b.WriteString(confirmationButton + "   ")
		b.WriteString(negationButtonOff)
	} else {
		b.WriteString(confirmationButtonOff + "   ")
		b.WriteString(negationButton)
	}

	return ""
}
