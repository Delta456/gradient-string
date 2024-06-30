package gradient

import "regexp"

func rgbToAnsi(r, g, b float64) int {
	return 16 + 36*int(r*5) + 6*int(g*5) + int(b*5)
}

func RemoveANSICodes(input string) string {
	// ANSI escape code pattern
	ansiEscapePattern := `\x1b\[[0-9;]*m`
	re := regexp.MustCompile(ansiEscapePattern)
	return re.ReplaceAllString(input, "")
}
