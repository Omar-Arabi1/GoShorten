-- name: AddUrl :exec
INSERT INTO urls (id, long_url, shortened_url_key) VALUES (?, ?, ?);

-- name: GetUrls :many
SELECT * FROM urls;

-- name: UpdateUrl :exec
UPDATE urls SET long_url = ? WHERE id = ?;

-- name: DeleteUrl :exec
DELETE FROM urls WHERE id = ?;

-- name: QueryUrlByShortUrl :one
SELECT * FROM urls WHERE shortened_url_key = ?;
