package cryptography

import "golang.org/x/crypto/bcrypt"

type hashComparer struct {
}

func NewHashComparer() *hashComparer {
	return &hashComparer{}
}

func (h *hashComparer) Compare(plaintext, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext))
	return err == nil
}
