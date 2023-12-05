package entities

import (
	"bytes"
	"errors"
)

// ErrFileIsTooBig возвращается, если размер файла превышает ограничение файлового хранилища.
var ErrFileIsTooBig = errors.New("file is too big")

// File описывает файл пользователя.
type File struct {
	UserID   int64        // UserID - идентификатор владельца файла.
	FileName string       // FileName - имя файла.
	Data     bytes.Buffer // Data - данные файла.
}
