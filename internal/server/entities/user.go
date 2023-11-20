// Package entities содержит описание сущностей.
package entities

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"github.com/KryukovO/goph-keeper/pkg/utils"
)

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

// Encrypt выполняет шифрование SHA-256 поля Password с добавлением соли.
func (user *User) Encrypt(secret []byte) error {
	if user.Salt == "" {
		salt, err := utils.GenerateRandomSalt(rand.NewSource(time.Now().UnixNano()))
		if err != nil {
			return err
		}

		user.Salt = salt
	}

	enc := hmac.New(sha256.New, secret)

	_, err := enc.Write([]byte(user.Password + user.Salt))
	if err != nil {
		return err
	}

	user.Password = hex.EncodeToString(enc.Sum(nil))

	return nil
}

// Validate возвращает ErrInvalidLoginPassword, если результат SHA-256 шифрования Password
// не соответствует EncryptedPassword с учетом Salt и secret.
// Всегда nil, если EncryptedPassword не установлен.
func (user *User) Validate(secret []byte) error {
	if user.EncryptedPassword == "" {
		return nil
	}

	enc := hmac.New(sha256.New, secret)

	_, err := enc.Write([]byte(user.Password + user.Salt))
	if err != nil {
		return err
	}

	hash := hex.EncodeToString(enc.Sum(nil))

	if user.EncryptedPassword != hash {
		return ErrInvalidLoginPassword
	}

	return nil
}
