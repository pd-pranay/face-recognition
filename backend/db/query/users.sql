-- name: CreateUser :one
INSERT INTO users (name, college_name, address, mobile_no, image_path, image_uid) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, college_name, address, mobile_no, image_uid;

-- name: ReadAllUsers :many
SELECT id, name, college_name, address, mobile_no, image_path, image_uid FROM users WHERE is_deleted = false;

-- name: ReadUserByID :one
SELECT id, name, college_name, address, mobile_no, image_path, image_uid FROM users WHERE is_deleted = false AND image_uid = ($1);

-- name: ReadUsersByFace :many
SELECT id, name, college_name, address, mobile_no, image_path, image_uid FROM users WHERE is_deleted = false AND image_uid IN ($1);

-- name: UpdateUser :one
UPDATE users SET name = ($2), college_name = ($3), address = ($4), mobile_no = ($5), image_path = ($6), image_uid = ($7) WHERE image_uid = ($1) RETURNING *;

-- name: DeleteUsersById :exec
DELETE FROM users WHERE image_uid = $1;