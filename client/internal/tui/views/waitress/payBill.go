package waitress

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/db"
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	"github.com/ElrohirGT/Ratatouille/internal/tui/global"
	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var focusedButton = styles.GetFocusedStyle().Render("[ Pay ]")
var blurredButton = fmt.Sprintf("[ %s ]", styles.GetDeactivateStyle().Render("Pay"))

type payBillModel struct {
	noBill      int32
	amountToPay float64
	optionFocus int

	sliderOptions []string
	sliderFocus   int

	canWriteInput bool

	moneyInput textinput.Model
	errorMsg   string
}

func CreatePayBillView(noBill int32, amountToPay float64) payBillModel {

	sliderOptions := []string{"Cash", "Credit Card", "Debit Card"}
	moneyInput := textinput.New()
	moneyInput.Cursor.Style = styles.GetFocusedStyle()
	moneyInput.PromptStyle = styles.GetDeactivateStyle()
	moneyInput.TextStyle = styles.GetDeactivateStyle()
	moneyInput.CharLimit = 30
	moneyInput.Placeholder = "10.00"

	return payBillModel{
		noBill: noBill,
		amountToPay:   amountToPay,
		sliderOptions: sliderOptions,
		moneyInput:    moneyInput}
}

func (m payBillModel) Init() tea.Cmd {
	return m.moneyInput.Focus()
}

func (m payBillModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch newMsg := msg.(type) {
	case tea.KeyMsg:
		switch newMsg.Type {
		case tea.KeyUp, tea.KeyDown, tea.KeyEnter, tea.KeyLeft, tea.KeyRight:
			if newMsg.Type == tea.KeyEnter && m.optionFocus == 2 {
				return m, handleAddPayment( m.noBill,
					m.moneyInput.Value(), 
					m.sliderOptions[m.sliderFocus])
			}
			if newMsg.Type == tea.KeyUp {
				newPointer := (m.optionFocus - 1) % 3
				m.optionFocus = int(math.Abs(float64(newPointer)))
			}
			if newMsg.Type == tea.KeyDown {
				newPointer := (m.optionFocus + 1) % 3
				m.optionFocus = int(math.Abs(float64(newPointer)))
			}
			if newMsg.Type == tea.KeyLeft && m.optionFocus == 0{
				newPointer := (m.sliderFocus - 1) % 3
				m.sliderFocus = int(math.Abs(float64(newPointer)))
			}
			if newMsg.Type == tea.KeyRight && m.optionFocus == 0{
				newPointer := (m.sliderFocus + 1) % 3
				m.sliderFocus = int(math.Abs(float64(newPointer)))
			}

			if m.optionFocus == 1 {
				m.canWriteInput = true
				m.moneyInput.PromptStyle = styles.GetFocusedStyle()
				m.moneyInput.TextStyle = styles.GetFocusedStyle()
				cmd := m.moneyInput.Focus()
				return m, cmd
			} else {
				// Remove focused state
				m.canWriteInput = false
				m.moneyInput.Blur()
				m.moneyInput.PromptStyle = styles.GetDeactivateStyle()
				m.moneyInput.TextStyle = styles.GetDeactivateStyle()
				return m, nil
			}
		}
	case global.ErrorDB:
		m.errorMsg = newMsg.Description
		return m, nil
	case global.PaymentSuccess:
		m.amountToPay -= newMsg.Amount
		if m.amountToPay <= 0 {
			onConfirmation := func() (tea.Model, tea.Cmd) { return CreateTakeSurvey(), nil }
			return components.CreateAlert(
				"Bill payed succesfully!",
				onConfirmation), nil
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.moneyInput, cmd = m.moneyInput.Update(msg)
	return m, cmd
}

func (m payBillModel) View() string {
	var b strings.Builder
	var sliderStyle lipgloss.Style
	if m.optionFocus == 0 {
		sliderStyle = styles.GetFocusedStyle()
	} else {
		sliderStyle = styles.GetDeactivateStyle()
	}
	button := &blurredButton
	if m.optionFocus == 2 {
		button = &focusedButton
	}

	b.WriteString(styles.GetTitleStyle().Render("Payment Portal"))
	b.WriteRune('\n')
	b.WriteRune('\n')
	b.WriteString(fmt.Sprintf("Remain to pay: %f", m.amountToPay))
	b.WriteRune('\n')

	b.WriteString("╭─・ Payment Method \n")
	b.WriteString(sliderStyle.Render(
		fmt.Sprintf("<< %s >>", m.sliderOptions[m.sliderFocus])))
	b.WriteString("\n")

	b.WriteString("╭─・ Amount \n")
	b.WriteString(m.moneyInput.View())

	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	if m.errorMsg != "" {
		b.WriteString(styles.GetErrorStyle().Render(m.errorMsg))
	}

	return b.String()
}

func handleAddPayment(noBill int32, amount string, paymentMethod string) tea.Cmd {

	amountConverted, err := strconv.ParseFloat(strings.TrimSpace(amount), 64)
	
	if amountConverted <= 0 {
		return func() tea.Msg {
			return global.ErrorDB{Description: "Amount most be different of zero"}
		}
	}

	var selectedMethod int32
	switch paymentMethod {
	case "Cash":
		selectedMethod = 1
	case "Credit Card":
		selectedMethod = 2
	case "Debit Card":
		selectedMethod = 3
	}

	err = global.Driver.AddPayment(context.Background(),
		db.AddPaymentParams{
			Tipo:    selectedMethod,
			Monto:   amountConverted,
			Factura: noBill,
		})

	if err != nil {
		return func() tea.Msg {
			return global.ErrorDB{Description: err.Error()}
		}
	} else {
		return func() tea.Msg {
			return global.PaymentSuccess{Amount: amountConverted}
		}
	}

}
