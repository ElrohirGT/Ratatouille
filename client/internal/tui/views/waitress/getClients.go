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

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type getClientsViewModel struct {
	table table.Model
	errorMsg string
}

func (m getClientsViewModel) Init() tea.Cmd { return handleGetClients() }

func (m getClientsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter": 
			return m, handleGetClients()
		}
	case global.ErrorDB:
		m.errorMsg = msg.Description
		return m, nil
	case table.Model:
		fmt.Println("SOMEONE CALL ME")
		m.table = msg
		return m, nil
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
func (m getClientsViewModel) View() string {
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

func handleGetClients() tea.Cmd {
	clients, err := global.Driver.GetClients(context.Background())
	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Error retreiving clients"}
		} 
	} else {
		table := parseClientsToTable(clients)
		return func() tea.Msg {
			return table
		} 
	}
}

func parseClientsToTable(clients []db.Cliente) table.Model{
	columns := []table.Column{
		{Title: "ID", Width: 15},
		{Title: "Nombre", Width: 30},
		{Title: "Direccion", Width: 30},
		{Title: "Nit", Width: 15},
	}

	rows := make([]table.Row, len(clients))
	
	for i := range clients {
		client := clients[i]
		fmt.Println(string(client.ID))
		rows[i] = table.Row{ fmt.Sprint(client.ID), client.Nombre, client.Direccion.String ,client.Nit }
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