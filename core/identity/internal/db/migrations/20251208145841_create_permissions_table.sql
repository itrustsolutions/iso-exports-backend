-- +goose Up
-- +goose StatementBegin

CREATE TABLE permissions (
    id CHAR(26) PRIMARY KEY NOT NULL,

    -- The service/application the permission belongs to
    service VARCHAR(255) NOT NULL,
    -- The resource the permission applies to
    resource VARCHAR(255) NOT NULL,
    -- The action that can be performed on the resource in the service
    action VARCHAR(255) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (service, resource, action)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd