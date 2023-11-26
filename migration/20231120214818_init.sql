-- +goose Up
-- +goose StatementBegin
    CREATE TYPE "subscription" AS ENUM ('UNKNOWN', 'REGULAR', 'PREMIUM');

    CREATE TABLE IF NOT EXISTS "users" (
        "id" BIGINT GENERATED ALWAYS AS IDENTITY,
        "login" TEXT NOT NULL,
        "password" TEXT NOT NULL,
        "salt" TEXT NOT NULL,
        "subscr" "subscription" NOT NULL,
        PRIMARY KEY(id),
        CONSTRAINT "users_uniq" UNIQUE("login")
    );

    CREATE TABLE IF NOT EXISTS "auth_data" (
        "id" BIGINT GENERATED ALWAYS AS IDENTITY,
        "user_id" BIGINT NOT NULL,
        "resource" TEXT NOT NULL,
        "login" TEXT NOT NULL,
        "password" TEXT NOT NULL,
        "metadata" TEXT,
        PRIMARY KEY(id),
        CONSTRAINT "auth_data_uniq" UNIQUE("user_id", "resource", "login")
    );
    
    CREATE TABLE IF NOT EXISTS "text_data" (
        "id" BIGINT GENERATED ALWAYS AS IDENTITY,
        "user_id" BIGINT NOT NULL,
        "label" TEXT NOT NULL,
        "data" TEXT NOT NULL,
        "metadata" TEXT,
        PRIMARY KEY(id),
        CONSTRAINT "text_data_uniq" UNIQUE("user_id", "label")
    );
    
    CREATE TABLE IF NOT EXISTS "bank_data" (
        "id" BIGINT GENERATED ALWAYS AS IDENTITY,
        "user_id" BIGINT NOT NULL,
        "number" TEXT NOT NULL,
        "name" TEXT NOT NULL,
        "expired_at" TEXT NOT NULL,
        "cvv" TEXT NOT NULL,
        "metadata" TEXT,
        PRIMARY KEY(id),
        CONSTRAINT "bank_data_uniq" UNIQUE("user_id", "number")
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE IF EXISTS "users";
    DROP TABLE IF EXISTS "auth_data";
    DROP TABLE IF EXISTS "text_data";
    DROP TABLE IF EXISTS "bank_data";
    DROP TYPE "subscription";
-- +goose StatementEnd
