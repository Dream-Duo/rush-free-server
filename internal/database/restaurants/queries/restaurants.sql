-- name: CreateRestaurant :one
INSERT INTO restaurants (
    vendor_id, name, description, cuisine_type, address, location, phone, email, seating_capacity, status
) VALUES (
    $1, $2, $3, $4, $5, ST_GeographyFromText('SRID=4326;POINT(' || $6 || ' ' || $7 || ')'), $8, $9, $10, $11
) RETURNING *;

-- name: GetRestaurantByID :one
SELECT * FROM restaurants WHERE restaurant_id = $1;

-- name: GetAllRestaurants :many
SELECT * FROM restaurants ORDER BY created_at DESC;

-- name: GetRestaurantsNearLocation :many
SELECT * FROM restaurants
WHERE ST_DWithin(location, ST_GeographyFromText('SRID=4326;POINT(' || $1 || ' ' || $2 || ')'), $3);

-- name: UpdateRestaurantStatus :exec
UPDATE restaurants
SET status = $2, updated_at = CURRENT_TIMESTAMP
WHERE restaurant_id = $1;

-- name: DeleteRestaurant :exec
DELETE FROM restaurants WHERE restaurant_id = $1;
