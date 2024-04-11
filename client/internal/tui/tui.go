package tui

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/views/auth"
	tea "github.com/charmbracelet/bubbletea"
)

type TUI struct {
	id          int
	username    string
	role        string
	currentView tea.Model
}

func CreateTUI() TUI {
	return TUI{currentView: auth.CreateAuthView()}
}

func (t TUI) Init() tea.Cmd {
	return nil
}

func (t TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// returns the new view
	newView, cmd := t.currentView.Update(msg)

	t.currentView = newView

	return t, cmd
}

func (t TUI) View() string {
	return t.currentView.View()
}
