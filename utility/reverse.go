package utility

// Reverse is a function to reverse a string
func Reverse(s string) string {

	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}

	return string(runes[n:])
}
