package auth

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AuthModel struct {
	Username   string
	Password   string
	FocusIndex int
	Inputs     []textinput.Model
}

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle.Copy()
	noStyle      = lipgloss.NewStyle()
	focusedButton = fmt.Sprintf("[ %s ]", focusedStyle.Render("Submit"))
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

func InitialModel() AuthModel {

	// Initializing model.
	m := AuthModel{
		Username: "",
		Password: "",
		Inputs:   make([]textinput.Model, 2),
	}

	// Creating TextInputs.
	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 30

		switch i {
		case 0:
			t.Placeholder = "Username"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}
		m.Inputs[i] = t
	}

	return m
}

func (m AuthModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AuthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter, tea.KeyUp, tea.KeyDown:

			s := msg.Type

			// Submiting form
			if s == tea.KeyEnter && m.FocusIndex == len(m.Inputs) {
				return m, tea.Quit
			}
			
			// Changin input focus
			if s == tea.KeyUp || s == tea.KeyShiftTab {
				m.FocusIndex--

			} else {
				m.FocusIndex++
			}

			// Cycling inputs if necessary
			if m.FocusIndex > len(m.Inputs) {
				m.FocusIndex = 0
			} else if m.FocusIndex < 0 {
				m.FocusIndex = len(m.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.Inputs))
			for i := 0; i <= len(m.Inputs)-1; i++ {
				if i == m.FocusIndex {
					// Set focused state
					cmds[i] = m.Inputs[i].Focus()
					m.Inputs[i].PromptStyle = focusedStyle
					m.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.Inputs[i].Blur()
				m.Inputs[i].PromptStyle = noStyle
				m.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *AuthModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.Inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}

	m.Username = m.Inputs[0].Value()
	m.Password = m.Inputs[1].Value()

	return tea.Batch(cmds...)
}

func (m AuthModel) View() string {
	var b strings.Builder
	
	for i := range m.Inputs {
		b.WriteString(m.Inputs[i].View())
		if i < len(m.Inputs) - 1 {
			b.WriteRune('\n')
		}
	}
	
	button := &blurredButton
	
	if m.FocusIndex == len(m.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	
	b.WriteString(blurredStyle.Render("Move with arrows. Submit with <Enter>"))
	
	return b.String()
}
