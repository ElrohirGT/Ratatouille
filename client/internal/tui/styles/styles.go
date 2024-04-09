package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

// COLORS
const (
	PrimaryColor    = lipgloss.Color("#94ffa6")
	DeactivateColor = lipgloss.Color("#5d9e68")
	BlackColor      = lipgloss.Color("#000000")
)

// STYLES
var focusedStyle = lipgloss.NewStyle().Foreground(PrimaryColor)
var deactiveStyle = lipgloss.NewStyle().Foreground(DeactivateColor)
var titleStyle = lipgloss.NewStyle().Foreground(BlackColor).Background(PrimaryColor).Bold(true)

func GetFocusedStyle() lipgloss.Style {
	return focusedStyle
}
func GetDeactivateStyle() lipgloss.Style {
	return deactiveStyle
}
func GetTitleStyle() lipgloss.Style {
	return titleStyle
}

func GetMenuItemStyle() list.DefaultItemStyles {

	style := list.NewDefaultDelegate().Styles

	style.SelectedTitle = style.SelectedTitle.
		Foreground(PrimaryColor).
		BorderLeftForeground(PrimaryColor)
	style.SelectedDesc = style.SelectedDesc.
		Foreground(DeactivateColor).
		BorderLeftForeground(PrimaryColor)

	return style

}
