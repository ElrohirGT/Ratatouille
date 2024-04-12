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

type getActiveAccountsModel struct {
	table    table.Model
	errorMsg string
}

func (m getActiveAccountsModel) Init() tea.Cmd { return handleGetClients() }

func (m getActiveAccountsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
func (m getActiveAccountsModel) View() string {
	var b strings.Builder

	b.WriteString("\n\n")

	b.WriteString(styles.GetTitleStyle().Render("Clients"))

	b.WriteString("\n\n")

	b.WriteString(baseStyle.Render(m.table.View()))

	b.WriteString("\n\n")

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleGetActiveAccounts() tea.Cmd {
	accounts, err := global.Driver.GetActiveAccounts(context.Background())
	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Error retreiving accounts"}
		}
	} else {
		table := parseAccountsToTable(accounts)
		return func() tea.Msg {
			return table
		}
	}
}

func parseAccountsToTable(accounts []db.Cuentum) table.Model {
	columns := []table.Column{
		{Title: "Table", Width: 15},
		{Title: "NO. Account", Width: 30},
		{Title: "Total", Width: 30},
	}

	rows := make([]table.Row, len(accounts))

	for i := range accounts {
		account := accounts[i]
		rows[i] = table.Row{
			fmt.Sprint(account.Mesa),
			fmt.Sprint(account.Numcuenta),
			fmt.Sprint(account.Total)}
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

/*
type getClientsViewModel struct {
	table table.Model
}

func (m getClientsViewModel) Init() tea.Cmd { return nil }

func (m getClientsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
func (m getClientsViewModel) View() string {
	return ""
}
*/
