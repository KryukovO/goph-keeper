package app

import (
	"context"
	"fmt"

	"github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/rivo/tview"
)

var subscriptions = []string{
	string(entities.RegularSubscription),
	string(entities.PremiumSubscription),
}

func (a *App) setupWelcomeMenu() {
	a.flex.Clear()
	a.list.Clear()

	a.list.AddItem("Регистрация", "", '1', func() {
		a.setupRegistrationForm()
		a.pages.SwitchToPage(formPage)
	})

	a.list.AddItem("Авторизация", "", '2', func() {
		a.setupLoginForm()
		a.pages.SwitchToPage(formPage)
	})

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.list, 0, 1, true)
}

func (a *App) setupLoginForm() {
	var bufUser entities.User

	a.flex.Clear()
	a.form.Clear(true)

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.form, 0, 1, true)

	a.form.AddInputField("Логин", "", 20, nil, func(login string) {
		bufUser.Login = login
	})

	a.form.AddPasswordField("Пароль", "", 20, '*', func(password string) {
		bufUser.Password = password
	})

	a.form.AddButton("Войти", func() {
		a.user = bufUser

		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
		defer cancel()

		resp, err := a.client.Authorization(ctx, &serverpb.AuthorizationRequest{
			Login:    a.user.Login,
			Password: a.user.Password,
		})
		if err != nil {
			a.logCh <- err.Error()

			return
		}

		a.token = resp.GetToken()

		a.setupMainMenu()

		a.pages.SwitchToPage(menuPage)
	})

	a.form.AddButton("Назад", func() {
		a.setupWelcomeMenu()

		a.pages.SwitchToPage(menuPage)
	})
}

func (a *App) setupRegistrationForm() {
	var bufUser entities.User

	a.flex.Clear()
	a.form.Clear(true)

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.form, 0, 1, true)

	a.form.AddInputField("Логин", "", 20, nil, func(login string) {
		bufUser.Login = login
	})

	a.form.AddPasswordField("Пароль", "", 20, '*', func(password string) {
		bufUser.Password = password
	})

	a.form.AddDropDown("Подписка", subscriptions, 0, func(subscription string, _ int) {
		bufUser.Subscription = entities.Subscription(subscription)
	})

	a.form.AddButton("Сохранить", func() {
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
		defer cancel()

		resp, err := a.client.Registration(ctx, &serverpb.RegistrationRequest{
			Login:        bufUser.Login,
			Password:     bufUser.Password,
			Subscription: entities.ConvertSubscription(bufUser.Subscription),
		})
		if err != nil {
			a.logCh <- err.Error()

			return
		}

		a.logCh <- fmt.Sprintf("Авторизован: %s", bufUser.Login)

		a.token = resp.GetToken()
		a.user = bufUser

		a.setupMainMenu()

		a.pages.SwitchToPage(menuPage)
	})

	a.form.AddButton("Назад", func() {
		a.setupWelcomeMenu()

		a.pages.SwitchToPage(menuPage)
	})
}
