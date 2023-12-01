package main

import "fmt"

// ROT13 Cipher Encryption Function
func rot13Encrypt(plainText string) string {
	cipherText := ""

	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			cipherText += string((char-'A'+13)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			cipherText += string((char-'a'+13)%26 + 'a')
		} else {
			cipherText += string(char)
		}
	}

	return cipherText
}

// ROT13 Cipher Decryption Function
func rot13Decrypt(cipherText string) string {
	plainText := ""

	for _, char := range cipherText {
		if char >= 'A' && char <= 'Z' {
			plainText += string((char-'A'+13)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			plainText += string((char-'a'+13)%26 + 'a')
		} else {
			plainText += string(char)
		}
	}

	return plainText
}

func main() {
	plainText := "HELLO WORLD"

	cipherText := rot13Encrypt(plainText)

	fmt.Printf("Plaintext: %s\n", plainText)
	fmt.Printf("Ciphertext: %s\n", cipherText)

	plainText = rot13Decrypt(cipherText)
	fmt.Printf("Plaintext: %s\n", plainText)
}
