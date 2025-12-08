-- +goose Up
-- +goose StatementBegin

CREATE TABLE users_roles (
    user_id CHAR(26) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id CHAR(26) NOT NULL REFERENCES roles(id) ON DELETE RESTRICT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id, role_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_roles;
-- +goose StatementEnd