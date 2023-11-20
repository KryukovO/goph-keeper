-- +goose Up
-- +goose StatementBegin
    CREATE TABLE IF NOT EXISTS "users" (
        id BIGINT GENERATED ALWAYS AS IDENTITY,
        login TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        salt TEXT NOT NULL,
        PRIMARY KEY(id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
