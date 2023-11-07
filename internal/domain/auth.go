package domain

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func generateSalt() (salt []byte, err error) {
	salt = make([]byte, 16)
	_, err = rand.Read(salt)
	return
}

func hashPassword(password []byte, salt []byte) []byte {
	key := argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
	key = append(key, salt...)
	return key
}

func verifyPassword(password []byte, encodedPassword []byte) bool {
	keyExpected := encodedPassword[:32]
	salt := encodedPassword[32:]
	key := argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
	return bytes.Equal(key, keyExpected)
}
