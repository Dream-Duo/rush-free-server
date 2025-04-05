-- name: AddRestaurantImage :one
INSERT INTO restaurant_images (
    restaurant_id, image_url, image_type, is_primary
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetRestaurantImages :many
SELECT * FROM restaurant_images WHERE restaurant_id = $1 ORDER BY uploaded_at DESC;

-- name: DeleteRestaurantImage :exec
DELETE FROM restaurant_images WHERE image_id = $1;
