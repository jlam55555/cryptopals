package set1

import (
	"crypto/aes"
	"log"
)

func AesDecode(ciphertext, key []byte) string {
	// generate aes cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// buffer for decrypted
	plaintext := make([]byte, len(ciphertext))

	// block size
	bs := len(key)

	// perform ECB block by block
	for start, end := 0, bs; start < len(ciphertext); start, end = start+bs, end+bs {
		c.Decrypt(plaintext[start:end], ciphertext[start:end])
	}

	return string(plaintext)
}
