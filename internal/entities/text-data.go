package entities

// TextData описывает текстовые данные пользователя.
type TextData struct {
	UserID   int64  // UserID - идентификатор владельца.
	Label    string // Label - заголовок данных.
	Data     string // Data - данные.
	Metadata string // Metadata - метаданные.
}
