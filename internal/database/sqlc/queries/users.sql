-- name: GetUserById :one
SELECT * from "user" WHERE id=$1;