-- +goose Up
-- +goose StatementBegin

CREATE TABLE namespaces (
    id CHAR(26) PRIMARY KEY NOT NULL,

    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Creating index for quick lookup by name (e.g., in API validation)
CREATE INDEX idx_namespaces_name ON namespaces (name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_namespaces_name;
DROP TABLE IF EXISTS namespaces;
-- +goose StatementEnd