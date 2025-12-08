-- +goose Up
-- +goose StatementBegin

CREATE TABLE users_passwords_history (
    id CHAR(26) PRIMARY KEY NOT NULL,
    
    user_id CHAR(26) NOT NULL REFERENCES users(id)
        ON DELETE CASCADE,

    password_hash TEXT NOT NULL,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Creating indexes
CREATE INDEX idx_users_passwords_history_user_id 
    ON users_passwords_history (user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_passwords_history_user_id;
DROP TABLE IF EXISTS users_passwords_history;
-- +goose StatementEnd