package entities

import "errors"

var ErrBankDataAlreadyExists = errors.New("data with the same card number already exists")

// BankData описывает данные банковской карты пользователя.
type BankData struct {
	UserID         int64  // UserID - идентификатор владельца.
	Number         string // Number - номер карты.
	CardholderName string // CardholderName - имя владельца карты.
	ExpiredAt      string // ExpiredAt - дата окончания действия карты.
	CVV            string // CVV-код.
	Metadata       string // Metadata - метаданные.
}
