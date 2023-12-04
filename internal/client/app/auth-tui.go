package app

import (
	"context"
	"fmt"
	"sort"

	"github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/rivo/tview"
)

func (a *App) setupAuthDataMenu() {
	a.flex.Clear()
	a.list.Clear()

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
	defer cancel()

	resp, err := a.client.AuthDataList(ctx, &empty.Empty{})
	if err != nil {
		a.logCh <- err.Error()

		a.setupMainMenu()

		return
	}

	data := resp.GetData()
	sort.Slice(data, func(i, j int) bool {
		return data[i].GetResource() < data[j].GetLogin() ||
			!(data[i].GetResource() < data[j].GetLogin()) &&
				data[i].GetLogin() < data[j].GetLogin()
	})

	for index, item := range data {
		a.list.AddItem(
			fmt.Sprintf("%s: %s", item.GetResource(), item.GetLogin()),
			"",
			rune(49+index),
			nil,
		)
	}

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.list, 0, 4, true).
		AddItem(tview.NewButton("Добавить").SetSelectedFunc(
			func() {
				a.setupAuthDataForm(nil)

				a.pages.SwitchToPage(formPage)
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Изменить").SetSelectedFunc(
			func() {
				if len(data) < a.list.GetCurrentItem()+1 {
					return
				}

				item := data[a.list.GetCurrentItem()]

				a.setupAuthDataForm(&entities.AuthData{
					Resource: item.GetResource(),
					Login:    item.GetLogin(),
					Password: item.GetUserPassword(),
					Metadata: item.GetMetadata(),
				})

				a.pages.SwitchToPage(formPage)
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Удалить").SetSelectedFunc(
			func() {
				if len(data) < a.list.GetCurrentItem()+1 {
					return
				}

				item := data[a.list.GetCurrentItem()]

				ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
				defer cancel()

				_, err := a.client.DeleteAuthData(ctx, &serverpb.DeleteAuthDataRequest{
					Resource: item.GetResource(),
					Login:    item.GetLogin(),
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

func (a *App) setupAuthDataForm(authData *entities.AuthData) {
	add := authData == nil

	if add {
		authData = &entities.AuthData{}
	}

	oldResource := authData.Resource
	oldLogin := authData.Login

	a.flex.Clear()
	a.form.Clear(true)

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.form, 0, 1, true)

	a.form.AddInputField("Ресурс", authData.Resource, 20, nil, func(resource string) {
		authData.Resource = resource
	})

	a.form.AddInputField("Логин", authData.Login, 20, nil, func(login string) {
		authData.Login = login
	})

	a.form.AddInputField("Пароль", authData.Password, 20, nil, func(password string) {
		authData.Password = password
	})

	a.form.AddTextArea("Примечание", authData.Metadata, 40, 0, 0, func(metadata string) {
		authData.Metadata = metadata
	})

	a.form.AddButton("Сохранить", func() {
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
		defer cancel()

		var err error

		if add {
			_, err = a.client.AddAuthData(ctx, &serverpb.AddAuthDataRequest{
				Data: &serverpb.AuthData{
					Resource:     authData.Resource,
					Login:        authData.Login,
					UserPassword: authData.Password,
					Metadata:     authData.Metadata,
				},
			})
		} else {
			_, err = a.client.UpdateAuthData(ctx, &serverpb.UpdateAuthDataRequest{
				OldResource: oldResource,
				OldLogin:    oldLogin,
				Data: &serverpb.AuthData{
					Resource:     authData.Resource,
					Login:        authData.Login,
					UserPassword: authData.Password,
					Metadata:     authData.Metadata,
				},
			})
		}

		if err != nil {
			a.logCh <- err.Error()

			return
		}

		a.logCh <- "Данные сохранены"

		a.setupAuthDataMenu()

		a.pages.SwitchToPage(menuPage)
	})

	a.form.AddButton("Назад", func() {
		a.setupAuthDataMenu()

		a.pages.SwitchToPage(menuPage)
	})
}
