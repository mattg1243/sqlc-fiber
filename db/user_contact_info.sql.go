// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user_contact_info.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUserContactInfo = `-- name: CreateUserContactInfo :one
INSERT INTO user_contact_info (
  id, user_id, phone, city, "state", street, zip, paymentInfo, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW()
)
RETURNING id, user_id, phone, city, state, street, zip, paymentinfo, created_at, updated_at
`

type CreateUserContactInfoParams struct {
	ID          uuid.UUID   `json:"id"`
	UserID      uuid.UUID   `json:"user_id"`
	Phone       pgtype.Text `json:"phone"`
	City        pgtype.Text `json:"city"`
	State       pgtype.Text `json:"state"`
	Street      pgtype.Text `json:"street"`
	Zip         pgtype.Text `json:"zip"`
	Paymentinfo []byte      `json:"paymentinfo"`
}

func (q *Queries) CreateUserContactInfo(ctx context.Context, arg CreateUserContactInfoParams) (UserContactInfo, error) {
	row := q.db.QueryRow(ctx, createUserContactInfo,
		arg.ID,
		arg.UserID,
		arg.Phone,
		arg.City,
		arg.State,
		arg.Street,
		arg.Zip,
		arg.Paymentinfo,
	)
	var i UserContactInfo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Phone,
		&i.City,
		&i.State,
		&i.Street,
		&i.Zip,
		&i.Paymentinfo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserContactInfo = `-- name: GetUserContactInfo :one
SELECT id, user_id, phone, city, state, street, zip, paymentinfo, created_at, updated_at FROM user_contact_info WHERE user_id = $1
`

func (q *Queries) GetUserContactInfo(ctx context.Context, userID uuid.UUID) (UserContactInfo, error) {
	row := q.db.QueryRow(ctx, getUserContactInfo, userID)
	var i UserContactInfo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Phone,
		&i.City,
		&i.State,
		&i.Street,
		&i.Zip,
		&i.Paymentinfo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserContactInfo = `-- name: UpdateUserContactInfo :one
UPDATE user_contact_info 
SET
  phone = $1,
  city = $2,
  "state" = $3,
  street = $4,
  zip = $5,
  paymentInfo = $6
WHERE
  user_id = $7
RETURNING id, user_id, phone, city, state, street, zip, paymentinfo, created_at, updated_at
`

type UpdateUserContactInfoParams struct {
	Phone       pgtype.Text `json:"phone"`
	City        pgtype.Text `json:"city"`
	State       pgtype.Text `json:"state"`
	Street      pgtype.Text `json:"street"`
	Zip         pgtype.Text `json:"zip"`
	Paymentinfo []byte      `json:"paymentinfo"`
	UserID      uuid.UUID   `json:"user_id"`
}

func (q *Queries) UpdateUserContactInfo(ctx context.Context, arg UpdateUserContactInfoParams) (UserContactInfo, error) {
	row := q.db.QueryRow(ctx, updateUserContactInfo,
		arg.Phone,
		arg.City,
		arg.State,
		arg.Street,
		arg.Zip,
		arg.Paymentinfo,
		arg.UserID,
	)
	var i UserContactInfo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Phone,
		&i.City,
		&i.State,
		&i.Street,
		&i.Zip,
		&i.Paymentinfo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
