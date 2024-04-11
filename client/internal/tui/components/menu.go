package components

import (
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuItem struct {
	Route           string
	ItemTitle       string
	ItemDescription string
}

func (i MenuItem) Title() string       { return i.ItemTitle }
func (i MenuItem) Description() string { return i.ItemDescription }
func (i MenuItem) FilterValue() string { return i.ItemTitle }

type MenuModel struct {
	Title        string
	Items        []MenuItem
	FocusIndex   int
	SelectedItem MenuItem
}

func CreateMenu(title string, items []MenuItem) MenuModel {

	return MenuModel{Title: title, Items: items}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter, tea.KeyUp, tea.KeyDown:
			s := msg.Type

			if s == tea.KeyUp || s == tea.KeyShiftTab {
				m.FocusIndex--
			} else {
				m.FocusIndex++
			}

			if m.FocusIndex > len(m.Items)-1 {
				m.FocusIndex = 0
			} else if m.FocusIndex < 0 {
				m.FocusIndex = len(m.Items) - 1
			}
		}
	}

	m.SelectedItem = m.Items[m.FocusIndex]

	return m, nil
}

func (m MenuModel) View() string {
	var sb strings.Builder

	sb.WriteString("\n")
	sb.WriteString(styles.GetTitleStyle().Render(m.Title))
	sb.WriteString("\n\n")

	for i, v := range m.Items {

		var titleStyle, descriptionStyle lipgloss.Style

		if i == m.FocusIndex {
			titleStyle, descriptionStyle = styles.GetMenuItemFocusStyle()
		} else {
			titleStyle, descriptionStyle = styles.GetMenuItemDefaultStyle()
		}
		sb.WriteString(titleStyle.Render(v.ItemTitle) + "\n" + descriptionStyle.Render(v.Description()) + "\n\n")

	}
	sb.WriteString(styles.GetHelpStyle().Render("↑ up • ↓ down • Enter select  • Esc quit"))

	return sb.String()
}
