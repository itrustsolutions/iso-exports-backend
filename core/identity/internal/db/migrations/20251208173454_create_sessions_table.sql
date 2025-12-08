-- +goose Up
-- +goose StatementBegin

CREATE TABLE sessions (
    id CHAR(26) PRIMARY KEY NOT NULL,
    
    -- User who owns this session
    user_id CHAR(26) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- The session token (secure, random)
    token VARCHAR(255) UNIQUE NOT NULL,
    
    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_activity_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    -- Timeout tracking (calculated values, not stored directly - use system_config)
    -- idle_expires_at = last_activity_at + system_config.session_idle_timeout_minutes
    -- absolute_expires_at = created_at + system_config.session_absolute_timeout_hours
    
    -- Network information
    ip_address_at_creation INET NOT NULL,
    ip_address_current INET NOT NULL,
    
    -- User agent and parsed device information
    user_agent TEXT NOT NULL,
    device_type VARCHAR(50),
    operating_system VARCHAR(100),
    browser VARCHAR(100),
    
    -- Geolocation (coarse, from IP)
    geo_country VARCHAR(2),
    geo_city VARCHAR(100),
    
    -- Session status
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    
    -- When the session ended (NULL if still active)
    ended_at TIMESTAMPTZ,
    
    -- Who terminated the session (NULL if user logout or system expiration)
    terminated_by_user_id CHAR(26) REFERENCES users(id) ON DELETE SET NULL
);

-- Creating indexes
CREATE INDEX idx_sessions_token ON sessions (token) WHERE status = 'active';
CREATE INDEX idx_sessions_user_active ON sessions (user_id, status) WHERE status = 'active';
CREATE INDEX idx_sessions_user_id ON sessions (user_id, created_at DESC);
CREATE INDEX idx_sessions_last_activity ON sessions (last_activity_at) WHERE status = 'active';
CREATE INDEX idx_sessions_ip_address ON sessions (ip_address_current);
CREATE INDEX idx_sessions_status ON sessions (status, created_at DESC);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_sessions_status;
DROP INDEX IF EXISTS idx_sessions_ip_address;
DROP INDEX IF EXISTS idx_sessions_last_activity;
DROP INDEX IF EXISTS idx_sessions_user_id;
DROP INDEX IF EXISTS idx_sessions_user_active;
DROP INDEX IF EXISTS idx_sessions_token;
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd