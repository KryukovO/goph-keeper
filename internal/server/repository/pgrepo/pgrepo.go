// Package pgstorage описывает репозиторий сервера, реализованный в PostgreSQL.
package pgrepo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/pkg/postgres"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// PgRepo - репозиторий PostgreSQL.
type PgRepo struct {
	db *postgres.Postgres
}

// NewPgRepo создаёт новый объект PgRepo.
func NewPgRepo(db *postgres.Postgres) *PgRepo {
	return &PgRepo{
		db: db,
	}
}

// CreateUser выполняет создание пользователя user.
func (repo *PgRepo) CreateUser(ctx context.Context, user entities.User) (int64, error) {
	query := `
		INSERT INTO users(login, password, salt, subscr) VALUES($1, $2, $3, $4)
		RETURNING id
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var id int64

	err = tx.QueryRowContext(
		ctx, query, user.Login, user.EncryptedPassword,
		user.Salt, user.Subscription,
	).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, entities.ErrUserAlreadyExists
		}

		return 0, err
	}

	return id, tx.Commit()
}

// User выполняет заполнение полей EncryptedPassword и Salt значениями из репозитория.
func (repo *PgRepo) User(ctx context.Context, user *entities.User) error {
	query := `
		SELECT 
			id, password, salt 
		FROM users
		WHERE login = $1
	`

	err := repo.db.QueryRowContext(ctx, query, user.Login).Scan(&user.ID, &user.EncryptedPassword, &user.Salt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.ErrInvalidLoginPassword
		}

		return err
	}

	return nil
}

// AddAuthData выполняет сохранение пары логин/пароль в репозитории.
func (repo *PgRepo) AddAuthData(ctx context.Context, data entities.AuthData) error {
	query := `
		INSERT INTO auth_data(
			user_id, resource, login, password, metadata)
			VALUES ($1, $2, $3, $4, $5)
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, query, data.UserID, data.Resource, data.Login, data.Password, data.Metadata)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return entities.ErrAuthDataAlreadyExists
		}

		return err
	}

	return tx.Commit()
}

// AddTextData выполняет сохранение текстовых данных в репозитории.
func (repo *PgRepo) AddTextData(ctx context.Context, data entities.TextData) error {
	query := `
		INSERT INTO text_data(
			user_id, label, data, metadata)
			VALUES ($1, $2, $3, $4)
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, query, data.UserID, data.Label, data.Data, data.Metadata)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return entities.ErrTextDataAlreadyExists
		}

		return err
	}

	return tx.Commit()
}

// AddBankData выполняет сохранение данных банковских карт в репозитории.
func (repo *PgRepo) AddBankData(ctx context.Context, data entities.BankData) error {
	query := `
		INSERT INTO bank_data(
			user_id, "number", name, expired_at, cvv, metadata)
			VALUES ($1, $2, $3, $4, $5, $6)
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(
		ctx, query, data.UserID, data.Number, data.CardholderName,
		data.ExpiredAt, data.CVV, data.Metadata,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return entities.ErrBankDataAlreadyExists
		}

		return err
	}

	return tx.Commit()
}

// UpdateAuthData выполняет обновление пары логин/пароль в репозитории.
func (repo *PgRepo) UpdateAuthData(ctx context.Context, oldResource, oldLogin string, data entities.AuthData) error {
	query := `
		UPDATE auth_data
			SET resource=$1, login=$2, password=$3, metadata=$4
		WHERE user_id = $5 AND resource = $6 AND login = $7
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(
		ctx, query, data.Resource, data.Login, data.Password, data.Metadata,
		data.UserID, oldResource, oldLogin,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return entities.ErrAuthDataAlreadyExists
		}

		return err
	}

	return tx.Commit()
}

// UpdateTextData выполняет обновление текстовых данных в репозитории.
func (repo *PgRepo) UpdateTextData(ctx context.Context, oldLabel string, data entities.TextData) error {
	query := `
		UPDATE text_data
			SET label=$1, data=$2, metadata=$3
		WHERE user_id = $4 AND label = $5
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, query, data.Label, data.Data, data.Metadata, data.UserID, oldLabel)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return entities.ErrTextDataAlreadyExists
		}

		return err
	}

	return tx.Commit()
}

