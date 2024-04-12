package waitress

import (
	"context"
	"strconv"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type TakeOrderView struct {
	forms    components.FormsModel
	amount   string
	account  string
	item     string
	errorMsg string
}

func (m TakeOrderView) Init() tea.Cmd {
	return nil
}

func (m TakeOrderView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateWaitressView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			m.account = m.forms.FormInputs["NO.account"].Value
			m.item = m.forms.FormInputs["ItemID"].Value
			m.amount = m.forms.FormInputs["Amount"].Value

			return m, handleTakeOrder(m.account, m.item, m.amount)
		}

		newForm, cmds := m.forms.Update(msg)
		m.forms = newForm.(components.FormsModel)
		return m, cmds

	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		return components.CreateConfirmation("Do you confirm your order?",
			func() (tea.Model, tea.Cmd) { return CreateWaitressView(), nil },
			func() (tea.Model, tea.Cmd) { return CreateTakeOrder(), nil }), nil
	}

	return m, nil
}

func (m TakeOrderView) View() string {
	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleTakeOrder(account, item, amount string) tea.Cmd {
	if account == "" || item == "" || amount == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	if validIntegers := global.CanConvertToInt32(account, item, amount); !validIntegers {
		return func() tea.Msg {
			return global.ErrorDB{Description: "All fields must be integers"}
		}

	}
	cuenta, _ := strconv.Atoi(account)
	platillo, _ := strconv.Atoi(item)
	cantidad, _ := strconv.Atoi(amount)

	return func() tea.Msg {
		order, err := global.Driver.TakeOrder(context.Background(),
			db.TakeOrderParams{Cuenta: int32(cuenta), Item: int32(platillo), Cantidad: int32(cantidad)})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "Order created correctly", Value: order}
		}
	}

}
