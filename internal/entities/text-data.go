package entities

import "errors"

var ErrTextDataAlreadyExists = errors.New("data with the same label already exists")

// TextData описывает текстовые данные пользователя.
type TextData struct {
	UserID   int64  // UserID - идентификатор владельца.
	Label    string // Label - заголовок данных.
	Data     string // Data - данные.
	Metadata string // Metadata - метаданные.
}
