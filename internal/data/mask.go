package data

// MaskWithStars replaces part of the text in the string with stars.
//
// The length of the input will be masked by all when it is less than or equal to 6.
func MaskWithStars(input string) string {
	chars := []rune(input)
	count := len(chars)
	for i := range count {
		if i != 0 && i != count-1 {
			chars[i] = '*'
		} else if i == 0 && count < 7 {
			chars[i] = '*'
		} else if i == count-1 && count < 8 {
			chars[i] = '*'
		}
	}

	return string(chars)
}
