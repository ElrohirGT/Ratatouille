package constants

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/golang-collections/collections/stack"
)

var ViewStack *stack.Stack = stack.New()

func GetViewStack() *stack.Stack {
	return ViewStack
}

const (
	PrimaryColor    = lipgloss.Color("#94ffa6")
	DeactivateColor = lipgloss.Color("#5d9e68")
)

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
