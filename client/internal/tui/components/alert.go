package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type AlertModel struct {
	DisplayText    string
	onConfirmation func() (tea.Model, tea.Cmd)
}

func CreateAlert(
	displayText string,
	onConfirmation func() (tea.Model, tea.Cmd)) AlertModel {
	return AlertModel{DisplayText: displayText,
		onConfirmation: onConfirmation,
	}
}

func (m AlertModel) Init() tea.Cmd {
	return nil
}

func (m AlertModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEnter {
			return m.onConfirmation()
		}
	}
	return m, nil
}

func (m AlertModel) View() string {
	var b strings.Builder

	b.WriteString(m.DisplayText)
	b.WriteString("\n\n")

	b.WriteString(confirmationButton + "   ")

	return b.String()
}
