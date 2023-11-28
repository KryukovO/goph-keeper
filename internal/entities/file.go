package entities

import "bytes"

// File описывает файл пользователя.
type File struct {
	UserID   int64        // UserID - идентификатор владельца файла.
	FileName string       // FileName - имя файла.
	Data     bytes.Buffer // Data - данные файла.
}
