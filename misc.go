package main

import (
	"crypto/sha512"
	"encoding/hex"
)

func EncryptData(data string) string {
	EncrypHash := sha512.New512_256()
	EncrypData := EncrypHash.Sum([]byte(data))
	return hex.EncodeToString(EncrypData)
}
