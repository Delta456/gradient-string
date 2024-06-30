package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/delta456/gradient-strings"
)

func main() {
	f, _ := gradient.NewGradientBuilder(gradient.WithColors(lipgloss.Color("#00FF00"), lipgloss.Color("#FF0000")),
		gradient.WithDomain(0.0, 1.0)).Build()

	b, _ := gradient.NewGradientBuilder(gradient.WithColors(lipgloss.Color("#FF0000"), lipgloss.Color("#00FF00")), gradient.WithDomain(0.0, 1.0)).Build()

	fmt.Println(b.RenderBackground(f.RenderForeground("This is Gradient String")))
}
