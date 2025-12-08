-- +goose Up
-- +goose StatementBegin

CREATE TABLE system_config (
    id CHAR(26) PRIMARY KEY NOT NULL,
    
    -- Session settings
    session_idle_timeout_minutes INTEGER NOT NULL DEFAULT 90,
    session_absolute_timeout_hours INTEGER NOT NULL DEFAULT 720,
    session_ip_binding_enabled BOOLEAN NOT NULL DEFAULT true,
    
    -- Password settings
    password_expiration_days INTEGER NOT NULL DEFAULT 90,
    password_history_count INTEGER NOT NULL DEFAULT 10,
    password_reset_token_expiration_minutes INTEGER NOT NULL DEFAULT 10,
    
    -- Account lockout settings
    max_failed_login_attempts INTEGER NOT NULL DEFAULT 5,
    account_lockout_duration_minutes INTEGER NOT NULL DEFAULT 15,
    failed_attempt_reset_window_minutes INTEGER NOT NULL DEFAULT 60,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Ensure only one config row exists
CREATE UNIQUE INDEX idx_system_config_singleton ON system_config ((TRUE));

-- Insert default configuration (single row)
-- Using a safe, valid ULID for singleton row
INSERT INTO system_config (id) VALUES ('01F0000000000000000000ABC1');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS system_config;
-- +goose StatementEnd