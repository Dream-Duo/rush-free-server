-- name: AddRestaurantAmenity :one
INSERT INTO restaurant_amenities (
    restaurant_id, amenity_name, is_available
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetRestaurantAmenities :many
SELECT * FROM restaurant_amenities WHERE restaurant_id = $1;

-- name: UpdateRestaurantAmenity :exec
UPDATE restaurant_amenities
SET is_available = $3
WHERE restaurant_id = $1 AND amenity_name = $2;

-- name: DeleteRestaurantAmenity :exec
DELETE FROM restaurant_amenities WHERE restaurant_id = $1 AND amenity_name = $2;
