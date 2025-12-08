-- +goose Up
-- +goose StatementBegin

-- Enable required extensions 
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users (
    id CHAR(26) PRIMARY KEY NOT NULL,

    username CITEXT NOT NULL,
    email CITEXT NOT NULL,

    password_hash TEXT NOT NULL,

    -- Whether the user has access to the system (virtual user)
    has_system_access BOOLEAN NOT NULL DEFAULT true,

    -- Whether the user has access to all namespaces
    has_all_namespaces_access BOOLEAN NOT NULL DEFAULT false,

    -- Whether the user is active (not disabled by admin)
    is_active BOOLEAN NOT NULL DEFAULT true,

    -- LOCKOUT STATE FIELDS:
    failed_attempts_count INTEGER NOT NULL DEFAULT 0,
    last_failed_attempt_at TIMESTAMPTZ,
    locked_until_at TIMESTAMPTZ,

    -- Soft deletion
    deleted_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT uq_users_username UNIQUE (username),
    CONSTRAINT uq_users_email UNIQUE (email)
);

-- Indexes
CREATE INDEX idx_users_username ON users (username);
CREATE INDEX idx_users_email ON users (email);

-- Soft-delete helper index
CREATE INDEX idx_users_not_deleted ON users (deleted_at)
    WHERE deleted_at IS NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS idx_users_not_deleted;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_username;
DROP TABLE IF EXISTS users;

-- +goose StatementEnd