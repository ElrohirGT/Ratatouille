package chef

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

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

type ChefViewModel struct {
	table    table.Model
	errorMsg string
}

func CreateChefViewModel() ChefViewModel {
	return ChefViewModel{table: table.New()}
}

func (m ChefViewModel) Init() tea.Cmd { return HandleGetDishes() }

func (m ChefViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		case "q":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "enter":
			row := m.table.SelectedRow()
			println(row[0])
			if row[4] == "Pedido" {
				return m, handlePrepareDish(row[0])
			} else if row[4] == "En preparación" {
				return m, handleFinishDish(row[0])
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
func (m ChefViewModel) View() string {
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

func HandleGetDishes() tea.Cmd {
	orders, err := global.Driver.GetPendingDishes(context.Background())
	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Error retreiving clients"}
		}
	} else {
		table := parseOrdersToTable(orders)
		return func() tea.Msg {
			return table
		}
	}
}

func handlePrepareDish(orderId string) tea.Cmd {

	v, err := strconv.Atoi(orderId)

	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Order ID most be an integer"}
		}
	}

	err = global.Driver.SetOrderPreparing(context.Background(), int32(v))
	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: err.Error()}
		}
	} else {
		return HandleGetDishes()
	}

}

func handleFinishDish(orderId string) tea.Cmd {

	v, err := strconv.Atoi(orderId)

	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Order ID most be an integer"}
		}
	}

	err = global.Driver.SetOrderDelivered(context.Background(), int32(v))
	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Error Setting orders"}
		}
	} else {
		return HandleGetDishes()
	}

}

func parseOrdersToTable(orders []db.GetPendingDishesRow) table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 15},
		{Title: "Date", Width: 30},
		{Title: "Name", Width: 30},
		{Title: "Quantity", Width: 15},
		{Title: "State", Width: 15},
	}

	rows := make([]table.Row, len(orders))

	for i := range orders {
		order := orders[i]
		fmt.Println(string(order.ID))
		rows[i] = table.Row{
			fmt.Sprint(order.ID),
			order.Fecha.Format(time.Layout),
			order.Nombredelitemmenu,
			fmt.Sprint(order.Cantidad),
			order.Estadodelpedido,
		}
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
