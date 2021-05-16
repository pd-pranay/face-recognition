-- name: CreateUser :one
INSERT INTO users (name, college_name, address, mobile_no, image_path, image_uid) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, college_name, address, mobile_no, image_uid;

-- name: ReadAllUsers :many
SELECT id, name, college_name, address, mobile_no, image_path, image_uid FROM users WHERE is_deleted = false;

-- name: ReadUserByID :one
SELECT id, name, college_name, address, mobile_no, image_path, image_uid FROM users WHERE is_deleted = false AND image_uid = ($1);

--name: DeleteUser :one
UPDATE users SET is_deleted = true WHERE id = ($1) RETURNING id, name;

--name: UpdateUserByID :one
UPDATE users SET name = ($1), college_name = ($2), address = ($3), mobile_no = ($4), image_path = ($5), image_uid = ($6) WHERE id = ($1) RETURNING id, name, college_name, address, mobile_no, image_uid, image_path;