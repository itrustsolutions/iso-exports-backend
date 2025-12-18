-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    email,
    password_hash,
    has_system_access,
    has_all_namespaces_access,
    is_active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING
    id,
    username,
    email,
    has_system_access,
    has_all_namespaces_access,
    is_active,
    created_at,
    updated_at;
