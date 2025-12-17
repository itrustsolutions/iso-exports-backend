-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    email,
    password_hash,
    has_system_access,
    has_all_namespaces_access,
    is_active,
    failed_attempts_count,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING
    id,
    username,
    email,
    has_system_access,
    has_all_namespaces_access,
    is_active,
    failed_attempts_count,
    last_failed_attempt_at,
    locked_until_at,
    deleted_at,
    created_at,
    updated_at;
