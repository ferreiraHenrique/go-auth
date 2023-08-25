package domain

import "golang.org/x/crypto/bcrypt"

type Password struct {
	Password string
}

func NewPassword(pass string) *Password {
	if pass == "" {
		return nil
	}

	password := Password{Password: pass}

	return &password
}

func (password *Password) Hash() error {
	bytePlain := []byte(password.Password)
	hash, err := bcrypt.GenerateFromPassword(bytePlain, bcrypt.MinCost)
	if err != nil {
		return err
	}

	password.Password = string(hash)
	return nil
}

func (password *Password) Compare(plain string) bool {
	byteHash := []byte(password.Password)
	bytePlain := []byte(plain)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	return err == nil
}
