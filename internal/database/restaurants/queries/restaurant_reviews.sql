-- name: AddRestaurantReview :one
INSERT INTO restaurant_reviews (
    user_id, restaurant_id, rating, review_text
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateRestaurantReview :exec
UPDATE restaurant_reviews
SET rating = $2, review_text = $3
WHERE review_id = $1;

-- name: DeleteRestaurantReview :exec
DELETE FROM restaurant_reviews WHERE review_id = $1;
