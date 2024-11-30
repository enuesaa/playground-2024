-- name: ListBooks :many
SELECT * FROM books;
-- name: CreateBook :one
INSERT INTO books (id, title, author_id) VALUES (?, ?, ?) RETURNING *;
-- name: DeleteBook :exec
DELETE FROM books WHERE id = ?;

-- name: CreateAuthor :one
INSERT INTO authors (id, name) VALUES (?, ?) RETURNING *;
-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = ?;

