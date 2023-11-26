// Package filestorage описывает интерфейс файлового хранилища.
package filestorage

import "github.com/KryukovO/goph-keeper/internal/server/entities"

// FileStorage - интерфейс файлового хранилища.
type FileStorage interface {
	// Save выполняет сохранение файла в хранилище.
	Save(file entities.File) error
	// List возвращает список имен файлов пользователя в хранилище.
	List(userID int64) (files []string, err error)
	// Load выгружает данные файла из хранилища в file по file.FileName и file.UserName.
	Load(file *entities.File) error
	// Delete удаляет файл из хранилища.
	Delete(file *entities.File) error

	// SetSubscriptions устанавливает информацию о подписках пользователей.
	SetSubscriptions(subscriptions map[int64]entities.Subscription)
	// UpdateSubscription обновляет информацию о подписке пользователя.
	UpdateSubscription(userID int64, subscription entities.Subscription)
}
