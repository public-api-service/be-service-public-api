package helper

func ToAlphaString(col int) string {
	var result string
	for col > 0 {
		col--
		result = string('A'+col%26) + result
		col /= 26
	}
	return result
}
