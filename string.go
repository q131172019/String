// Package string contains functions for working with strings.
package string

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
           j := len(b)-1-i
	   b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
