// Package entities содержит описание сущностей.
package entities

import (
	"errors"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/KryukovO/goph-keeper/pkg/utils"
)

const PswdEncCost = 14

var ErrInvalidLoginPassword = errors.New("invalid login/password")

// User описывает пользователя системы.
type User struct {
	// Login - логин пользователя.
	Login string
	// Password - пароль пользователя.
	Password string
	// EncryptedPassword - зашифрованный пароль пользователя.
	EncryptedPassword string
	// Salt - добавочная соль к паролю.
	Salt string
}

// Encrypt выполняет шифрование поля Password с добавлением соли.
func (user *User) Encrypt() error {
	if user.Salt == "" {
		salt, err := utils.GenerateRandomSalt(rand.NewSource(time.Now().UnixNano()))
		if err != nil {
			return err
		}

		user.Salt = salt
	}

	enc, err := bcrypt.GenerateFromPassword([]byte(user.Password+user.Salt), PswdEncCost)
	if err != nil {
		return err
	}

	user.EncryptedPassword = string(enc)

	return nil
}

// Validate возвращает ErrInvalidLoginPassword, если Password
// не соответствует EncryptedPassword с учетом Salt.
// Всегда nil, если EncryptedPassword не установлен.
func (user *User) Validate(secret []byte) error {
	if user.EncryptedPassword == "" {
		return nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(user.Password+user.Salt))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return ErrInvalidLoginPassword
	}

	if err != nil {
		return err
	}

	return nil
}
