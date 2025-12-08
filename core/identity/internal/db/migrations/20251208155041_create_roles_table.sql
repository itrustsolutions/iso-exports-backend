-- +goose Up
-- +goose StatementBegin

CREATE TABLE roles (
    id CHAR(26) PRIMARY KEY NOT NULL,

    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,

    -- Indicates if the role is a system-defined role
    is_system_role BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd