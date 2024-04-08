package components

import (
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/tui/constants"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type MenuItem struct {
	ItemTitle       string
	ItemDescription string
}

func (i MenuItem) Title() string       { return i.ItemTitle }
func (i MenuItem) Description() string { return i.ItemDescription }
func (i MenuItem) FilterValue() string { return i.ItemTitle }

type MenuModel struct {
	Items     []MenuItem
	list      list.Model
	FocusItem MenuItem
}

func InitialModel(title string, items []MenuItem) MenuModel {

	menuItems := make([]list.Item, 0)

	for i := 0; i < len(items); i++ {
		menuItems = append(menuItems, items[i])
	}

	// Config List component
	delegate := list.NewDefaultDelegate()
	delegate.Styles = constants.GetMenuItemStyle()
	modelList := list.New(menuItems, delegate, 0, 0)
	modelList.Title = title

	return MenuModel{Items: items, list: modelList}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		} else if msg.Type == tea.KeyEnter {
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				m.FocusItem = i
			}
		}
	case tea.WindowSizeMsg:
		x, y := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-x, msg.Height-y)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m MenuModel) View() string {
	var sb strings.Builder

	sb.WriteString(docStyle.Render(m.list.View()))
	return sb.String()
}
