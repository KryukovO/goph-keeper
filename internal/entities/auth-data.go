package entities

import "errors"

// ErrAuthDataAlreadyExists возвращается, если аналогичная пара логин/пароль уже существует.
var ErrAuthDataAlreadyExists = errors.New("data with the same login for this resource already exists")

// AuthData описывает пару логин/пароль пользователя.
type AuthData struct {
	UserID   int64  // UserID - идентификатор владельца.
	Resource string // Resource - ресурс в котором используются данные авторизации.
	Login    string // Login - логин.
	Password string // Password - пароль.
	Metadata string // Metadata - метаданные.
}
