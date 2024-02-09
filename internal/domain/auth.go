package domain

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func generateSalt() (salt []byte, err error) {
	salt = make([]byte, 16)
	_, err = rand.Read(salt)
	return
}

func hashPassword(password []byte, salt []byte) string {
	key := argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
	key = append(key, salt...)
	return base64.StdEncoding.EncodeToString(key)
}

func verifyPassword(password []byte, encodedPassword string) bool {
	decodedPassword, err := base64.StdEncoding.DecodeString(encodedPassword)
	if err != nil {
		return false
	}
	keyExpected := decodedPassword[:32]
	salt := decodedPassword[32:]
	key := argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
	return bytes.Equal(key, keyExpected)
}
