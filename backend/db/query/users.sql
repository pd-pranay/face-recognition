-- name: CreateUser :one
INSERT INTO users (name, college_name, address, mobile_no, image_path, image_uid) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, college_name, address, mobile_no, image_uid;

-- name: ReadAllUsers :many
SELECT id, name, college_name, address, mobile_no, image_path, image_uid FROM users WHERE is_deleted = false;

-- name: ReadUserByID :one
SELECT id, name, college_name, address, mobile_no, image_path, image_uid FROM users WHERE is_deleted = false AND id = ($1);

-- name: ReadUsersByFace :many
SELECT id, name, college_name, address, mobile_no, image_path, image_uid FROM users WHERE is_deleted = false AND image_uid IN ($1);

-- name: UpdateUserFlush :one
UPDATE users SET is_deleted = true WHERE id = $1 RETURNING *;