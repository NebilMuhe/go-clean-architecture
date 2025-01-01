-- name: AddUser :one
INSERT INTO users(email,password)
VALUES($1,$2) RETURNING *;

-- name: GetUserByEmail :one
SELECT id,email,password FROM users WHERE email = $1;

-- name: UserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);