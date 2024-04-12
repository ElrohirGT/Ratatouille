package waitress

import (
	"context"
	"fmt"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type getMenuItemsViewModel struct {
	table    table.Model
	errorMsg string
}

func (m getMenuItemsViewModel) Init() tea.Cmd { return handleGetMenuItems() }

func (m getMenuItemsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return CreateWaitressView(), nil
		case "q":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		}
	case global.ErrorDB:
		m.errorMsg = msg.Description
		return m, nil
	case table.Model:
		m.table = msg
		return m, nil
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
func (m getMenuItemsViewModel) View() string {
	var b strings.Builder

	b.WriteString("\n\n")

	b.WriteString(styles.GetTitleStyle().Render("Menu Items"))

	b.WriteString("\n\n")

	b.WriteString(baseStyle.Render(m.table.View()))

	b.WriteString("\n\n")

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleGetMenuItems() tea.Cmd {
	menuItems, err := global.Driver.GetMenuItems(context.Background())
	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Error retreiving clients"}
		}
	} else {
		table := parseMenuItemsToTable(menuItems)
		return func() tea.Msg {
			return table
		}
	}
}

func parseMenuItemsToTable(menuItems []db.GetMenuItemsRow) table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 15},
		{Title: "Name", Width: 30},
		{Title: "Price", Width: 30},
	}

	rows := make([]table.Row, len(menuItems))

	for i := range menuItems {
		item := menuItems[i]
		fmt.Println(string(item.ID))
		rows[i] = table.Row{fmt.Sprint(item.ID), item.Nombre, fmt.Sprint(item.Preciounitario)}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(8),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}