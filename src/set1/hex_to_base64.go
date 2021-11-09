package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToBase64(hex_str string) string {
	hex_raw, err := hex.DecodeString(hex_str)
	if err != nil {
		log.Fatal("invalid hex string")
	}

	return base64.StdEncoding.EncodeToString(hex_raw)
}
