-- name: GetAllPosts :many
SELECT id, post_name, post_data, tags FROM posts;

-- name: CreatePost :one
INSERT INTO posts (post_name, post_data, tags) VALUES ($1, $2, $3) RETURNING id, post_name, post_data, tags;

-- name: GetPostByID :one
SELECT id, post_name, post_data, tags FROM posts WHERE id = ($1);

-- name: GetPostsByTag :many
SELECT id, post_name, post_data, tags FROM posts WHERE tags @> ($1);