// UpdateBankData выполняет обновление данных банковских карт в репозитории.
func (repo *PgRepo) UpdateBankData(ctx context.Context, oldNumber string, data entities.BankData) error {
	query := `
		UPDATE bank_data
			SET "number"=$1, name=$2, expired_at=$3, cvv=$4, metadata=$5
		WHERE user_id = $6 AND number = $7;
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(
		ctx, query, data.Number, data.CardholderName,
		data.ExpiredAt, data.CVV, data.Metadata,
		data.UserID, oldNumber,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return entities.ErrBankDataAlreadyExists
		}

		return err
	}

	return tx.Commit()
}

// DeleteAuthData выполняет удаление пары логин/пароль из репозитория.
func (repo *PgRepo) DeleteAuthData(ctx context.Context, data entities.AuthData) error {
	query := `
		DELETE FROM auth_data
		WHERE user_id = $1 AND resource = $2 AND login = $3;
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, query, data.UserID, data.Resource, data.Login)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// DeleteTextData выполняет удаление текстовых данных из репозитория.
func (repo *PgRepo) DeleteTextData(ctx context.Context, data entities.TextData) error {
	query := `
		DELETE FROM text_data
		WHERE user_id = $1 AND label = $2;
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, query, data.UserID, data.Label)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// DeleteBankData выполняет удаление данных банковских карт из репозитория.
func (repo *PgRepo) DeleteBankData(ctx context.Context, data entities.BankData) error {
	query := `
		DELETE FROM bank_data
		WHERE user_id = $1 AND number = $2;
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, query, data.UserID, data.Number)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// AuthDataList возвращает список сохраненных пар логин/пароль из репозитория.
func (repo *PgRepo) AuthDataList(ctx context.Context, userID int64) ([]entities.AuthData, error) {
	query := `
		SELECT user_id, resource, login, password, metadata
		FROM auth_data
		WHERE user_id = $1
	`

	res := make([]entities.AuthData, 0)

	rows, err := repo.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.AuthData

		err = rows.Scan(
			&data.UserID, &data.Resource, &data.Login,
			&data.Password, &data.Metadata,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// TextLabelsList возвращает список заголовков сохранённых текстовых данных из репозитория.
func (repo *PgRepo) TextLabelsList(ctx context.Context, userID int64) ([]string, error) {
	query := `
		SELECT label
		FROM text_data
		WHERE user_id = $1
	`

	res := make([]string, 0)

	rows, err := repo.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var label string

		err = rows.Scan(&label)
		if err != nil {
			return nil, err
		}

		res = append(res, label)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// TextData возвращает сохранённые текстовые данные по заголовку из репозитория.
func (repo *PgRepo) TextData(ctx context.Context, data *entities.TextData) error {
	query := `
		SELECT data, metadata
		FROM text_data
		WHERE user_id = $1 AND label = $2
	`

	err := repo.db.QueryRowContext(ctx, query, data.UserID, data.Label).Scan(
		&data.Data, &data.Metadata,
	)

	return err
}

// BankCardList возвращает список номеров банковских карт из репозитория.
func (repo *PgRepo) BankCardNumbersList(ctx context.Context, userID int64) ([]string, error) {
	query := `
		SELECT "number"
		FROM bank_data
		WHERE user_id = $1
	`

	res := make([]string, 0)

	rows, err := repo.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var number string

		err = rows.Scan(&number)
		if err != nil {
			return nil, err
		}

		res = append(res, number)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// BankCardList возвращает данные банковской карты по номеру из репозитория.
func (repo *PgRepo) BankCard(ctx context.Context, data *entities.BankData) error {
	query := `
		SELECT name, expired_at, cvv, metadata
		FROM bank_data
		WHERE user_id = $1 AND number = $2
	`

	err := repo.db.QueryRowContext(ctx, query, data.UserID, data.Number).Scan(
		&data.CardholderName, &data.ExpiredAt, &data.CVV, &data.Metadata,
	)

	return err
}

// Sunbsriptions возвращает информацию о подписке пользователей из репозитория.
func (repo *PgRepo) Sunbsriptions(ctx context.Context) (map[int64]entities.Subscription, error) {
	query := `
		SELECT id, subscr FROM users
	`

	res := make(map[int64]entities.Subscription, 0)

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			userID       int64
			subscription string
		)

		err = rows.Scan(&userID, &subscription)
		if err != nil {
			return nil, err
		}

		res[userID] = entities.Subscription(subscription)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
