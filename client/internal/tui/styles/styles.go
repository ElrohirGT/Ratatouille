package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// COLORS
const (
	PrimaryColor    = lipgloss.Color("#94ffa6")
	SecondaryColor  = lipgloss.Color("#b259e3")
	DeactivateColor = lipgloss.Color("#5d9e68")
	BlackColor      = lipgloss.Color("#000000")
	WhiteColor      = lipgloss.Color("#ffffff")
	GreyColor       = lipgloss.Color("#b5b5b5")
)

// STYLES
var (
	focusedStyle  = lipgloss.NewStyle().Foreground(PrimaryColor)
	deactiveStyle = lipgloss.NewStyle().Foreground(DeactivateColor)
	titleStyle    = lipgloss.NewStyle().
			Foreground(WhiteColor).
			Bold(true).
			BorderStyle(lipgloss.RoundedBorder()).
			PaddingLeft(2).
			PaddingRight(2).
			MarginLeft(3)

	ItemTitleFocus = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(PrimaryColor).
			BorderLeft(true).
			PaddingLeft(1).
			MarginLeft(3)

	ItemDescriptionFocus = lipgloss.NewStyle().
				Foreground(DeactivateColor).
				Bold(true).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(PrimaryColor).
				BorderLeft(true).
				PaddingLeft(1).
				MarginLeft(3)
	
	ItemTitleDefault = lipgloss.NewStyle().
		Foreground(WhiteColor).
		Bold(true).
		PaddingLeft(1).
		MarginLeft(3)

	ItemDescriptionDefault = lipgloss.NewStyle().
		Foreground(GreyColor).
		PaddingLeft(1).
		MarginLeft(3)
	
	helpStyle = lipgloss.NewStyle().
		Foreground(GreyColor)
)

func GetFocusedStyle() lipgloss.Style {
	return focusedStyle
}
func GetDeactivateStyle() lipgloss.Style {
	return deactiveStyle
}
func GetTitleStyle() lipgloss.Style {
	return titleStyle
}

func GetMenuItemFocusStyle() (lipgloss.Style, lipgloss.Style) {
	return ItemTitleFocus, ItemDescriptionFocus
}

func GetMenuItemDefaultStyle() (lipgloss.Style, lipgloss.Style) {
	return ItemTitleDefault, ItemDescriptionDefault
}

func GetHelpStyle() (lipgloss.Style) {
	return helpStyle
}