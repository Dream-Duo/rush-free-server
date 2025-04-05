-- name: CreateRestaurantHour :one
INSERT INTO restaurant_hours (
    restaurant_id, day_of_week, opening_time, closing_time, is_closed
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetRestaurantHours :many
SELECT * FROM restaurant_hours WHERE restaurant_id = $1 ORDER BY day_of_week;

-- name: UpdateRestaurantHour :exec
UPDATE restaurant_hours
SET opening_time = $3, closing_time = $4, is_closed = $5
WHERE restaurant_id = $1 AND day_of_week = $2;

-- name: DeleteRestaurantHour :exec
DELETE FROM restaurant_hours WHERE restaurant_id = $1 AND day_of_week = $2;
