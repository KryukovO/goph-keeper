// Package localstorage реализует локальное файловое хранилище.
package localstorage

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/pkg/utils"
)

// FileCatalog описывает каталог файлов пользователей: [userID][file name]size.
type FileCatalog map[int64]map[string]int64

// Size возвращает суммарный объем файлов пользователя.
func (fc FileCatalog) Size(userID int64) int64 {
	res := int64(0)

	files, ok := fc[userID]
	if !ok {
		return res
	}

	for _, size := range files {
		res += size
	}

	return res
}

// LocalStorage реализует лдокальное файловое хранилище.
type LocalStorage struct {
	mutex sync.RWMutex
	open  atomic.Bool // open - признак доступности хранилища.

	storeFolder   string                          // storeFolder - папка для хранения файлов.
	subscriptions map[int64]entities.Subscription // subscriptions - информация о подписках пользователей
	files         FileCatalog                     // files - файлы пользователя
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
		storeFolder:   path,
		subscriptions: make(map[int64]entities.Subscription),
		files:         make(map[int64]map[string]int64),
	}

	err := s.initStorage()
	if err != nil {
		return nil, err
	}

	s.open.Store(true)

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

	if !s.open.Load() {
		return nil
	}

	subscription := s.subscriptions[file.UserID]
	limit := subscriptionLimits[subscription]

	if s.files.Size(file.UserID)+int64(file.Data.Len()) > limit {
		return entities.ErrFileIsTooBig
	}

	folderPath := fmt.Sprintf("%s/%s", s.storeFolder, strconv.FormatInt(file.UserID, 10))

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, os.ModePerm); err != nil {
			return err
		}

		s.files[file.UserID] = make(map[string]int64)
	}

	err := utils.SaveFile(folderPath, file.FileName, file.Data)
	if err != nil {
		return err
	}

	s.files[file.UserID][file.FileName] = int64(file.Data.Len())

	return nil
}

// List возвращает список имен файлов пользователя в хранилище.
func (s *LocalStorage) List(userID int64) []string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.open.Load() {
		return []string{}
	}

	files, ok := s.files[userID]
	if !ok {
		return []string{}
	}

	res := make([]string, 0, len(files))

	for file := range files {
		res = append(res, file)
	}

	return res
}

// Load выгружает данные файла из хранилища в file по file.FileName и file.UserID.
func (s *LocalStorage) Load(file *entities.File) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.open.Load() {
		return nil
	}

	filePath := fmt.Sprintf(
		"%s/%s/%s",
		s.storeFolder, strconv.FormatInt(file.UserID, 10), file.FileName,
	)

	f, err := utils.GetFile(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buffer := bytes.NewBuffer(make([]byte, 0))
	part := make([]byte, 1024)

	for {
		var count int

		if count, err = reader.Read(part); err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		buffer.Write(part[:count])
	}

	file.Data = *buffer

	return nil
}

// Delete удаляет файл из хранилища.
func (s *LocalStorage) Delete(file entities.File) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.open.Load() {
		return nil
	}

	filePath := fmt.Sprintf(
		"%s/%s/%s",
		s.storeFolder, strconv.FormatInt(file.UserID, 10), file.FileName,
	)

	err := utils.RemoveFile(filePath)
	if err != nil {
		return err
	}

	delete(s.files[file.UserID], file.FileName)

	return nil
}

// SetSubscriptions устанавливает информацию о подписках пользователей.
func (s *LocalStorage) SetSubscriptions(subscriptions map[int64]entities.Subscription) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.open.Load() {
		return
	}

	s.subscriptions = subscriptions
}

// UpdateSubscription обновляет информацию о подписке пользователя.
func (s *LocalStorage) UpdateSubscription(userID int64, subscription entities.Subscription) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.open.Load() {
		return
	}

	s.subscriptions[userID] = subscription
}

// Close закрывает хранилище.
func (s *LocalStorage) Close() {
	s.open.Store(false)
}
