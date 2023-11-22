-- +goose Up
-- +goose StatementBegin
    CREATE TABLE IF NOT EXISTS "users" (
        "id" BIGINT GENERATED ALWAYS AS IDENTITY,
        "login" TEXT NOT NULL UNIQUE,
        "password" TEXT NOT NULL,
        "salt" TEXT NOT NULL,
        PRIMARY KEY(id)
    );

    CREATE TABLE IF NOT EXISTS "auth_data" (
        "id" BIGINT GENERATED ALWAYS AS IDENTITY,
        "user_id" BIGINT,
        "resource" TEXT,
        "login" TEXT,
        "password" TEXT,
        "metadata" TEXT,
        CONSTRAINT "auth_data_uniq" PRIMARY KEY("user_id", "resource", "login")
    );
    
    CREATE TABLE IF NOT EXISTS "text_data" (
        "id" BIGINT GENERATED ALWAYS AS IDENTITY,
        "user_id" BIGINT,
        "label" TEXT,
        "data" TEXT,
        "metadata" TEXT,
        CONSTRAINT "text_data_uniq" PRIMARY KEY("user_id", "label")
    );
    
    CREATE TABLE IF NOT EXISTS "bank_data" (
        "id" BIGINT GENERATED ALWAYS AS IDENTITY,
        "user_id" BIGINT,
        "number" TEXT,
        "name" TEXT,
        "expired_at" TEXT,
        "cvv" TEXT,
        "metadata" TEXT,
        CONSTRAINT "bank_data_uniq" PRIMARY KEY("user_id", "number")
    );

    CREATE TABLE IF NOT EXISTS "file_data" (
        "id" BIGINT GENERATED ALWAYS AS IDENTITY,
        "user_id" BIGINT,
        "filename" TEXT,
        "data" BYTEA,
        "metadata" TEXT,
        CONSTRAINT "file_data_uniq" PRIMARY KEY("user_id", "filename")
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE IF EXISTS "users";
    DROP TABLE IF EXISTS "auth_data";
    DROP TABLE IF EXISTS "text_data";
    DROP TABLE IF EXISTS "bank_data";
    DROP TABLE IF EXISTS "file_data";
-- +goose StatementEnd
