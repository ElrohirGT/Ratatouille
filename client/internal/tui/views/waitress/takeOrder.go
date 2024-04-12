package waitress

import (
	"context"
	"fmt"
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
		if newMsg.Type == tea.KeyEnter && m.forms.FocusIndex == len(m.forms.FormInputs) {
			m.account = m.forms.FormInputs["NO.account"].Value
			m.item = m.forms.FormInputs["ItemID"].Value
			m.amount = m.forms.FormInputs["Amount"].Value

			onConfirmation := func() tea.Cmd { return handleTakeOrder(m.account, m.item, m.amount) }
			onNegation := func() (tea.Model, tea.Cmd) { return CreateWaitressView(), nil }
			onSuccess := func() (tea.Model, tea.Cmd) { return CreateWaitressView(), nil }
			onError := func() (tea.Model, tea.Cmd) { return CreateWaitressView(), nil }

			return components.CreateConfirmation(renderOrder(m.account, m.item, m.amount), 
			onConfirmation, 
			onNegation, 
			onSuccess, 
			onError), nil
		}

		newForm, cmds := m.forms.Update(msg)
		m.forms = newForm.(components.FormsModel)
		return m, cmds
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

func renderOrder(account, item, amount string) string {
	return fmt.Sprintf(`
	NEW ORDER DETAILS
	=============================
	Account No: %s
	Item: %s
	Amount: %s units
	=============================
	CONFIRM?
`, account, item, amount)	
}