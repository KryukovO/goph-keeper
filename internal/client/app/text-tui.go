package app

import (
	"context"
	"sort"

	"github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/rivo/tview"
)

func (a *App) setupTextDataMenu() {
	a.flex.Clear()
	a.list.Clear()

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
	defer cancel()

	resp, err := a.client.TextLabelsList(ctx, nil)
	if err != nil {
		a.logCh <- err.Error()

		a.setupMainMenu()
	}

	labels := resp.GetLabels()
	sort.Strings(labels)

	for index, label := range labels {
		a.list.AddItem(label, "", rune(49+index), nil)
	}

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.list, 0, 4, true).
		AddItem(tview.NewButton("Добавить").SetSelectedFunc(
			func() {
				a.setupTextDataForm(nil)

				a.pages.SwitchToPage(formPage)
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Изменить").SetSelectedFunc(
			func() {
				label := labels[a.list.GetCurrentItem()]

				ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
				defer cancel()

				resp, err := a.client.TextData(ctx, &serverpb.TextDataRequest{
					Label: label,
				})
				if err != nil {
					a.logCh <- err.Error()

					return
				}

				a.setupTextDataForm(&entities.TextData{
					Label:    resp.GetData().GetLabel(),
					Data:     resp.GetData().GetText(),
					Metadata: resp.GetData().GetMetadata(),
				})

				a.pages.SwitchToPage(formPage)
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Удалить").SetSelectedFunc(
			func() {
				label := labels[a.list.GetCurrentItem()]

				ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
				defer cancel()

				_, err := a.client.DeleteTextData(ctx, &serverpb.DeleteTextDataRequest{
					Label: label,
				})
				if err != nil {
					a.logCh <- err.Error()

					return
				}

				a.logCh <- "Данные удалены"

				a.setupTextDataMenu()
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Назад").SetSelectedFunc(
			func() {
				a.setupMainMenu()

				a.pages.SwitchToPage(menuPage)
			},
		), 0, 1, false)
}

func (a *App) setupTextDataForm(textData *entities.TextData) {
	add := textData == nil

	if add {
		textData = &entities.TextData{}
	}

	oldLabel := textData.Label

	a.flex.Clear()
	a.form.Clear(true)

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.form, 0, 1, true)

	a.form.AddInputField("Заголовок", textData.Label, 20, nil, func(label string) {
		textData.Label = label
	})

	a.form.AddTextArea("Текст", textData.Data, 40, 0, 0, func(data string) {
		textData.Data = data
	})

	a.form.AddTextArea("Примечание", textData.Metadata, 40, 0, 0, func(metadata string) {
		textData.Metadata = metadata
	})

	a.form.AddButton("Сохранить", func() {
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
		defer cancel()

		var err error

		if add {
			_, err = a.client.AddTextData(ctx, &serverpb.AddTextDataRequest{
				Data: &serverpb.TextData{
					Label:    textData.Label,
					Text:     textData.Data,
					Metadata: textData.Metadata,
				},
			})
		} else {
			_, err = a.client.UpdateTextData(ctx, &serverpb.UpdateTextDataRequest{
				OldLabel: oldLabel,
				Data: &serverpb.TextData{
					Label:    textData.Label,
					Text:     textData.Data,
					Metadata: textData.Metadata,
				},
			})
		}

		if err != nil {
			a.logCh <- err.Error()

			return
		}

		a.logCh <- "Данные сохранены"

		a.setupTextDataMenu()

		a.pages.SwitchToPage(menuPage)
	})

	a.form.AddButton("Назад", func() {
		a.setupTextDataMenu()
	})
}
