package entities

// File описывает файл пользователя.
type File struct {
	UserID   int64  // UserID - идентификатор владельца файла.
	FileName string // FileName - имя файла.
	Data     []byte // Data - данные файла.
	Metadata string // Metadata - метаданные
}
