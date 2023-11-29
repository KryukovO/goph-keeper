package app

import (
	"context"
	"sort"

	"github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/rivo/tview"
)

func (a *App) setupBankDataMenu() {
	a.flex.Clear()
	a.list.Clear()

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
	defer cancel()

	resp, err := a.client.BankCardNumbersList(ctx, nil)
	if err != nil {
		a.logCh <- err.Error()

		a.setupMainMenu()
	}

	numbers := resp.GetCardNumbers()
	sort.Strings(numbers)

	for index, number := range numbers {
		a.list.AddItem(number, "", rune(49+index), nil)
	}

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.list, 0, 4, true).
		AddItem(tview.NewButton("Добавить").SetSelectedFunc(
			func() {
				a.setupBankDataForm(nil)

				a.pages.SwitchToPage(formPage)
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Изменить").SetSelectedFunc(
			func() {
				number := numbers[a.list.GetCurrentItem()]

				ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
				defer cancel()

				resp, err := a.client.BankCard(ctx, &serverpb.BankCardRequest{
					CardNumber: number,
				})
				if err != nil {
					a.logCh <- err.Error()

					return
				}

				a.setupBankDataForm(&entities.BankData{
					Number:         resp.GetData().GetNumber(),
					CardholderName: resp.GetData().GetCardholderName(),
					ExpiredAt:      resp.GetData().GetExpirationDate(),
					CVV:            resp.GetData().GetCVV(),
					Metadata:       resp.GetData().GetMetadata(),
				})

				a.pages.SwitchToPage(formPage)
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Удалить").SetSelectedFunc(
			func() {
				number := numbers[a.list.GetCurrentItem()]

				ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
				defer cancel()

				_, err := a.client.DeleteBankData(ctx, &serverpb.DeleteBankDataRequest{
					Number: number,
				})
				if err != nil {
					a.logCh <- err.Error()

					return
				}

				a.logCh <- "Данные удалены"

				a.setupAuthDataMenu()
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Назад").SetSelectedFunc(
			func() {
				a.setupMainMenu()
			},
		), 0, 1, false)
}

func (a *App) setupBankDataForm(bankData *entities.BankData) {
	add := bankData == nil

	if add {
		bankData = &entities.BankData{}
	}

	oldNumber := bankData.Number

	a.flex.Clear()
	a.form.Clear(true)

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.form, 0, 1, true)

	a.form.AddInputField("Номер карты", bankData.Number, 20, nil, func(number string) {
		bankData.Number = number
	})

	a.form.AddInputField("ФИО держателя", bankData.CardholderName, 20, nil, func(name string) {
		bankData.CardholderName = name
	})

	a.form.AddInputField("Дата окончания действия", bankData.ExpiredAt, 20, nil, func(expired string) {
		bankData.ExpiredAt = expired
	})

	a.form.AddInputField("CVV", bankData.CVV, 20, nil, func(cvv string) {
		bankData.CVV = cvv
	})

	a.form.AddTextArea("Примечание", bankData.Metadata, 40, 0, 0, func(metadata string) {
		bankData.Metadata = metadata
	})

	a.form.AddButton("Сохранить", func() {
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
		defer cancel()

		var err error

		if add {
			_, err = a.client.AddBankData(ctx, &serverpb.AddBankDataRequest{
				Data: &serverpb.BankData{
					Number:         bankData.Number,
					CardholderName: bankData.CardholderName,
					ExpirationDate: bankData.ExpiredAt,
					CVV:            bankData.CVV,
					Metadata:       bankData.Metadata,
				},
			})
		} else {
			_, err = a.client.UpdateBankData(ctx, &serverpb.UpdateBankDataRequest{
				OldNumber: oldNumber,
				Data: &serverpb.BankData{
					Number:         bankData.Number,
					CardholderName: bankData.CardholderName,
					ExpirationDate: bankData.ExpiredAt,
					CVV:            bankData.CVV,
					Metadata:       bankData.Metadata,
				},
			})
		}

		if err != nil {
			a.logCh <- err.Error()

			return
		}

		a.logCh <- "Данные сохранены"

		a.setupMainMenu()

		a.pages.SwitchToPage(menuPage)
	})

	a.form.AddButton("Отмена", func() {
		a.setupMainMenu()

		a.pages.SwitchToPage(menuPage)
	})
}
