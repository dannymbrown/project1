package cipher

func TransformK(k int) int {
	// (a ∙ k + b) mod 52 where a = 250 and b = 479.
	output := (250*k + 479) % 52
	return output
}

func FromCipher(char string) string {
	return char
}

func ToCipher(char string) string {
	return char
}
