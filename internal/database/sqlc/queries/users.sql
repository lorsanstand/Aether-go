-- name: GetUserById :one
SELECT * from "user" WHERE id=$1;

-- name: CreateUser :exec
INSERT INTO "user" (display_name, username, email, is_active, is_verified, is_superuser, hashed_password)
VALUES  ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "user" SET display_name= $2, username = $3, birth_day = $4, description = $5
WHERE id = $1;