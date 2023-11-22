// Package usecases описывает слой логики взамодействия с репозиторием.
package usecases

import "context"

// UseCases описывает интерфейс логики взамодействия с репозиторием.
type UseCases interface {
	// Registration выполняет регистрацию пользователя.
	Registration(ctx context.Context, login string, pswd string) (token string, err error)
	// Authorization выполняет авторизацию пользователя.
	Authorization(ctx context.Context, login string, pswd string) (token string, err error)
}
