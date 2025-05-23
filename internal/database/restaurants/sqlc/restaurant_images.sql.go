// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: restaurant_images.sql

package restaurants

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addRestaurantImage = `-- name: AddRestaurantImage :one
INSERT INTO restaurant_images (
    restaurant_id, image_url, image_type, is_primary
) VALUES (
    $1, $2, $3, $4
) RETURNING image_id, restaurant_id, image_url, image_type, is_primary, uploaded_at
`

type AddRestaurantImageParams struct {
	RestaurantID int32
	ImageUrl     string
	ImageType    interface{}
	IsPrimary    pgtype.Bool
}

func (q *Queries) AddRestaurantImage(ctx context.Context, arg AddRestaurantImageParams) (RestaurantImage, error) {
	row := q.db.QueryRow(ctx, addRestaurantImage,
		arg.RestaurantID,
		arg.ImageUrl,
		arg.ImageType,
		arg.IsPrimary,
	)
	var i RestaurantImage
	err := row.Scan(
		&i.ImageID,
		&i.RestaurantID,
		&i.ImageUrl,
		&i.ImageType,
		&i.IsPrimary,
		&i.UploadedAt,
	)
	return i, err
}

const deleteRestaurantImage = `-- name: DeleteRestaurantImage :exec
DELETE FROM restaurant_images WHERE image_id = $1
`

func (q *Queries) DeleteRestaurantImage(ctx context.Context, imageID int32) error {
	_, err := q.db.Exec(ctx, deleteRestaurantImage, imageID)
	return err
}

const getRestaurantImages = `-- name: GetRestaurantImages :many
SELECT image_id, restaurant_id, image_url, image_type, is_primary, uploaded_at FROM restaurant_images WHERE restaurant_id = $1 ORDER BY uploaded_at DESC
`

func (q *Queries) GetRestaurantImages(ctx context.Context, restaurantID int32) ([]RestaurantImage, error) {
	rows, err := q.db.Query(ctx, getRestaurantImages, restaurantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RestaurantImage
	for rows.Next() {
		var i RestaurantImage
		if err := rows.Scan(
			&i.ImageID,
			&i.RestaurantID,
			&i.ImageUrl,
			&i.ImageType,
			&i.IsPrimary,
			&i.UploadedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
