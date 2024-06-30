package gradient

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/mazznoer/colorgrad"
)

const resetANSI = "\033[0m"

type GradientDirection int

const (
	Horizontal GradientDirection = iota
	Vertical
)

type GradientBuilder struct {
	builder   *colorgrad.GradientBuilder
	grad      colorgrad.Gradient
	direction GradientDirection
}

type Option func(*GradientBuilder)

func NewGradientBuilder(options ...Option) *GradientBuilder {
	grad := colorgrad.NewGradient()
	builder := &GradientBuilder{builder: grad, direction: Horizontal}

	for _, option := range options {
		option(builder)
	}

	return builder
}

func WithColors(colors ...color.Color) Option {
	return func(b *GradientBuilder) {
		b.builder = b.builder.Colors(colors...)
	}
}

func WithHTMLColors(colors ...string) Option {
	return func(b *GradientBuilder) {
		b.builder = b.builder.HtmlColors(colors...)
	}
}

func WithDomain(domain ...float64) Option {
	return func(b *GradientBuilder) {
		b.builder = b.builder.Domain(domain...)
	}
}

func WithDirection(direction GradientDirection) Option {
	return func(b *GradientBuilder) {
		b.direction = direction
	}
}

func (b *GradientBuilder) Build() (*GradientBuilder, error) {
	grad, err := b.builder.Build()
	if err != nil {
		return &GradientBuilder{}, err
	}

	b.grad = grad
	return b, nil
}

func (b *GradientBuilder) RenderForeground(s string) string {
	runes := []rune(RemoveANSICodes(s))
	length := len(runes)
	var result strings.Builder
	for i, c := range runes {
		var t float64
		if b.direction == Horizontal {
			t = float64(i) / float64(length-1)
		} else { // Vertical
			t = float64(length-i-1) / float64(length-1)
		}
		color := b.grad.At(t)
		r, g, blue := color.RGB255()
		ansiColor := rgbToAnsi(float64(r)/255, float64(g)/255, float64(blue)/255)
		result.WriteString(fmt.Sprintf("\033[38;5;%dm%c", ansiColor, c))
	}
	result.WriteString(resetANSI) // Reset color
	return result.String()
}

func (b *GradientBuilder) RenderBackground(s string) string {
	runes := []rune(RemoveANSICodes(s))
	length := len(runes)
	var result strings.Builder
	for i, c := range runes {
		var t float64
		if b.direction == Horizontal {
			t = float64(i) / float64(length-1)
		} else { // Vertical
			t = float64(length-i-1) / float64(length-1)
		}
		color := b.grad.At(t)
		r, g, blue := color.RGB255()
		ansiColor := rgbToAnsi(float64(r)/255, float64(g)/255, float64(blue)/255)
		result.WriteString(fmt.Sprintf("\033[48;5;%dm%c", ansiColor, c))
	}
	result.WriteString(resetANSI) // Reset color
	return result.String()
}

/*

 */
