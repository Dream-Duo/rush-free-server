-- name: GetRestaurantReviews :many
SELECT r.review_id, 
       r.rating, 
       r.review_text, 
       u.first_name || ' ' || u.last_name AS user_name, 
       r.created_at
FROM restaurant_reviews r
JOIN users u ON r.user_id = u.user_id
WHERE r.restaurant_id = $1
ORDER BY r.created_at DESC;

