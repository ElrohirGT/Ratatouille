package waitress

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type GenerateBillView struct {
	forms    components.FormsModel
	account  string
	client   string
	errorMsg string
}

func (m GenerateBillView) Init() tea.Cmd {
	return nil
}

func (m GenerateBillView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		if newMsg.Type == tea.KeyEsc {
			return CreateWaitressView(), nil
		}
		if newMsg.Type == tea.KeyEnter {
			m.account = m.forms.FormInputs["Account"].Value
			m.client = m.forms.FormInputs["Client"].Value

			return m, handleGenerateBill(m.account, m.client)
		}

		newForm, cmds := m.forms.Update(msg)
		m.forms = newForm.(components.FormsModel)
		return m, cmds

	case global.ErrorDB:
		m.errorMsg = newMsg.Description
	case global.SuccesDB:
		bill := newMsg.Value.(db.GenerateBillRow)
		client, _ := global.Driver.GetClient(context.Background(), bill.Cliente)

		billText := renderBill(bill.Fecha.Format(time.Layout), m.account, client.Nombre, client.Nit, bill.Total)

		onConfirmation := func() (tea.Model, tea.Cmd) { return CreateWaitressView(), nil }

		return components.CreateAlert(billText, onConfirmation), nil
	}

	return m, nil
}

func (m GenerateBillView) View() string {
	var b strings.Builder

	b.WriteString(m.forms.View() + "\n\n")

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleGenerateBill(NoAccount, client string) tea.Cmd {
	if NoAccount == "" || client == "" {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Cannot have empty fields!"}
		}
	}

	noAccountConverted, err := strconv.Atoi(NoAccount)
	clientConverted, err2 := strconv.Atoi(client)

	if err != nil || err2 != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: "NoAccount client must be integers!"}
		}
	}

	return func() tea.Msg {
		bill, err := global.Driver.GenerateBill(context.Background(),
			db.GenerateBillParams{Cuenta: int32(noAccountConverted), Cliente: int32(clientConverted)})
		if err != nil {
			return global.ErrorDB{Description: err.Error()}
		} else {
			return global.SuccesDB{Description: "Open account correctly", Value: bill}
		}
	}

}

func renderBill(date, account, client, nit, total string) string {
	return fmt.Sprintf(`
	NEW BILL DETAILS
	=============================
	Date: %s
	Account: %s
	Client: %s
	NIT: %s
	Total: $%s
	=============================
	
`, date, account, client, nit, total)
}
