package helper

// TextShortener takes string as an argument to shortened for fitting in the terminal
func TextShortener(text string) string {
	if text != "" {
		return text[0:100]
	}
	return text
}

/*
	it had written for fetching last-item
	not used in the curretn code
*/
