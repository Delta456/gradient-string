package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/delta456/gradient-strings"
)

func main() {
	b, _ := gradient.NewGradientBuilder(gradient.WithColors(lipgloss.Color("123"), lipgloss.Color("202")),
		gradient.WithDomain(0.0, 1.0)).Build()

	b1, _ := gradient.NewGradientBuilder(gradient.WithColors(lipgloss.Color("253"), lipgloss.Color("129")),
		gradient.WithDomain(0.0, 1.0)).Build()

	fmt.Println(b1.RenderBackground(b.RenderForeground("Hello World")))
}
