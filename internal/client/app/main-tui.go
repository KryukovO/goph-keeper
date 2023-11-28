package app

import (
	"github.com/rivo/tview"
)

func (a *App) setupMainMenu() {
	a.flex.Clear()
	a.list.Clear()

	a.list.AddItem("Пары логин/пароль", "", '1', func() {
		a.setupAuthDataMenu()
	})

	a.list.AddItem("Текстовые данные", "", '2', func() {
		a.setupTextDataMenu()
	})

	a.list.AddItem("Данные банковских карт", "", '3', func() {

	})

	a.list.AddItem("Файлы", "", '4', func() {

	})

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.list, 0, 1, true)
}
