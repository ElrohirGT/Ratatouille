package components

import (
	"fmt"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/tui/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var focusedButton = styles.GetFocusedStyle().Render("[ Submit ]")
var blurredButton = fmt.Sprintf("[ %s ]", styles.GetDeactivateStyle().Render("Submit"))

type FormsInput struct {
	Value       string
	Placeholder string
	InputType   string
}

type FormsModel struct {
	Title      string
	FormInputs map[string]FormsInput

	FocusIndex  int
	inputsTitle []string
	inputs      []textinput.Model
}

func CreateForms(title string, formsInputs map[string]FormsInput) FormsModel {

	m := FormsModel{
		Title:       title,
		FocusIndex:  0,
		FormInputs:  formsInputs,
		inputsTitle: make([]string, 0),
		inputs:      make([]textinput.Model, 0),
	}

	var t textinput.Model
	for k, v := range formsInputs {
		t = textinput.New()
		t.Cursor.Style = styles.GetFocusedStyle()
		t.PromptStyle = styles.GetDeactivateStyle()
		t.TextStyle = styles.GetDeactivateStyle()
		t.CharLimit = 30
		t.Placeholder = v.Placeholder

		m.inputsTitle = append(m.inputsTitle, k)

		if v.InputType == "password" {
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}
		m.inputs = append(m.inputs, t)
	}

	return m
}

func (m FormsModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m FormsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter, tea.KeyUp, tea.KeyDown:
			s := msg.Type

			if s == tea.KeyEnter {
				return m, nil
			}
			if s == tea.KeyUp || s == tea.KeyShiftTab {
				m.FocusIndex--
			} else {
				m.FocusIndex++
			}

			if m.FocusIndex > len(m.inputs) {
				m.FocusIndex = 0
			} else if m.FocusIndex < 0 {
				m.FocusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.FocusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = styles.GetFocusedStyle()
					m.inputs[i].TextStyle = styles.GetFocusedStyle()
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = styles.GetDeactivateStyle()
				m.inputs[i].TextStyle = styles.GetDeactivateStyle()
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *FormsModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		// Assigning updated value to each input Entry.
		entry := m.FormInputs[m.inputsTitle[i]]
		entry.Value = m.inputs[i].Value()
		m.FormInputs[m.inputsTitle[i]] = entry
	}

	return tea.Batch(cmds...)
}

func (m FormsModel) View() string {
	var b strings.Builder

	b.WriteString(styles.GetTitleStyle().Render(m.Title))
	b.WriteRune('\n')
	b.WriteRune('\n')

	for i := range m.inputsTitle {
		b.WriteString(fmt.Sprintf("╭─・ %s \n", m.inputsTitle[i]))
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteString("\n\n")
		}
	}

	button := &blurredButton
	if m.FocusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
