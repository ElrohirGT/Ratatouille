package components

import (
	"fmt"
	"strings"

	"github.com/ElrohirGT/Ratatouille/internal/tui/constants"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/exp/maps"
)

var focusedButton = constants.GetFocusedStyle().Render("[ Submit ]")
var blurredButton = fmt.Sprintf("[ %s ]", constants.GetDeactivateStyle().Render("Submit"))

type FormsInput struct {
	Value       string
	Placeholder string
	InputType   string
}

type FormsModel struct {
	Title      string
	FormInputs map[string]FormsInput

	FocusIndex int
	inputs     []textinput.Model
}

func CreateForms(title string, formsInputs map[string]FormsInput) FormsModel {

	m := FormsModel{
		Title:      title,
		FormInputs: formsInputs,
		inputs:     make([]textinput.Model, 0),
	}

	var t textinput.Model
	for _, v := range formsInputs {
		t = textinput.New()
		t.Cursor.Style = constants.GetFocusedStyle()
		t.PromptStyle = constants.GetDeactivateStyle()
		t.TextStyle = constants.GetDeactivateStyle()
		t.CharLimit = 30
		t.Placeholder = v.Placeholder

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

			if s == tea.KeyEnter && m.FocusIndex != len(m.FormInputs) {
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
					m.inputs[i].PromptStyle = constants.GetFocusedStyle()
					m.inputs[i].TextStyle = constants.GetFocusedStyle()
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = constants.GetDeactivateStyle()
				m.inputs[i].TextStyle = constants.GetDeactivateStyle()
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *FormsModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	keys := maps.Keys(m.FormInputs)

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		// Assigning updated value to each input Entry.
		entry := m.FormInputs[keys[i]]
		entry.Value = m.inputs[i].Value()
		m.FormInputs[keys[i]] = entry
	}

	return tea.Batch(cmds...)
}

func (m FormsModel) View() string {
	var b strings.Builder
	keys := maps.Keys(m.FormInputs)

	b.WriteString(constants.GetTitleStyle().Render(m.Title))
	b.WriteRune('\n')
	b.WriteRune('\n')

	for i := range keys {
		b.WriteString(fmt.Sprintf("╭─・ %s \n", keys[i]))
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.FocusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
