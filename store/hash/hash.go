package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct {
	Cost int
}

func New() *Hash {
	return &Hash{Cost: 10}
}

func (h *Hash) HashPassword(password string) (string, error) {
	b := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(b, h.Cost)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *Hash) VerifyPassword(stringPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(stringPassword))
	return err == nil
}
