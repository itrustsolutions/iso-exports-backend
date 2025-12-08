-- +goose Up
-- +goose StatementBegin

CREATE TABLE password_reset_tokens (
    id CHAR(26) PRIMARY KEY NOT NULL,
    
    -- User requesting the password reset
    user_id CHAR(26) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- The secure, random token sent to user's email
    token_hash VARCHAR(255) UNIQUE NOT NULL,
    
    -- IP address from where the reset was requested
    ip_address INET NOT NULL,
    
    -- User agent from the reset request
    user_agent TEXT,
    
    -- When the token was created
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Creating indexes
-- Index for finding valid tokens
CREATE INDEX idx_password_reset_tokens_token_hash ON password_reset_tokens (token_hash);
-- Index for finding user's active tokens (to invalidate old ones on new request)
CREATE INDEX idx_password_reset_tokens_user_id ON password_reset_tokens (user_id);
-- Index for cleanup/purging expired tokens (using created_at)
CREATE INDEX idx_password_reset_tokens_created_at ON password_reset_tokens (created_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_password_reset_tokens_created_at;
DROP INDEX IF EXISTS idx_password_reset_tokens_user_id;
DROP INDEX IF EXISTS idx_password_reset_tokens_token_hash;
DROP TABLE IF EXISTS password_reset_tokens;
-- +goose StatementEnd