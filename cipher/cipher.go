package cipher

var keystore = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func TransformK(k int) int {
	// (a ∙ k + b) mod 52 where a = 250 and b = 479.
	output := (250*k + 479) % 52
	return output
}

func DecryptMessage(ciphertext string, k int) string {
	y := textToKey(ciphertext)
	x := (y - TransformK(k)) % 52
	if x < 0 {
		x = x + 52
	}
	return keyToText(x)
}

func EncryptMessage(plaintext string, k int) string {
	x := textToKey(plaintext)
	y := (x + TransformK(k)) % 52
	return keyToText(y)
}

func keyToText(key int) string {
	return string(keystore[key])
}

func textToKey(text string) int {
	for index, value := range keystore {
		if text == string(value) {
			return index
		}
	}
	return -1
}
