// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: restaurant_reviews.sql

package restaurants

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addRestaurantReview = `-- name: AddRestaurantReview :one
INSERT INTO restaurant_reviews (
    user_id, restaurant_id, rating, review_text
) VALUES (
    $1, $2, $3, $4
) RETURNING review_id, user_id, restaurant_id, rating, review_text, created_at
`

type AddRestaurantReviewParams struct {
	UserID       int32
	RestaurantID int32
	Rating       pgtype.Int4
	ReviewText   pgtype.Text
}

func (q *Queries) AddRestaurantReview(ctx context.Context, arg AddRestaurantReviewParams) (RestaurantReview, error) {
	row := q.db.QueryRow(ctx, addRestaurantReview,
		arg.UserID,
		arg.RestaurantID,
		arg.Rating,
		arg.ReviewText,
	)
	var i RestaurantReview
	err := row.Scan(
		&i.ReviewID,
		&i.UserID,
		&i.RestaurantID,
		&i.Rating,
		&i.ReviewText,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRestaurantReview = `-- name: DeleteRestaurantReview :exec
DELETE FROM restaurant_reviews WHERE review_id = $1
`

func (q *Queries) DeleteRestaurantReview(ctx context.Context, reviewID int32) error {
	_, err := q.db.Exec(ctx, deleteRestaurantReview, reviewID)
	return err
}

const updateRestaurantReview = `-- name: UpdateRestaurantReview :exec
UPDATE restaurant_reviews
SET rating = $2, review_text = $3
WHERE review_id = $1
`

type UpdateRestaurantReviewParams struct {
	ReviewID   int32
	Rating     pgtype.Int4
	ReviewText pgtype.Text
}

func (q *Queries) UpdateRestaurantReview(ctx context.Context, arg UpdateRestaurantReviewParams) error {
	_, err := q.db.Exec(ctx, updateRestaurantReview, arg.ReviewID, arg.Rating, arg.ReviewText)
	return err
}
