package entities

// AuthData описывает пару логин/пароль пользователя.
type AuthData struct {
	UserID   int64  // UserID - идентификатор владельца.
	Resource string // Resource - ресурс в котором используются данные авторизации.
	Login    string // Login - логин.
	Password string // Password - пароль.
	Metadata string // Metadata - метаданные.
}
