package utils

import "github.com/charmbracelet/lipgloss"

var colors = map[string]lipgloss.Color{
	"red":    lipgloss.Color("1"),
	"green":  lipgloss.Color("2"),
	"yellow": lipgloss.Color("3"),
	"blue":   lipgloss.Color("4"),
	"purple": lipgloss.Color("5"),
	"dark":   lipgloss.Color("8"),
	"light":  lipgloss.Color("15"),
	"gray":   lipgloss.Color("7"),
	"white":  lipgloss.Color("15"),
}

var Styles = map[string]lipgloss.Style{
	"selectedNotChecked": lipgloss.NewStyle().Foreground(colors["yellow"]).Bold(true),
	"selectedChecked":    lipgloss.NewStyle().Foreground(colors["yellow"]).Faint(true),
	"checked":            lipgloss.NewStyle().Foreground(colors["white"]).Faint(true),
	"notChecked":         lipgloss.NewStyle().Foreground(colors["white"]).Bold(true),
	"footer":             lipgloss.NewStyle().Foreground(colors["blue"]), //.Italic(true),
	"title":              lipgloss.NewStyle().Foreground(colors["green"]).Bold(true).Align(lipgloss.Left),
	"alertTitle":         lipgloss.NewStyle().Foreground(colors["light"]).Bold(true).Padding(0, 1, 0, 1).Align(lipgloss.Left).Background(colors["red"]),
	"error":              lipgloss.NewStyle().Foreground(colors["red"]).Bold(true).Width(80),
}
