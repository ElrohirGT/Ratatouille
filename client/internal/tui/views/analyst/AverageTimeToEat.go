package analyst

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type AverageTimeToEatModel struct {
	
	startDate string
	endDate   string
	form      components.FormsModel
}

func CreateAverateTimeToEatModel() AverageTimeToEatModel {

	inputFields := map[string]components.FormsInput{
		"StartDate": {Placeholder: "XX/XX/XXXX"},
		"EndDate":   {Placeholder: "XX/XX/XXXX"},
	}

	form := components.CreateForms("Averate Time to Eat", inputFields)

	return AverageTimeToEatModel{form: form}
}

func (m AverageTimeToEatModel) Init() tea.Cmd {
	return nil
}

func (m AverageTimeToEatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	newForm, cmd := m.form.Update(msg)
	m.form = newForm.(components.FormsModel)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEsc {
			ViewStack.Pop()
			return m, nil
		}
	}
	return m, cmd
}

func (m AverageTimeToEatModel) View() string {

	return m.form.View()

}
