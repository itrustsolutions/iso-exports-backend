-- +goose Up
-- +goose StatementBegin

CREATE TABLE roles_permissions (
    role_id CHAR(26) NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id CHAR(26) NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (role_id, permission_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles_permissions;
-- +goose StatementEnd