-- name: CreateAdmin :one
INSERT INTO admin (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email;

-- name: Login :one
SELECT id, name, email, password FROM admin WHERE email = ($1);