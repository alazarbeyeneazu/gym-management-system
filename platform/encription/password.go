package encription

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password string) (string, error) {
	err := validation.Validate(password, validation.Required, validation.Length(8, 1000))
	if err != nil {
		return "", err
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		hashError := fmt.Errorf("can not hash the password %w", err)
		return "", hashError
	}
	return string(hasedPassword), nil
}

func CheckPassword(password, hasedpassword string) error {
	err := validation.Validate(password, validation.Required, validation.Length(8, 1000))
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(hasedpassword), []byte(password))

}
