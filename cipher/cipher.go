package cipher

var keystore = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// var keystore = map[int]string{
// 	0: "A",
// 	1: "B",
// 	2: "C",
// 	3: "D",
// 	4: "E",
// 	5: "F",
// 	6: "G",
// 	7: "H",
// 	8: "I",
// 	9: "J",
// 	10: "K",
// 	11: "L",
// 	12: "M",
// 	13: "N",
// 	14: "O",
// 	15: "P",
// 	16: ""

// }

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

// fromKey returns number mapping for char.
// if the key is invalid, "" is returned.
func textToKey(text string) int {
	for index, value := range keystore {
		if text == string(value) {
			return index
		}
	}
	return -1
}

// fromKey returns number mapping for char.
// if the char is invalid, -1 is returned.
// func toKey(char string) int {
// 	for index, value := range keystore {
// 		if char == string(value) {
// 			return index
// 		}
// 	}
// 	return -1
// }
