package cryptography

import "golang.org/x/crypto/bcrypt"

type hasher struct{}

func NewHasher() *hasher {
	return &hasher{}
}

func (h *hasher) Hash(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), 14)
	return string(bytes), err
}
