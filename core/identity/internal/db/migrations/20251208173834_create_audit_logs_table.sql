-- +goose Up
-- +goose StatementBegin

CREATE TABLE audit_logs (
    id CHAR(26) PRIMARY KEY NOT NULL,
    
    -- Actor who performed the action (always present, never null)
    user_id CHAR(26) NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    
    -- Which service generated this log
    service VARCHAR(255) NOT NULL,
    
    -- What kind of resource was affected
    resource VARCHAR(255) NOT NULL,
    
    -- What action was performed
    action VARCHAR(255) NOT NULL,
    
    -- Specific entity identifier(s) - array for bulk operations
    resource_ids TEXT[] NOT NULL,
    
    -- Which namespace the action occurred in (NULL for non-namespaced operations)
    namespace_id CHAR(26) REFERENCES namespaces(id) ON DELETE SET NULL,

    -- Which session the action was performed in
    session_id CHAR(26) NOT NULL REFERENCES sessions(id) ON DELETE RESTRICT,
    
    -- Network and device information
    ip_address INET NOT NULL,
    user_agent TEXT,
    device_type VARCHAR(50),
    operating_system VARCHAR(100),
    browser VARCHAR(100),
    
    -- Geolocation information
    geo_country VARCHAR(2),
    geo_city VARCHAR(100),
    
    -- For update actions: only the fields that changed
    changes JSONB,
    
    -- Flexible metadata for service-specific details and programmer-defined fields
    metadata JSONB,
    
    -- When the action occurred
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Creating indexes
CREATE INDEX idx_audit_logs_user_id ON audit_logs (user_id, created_at DESC);
CREATE INDEX idx_audit_logs_resource ON audit_logs (service, resource, created_at DESC);
CREATE INDEX idx_audit_logs_resource_ids ON audit_logs USING GIN (resource_ids);
CREATE INDEX idx_audit_logs_service ON audit_logs (service, created_at DESC);
CREATE INDEX idx_audit_logs_action ON audit_logs (action, created_at DESC);
CREATE INDEX idx_audit_logs_namespace_id ON audit_logs (namespace_id, created_at DESC) WHERE namespace_id IS NOT NULL;
CREATE INDEX idx_audit_logs_created_at ON audit_logs (created_at DESC);
CREATE INDEX idx_audit_logs_ip_address ON audit_logs (ip_address);
CREATE INDEX idx_audit_logs_metadata ON audit_logs USING GIN (metadata);
CREATE INDEX idx_audit_logs_changes ON audit_logs USING GIN (changes);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_audit_logs_changes;
DROP INDEX IF EXISTS idx_audit_logs_metadata;
DROP INDEX IF EXISTS idx_audit_logs_ip_address;
DROP INDEX IF EXISTS idx_audit_logs_created_at;
DROP INDEX IF EXISTS idx_audit_logs_namespace_id;
DROP INDEX IF EXISTS idx_audit_logs_action;
DROP INDEX IF EXISTS idx_audit_logs_service;
DROP INDEX IF EXISTS idx_audit_logs_resource_ids;
DROP INDEX IF EXISTS idx_audit_logs_resource;
DROP INDEX IF EXISTS idx_audit_logs_user_id;
DROP TABLE IF EXISTS audit_logs;
-- +goose StatementEnd