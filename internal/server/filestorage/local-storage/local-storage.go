package localstorage

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/KryukovO/goph-keeper/internal/entities"
)

// LocalStorage реализует лдокальное файловое хранилище.
type LocalStorage struct {
	mutex sync.RWMutex
	open  bool // open - признак доступности хранилища.

	storeFolder   string                          // storeFolder - папка для хранения файлов.
	subscriptions map[int64]entities.Subscription // subscriptions - информация о подписках пользователей
	files         map[int64]map[string]int64      // files - файлы пользователя: [userID][file name]size
}

// subscriptionLimits хранит ограничения объема хранилища для подписок.
var subscriptionLimits = map[entities.Subscription]int64{
	entities.UnknownSubscription: 0,                      // 0 bytes
	entities.RegularSubscription: 10 * 1024 * 1024,       // 10 MB
	entities.PremiumSubscription: 1 * 1024 * 1024 * 1024, // 1 GB
}

// NewLocalStorage возвращает новый объект LocalStorage.
func NewLocalStorage(path string) (*LocalStorage, error) {
	s := &LocalStorage{
		open:          true,
		storeFolder:   path,
		subscriptions: make(map[int64]entities.Subscription),
		files:         make(map[int64]map[string]int64),
	}

	err := s.initStorage()
	if err != nil {
		return nil, err
	}

	return s, nil
}

// initStorage инициализирует локальное хранилище.
func (s *LocalStorage) initStorage() error {
	entries, err := os.ReadDir(s.storeFolder)
	if err != nil {
		return err
	}

	for _, e := range entries {
		// Предполагаем, что в указанной директории есть файлы,
		// их не учитываем
		if !e.IsDir() {
			continue
		}

		userID, err := strconv.ParseInt(e.Name(), 10, 64)
		// Поддиректории, не являющиеся идентификаторами, пропускаем
		if err != nil {
			continue
		}

		s.files[userID] = make(map[string]int64)

		subEntries, err := os.ReadDir(fmt.Sprintf("%s/%s", s.storeFolder, e.Name()))
		if err != nil {
			return err
		}

		for _, se := range subEntries {
			// Пропускаем вложенные директории.
			if se.IsDir() {
				continue
			}

			info, err := se.Info()
			if err != nil {
				return err
			}

			s.files[userID][se.Name()] += info.Size()
		}
	}

	return nil
}

// Save выполняет сохранение файла в хранилище.
func (s *LocalStorage) Save(file entities.File) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.open {
		return nil
	}

	return nil
}

// List возвращает список имен файлов пользователя в хранилище.
func (s *LocalStorage) List(userID int64) ([]string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.open {
		return []string{}, nil
	}

	return []string{}, nil
}

// Load выгружает данные файла из хранилища в file по file.FileName и file.UserID.
func (s *LocalStorage) Load(file *entities.File) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.open {
		return nil
	}

	return nil
}

// Delete удаляет файл из хранилища.
func (s *LocalStorage) Delete(file *entities.File) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.open {
		return nil
	}

	return nil
}

// SetSubscriptions устанавливает информацию о подписках пользователей.
func (s *LocalStorage) SetSubscriptions(subscriptions map[int64]entities.Subscription) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.open {
		return
	}

	s.subscriptions = subscriptions
}

// UpdateSubscription обновляет информацию о подписке пользователя.
func (s *LocalStorage) UpdateSubscription(userID int64, subscription entities.Subscription) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.open {
		return
	}

	s.subscriptions[userID] = subscription
}

// Close закрывает хранилище.
func (s *LocalStorage) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.open = false
}
