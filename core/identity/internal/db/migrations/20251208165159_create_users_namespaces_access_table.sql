-- +goose Up
-- +goose StatementBegin

CREATE TABLE users_namespaces_access (
    user_id CHAR(26) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    namespace_id CHAR(26) NOT NULL REFERENCES namespaces(id) ON DELETE CASCADE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id, namespace_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_namespaces_access;
-- +goose StatementEnd