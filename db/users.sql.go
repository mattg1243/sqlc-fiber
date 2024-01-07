// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    fname, lname, email, salt, "hash", city, nameforheader, phone, "state", street, zip, license, paymentinfo
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING id, fname, lname, email, salt, hash, city, nameforheader, phone, state, street, zip, license, paymentinfo, created_at, updated_at
`

type CreateUserParams struct {
	Fname         string      `json:"fname"`
	Lname         string      `json:"lname"`
	Email         string      `json:"email"`
	Salt          string      `json:"salt"`
	Hash          string      `json:"hash"`
	City          pgtype.Text `json:"city"`
	Nameforheader string      `json:"nameforheader"`
	Phone         pgtype.Text `json:"phone"`
	State         pgtype.Text `json:"state"`
	Street        pgtype.Text `json:"street"`
	Zip           pgtype.Text `json:"zip"`
	License       pgtype.Text `json:"license"`
	Paymentinfo   []byte      `json:"paymentinfo"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Fname,
		arg.Lname,
		arg.Email,
		arg.Salt,
		arg.Hash,
		arg.City,
		arg.Nameforheader,
		arg.Phone,
		arg.State,
		arg.Street,
		arg.Zip,
		arg.License,
		arg.Paymentinfo,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fname,
		&i.Lname,
		&i.Email,
		&i.Salt,
		&i.Hash,
		&i.City,
		&i.Nameforheader,
		&i.Phone,
		&i.State,
		&i.Street,
		&i.Zip,
		&i.License,
		&i.Paymentinfo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, fname, lname, email, salt, hash, city, nameforheader, phone, state, street, zip, license, paymentinfo, created_at, updated_at FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fname,
		&i.Lname,
		&i.Email,
		&i.Salt,
		&i.Hash,
		&i.City,
		&i.Nameforheader,
		&i.Phone,
		&i.State,
		&i.Street,
		&i.Zip,
		&i.License,
		&i.Paymentinfo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    fname = $1,
    lname = $2,
    email = $3,
    city = $4,
    nameForHeader = $5,
    phone = $6,
    "state" = $7,
    street = $8,
    zip = $9,
    license = $10,
    paymentInfo = $11,
    updated_at = NOW()
WHERE
    id = $1
RETURNING id, fname, lname, email, salt, hash, city, nameforheader, phone, state, street, zip, license, paymentinfo, created_at, updated_at
`

type UpdateUserParams struct {
	Fname         string      `json:"fname"`
	Lname         string      `json:"lname"`
	Email         string      `json:"email"`
	City          pgtype.Text `json:"city"`
	Nameforheader string      `json:"nameforheader"`
	Phone         pgtype.Text `json:"phone"`
	State         pgtype.Text `json:"state"`
	Street        pgtype.Text `json:"street"`
	Zip           pgtype.Text `json:"zip"`
	License       pgtype.Text `json:"license"`
	Paymentinfo   []byte      `json:"paymentinfo"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Fname,
		arg.Lname,
		arg.Email,
		arg.City,
		arg.Nameforheader,
		arg.Phone,
		arg.State,
		arg.Street,
		arg.Zip,
		arg.License,
		arg.Paymentinfo,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fname,
		&i.Lname,
		&i.Email,
		&i.Salt,
		&i.Hash,
		&i.City,
		&i.Nameforheader,
		&i.Phone,
		&i.State,
		&i.Street,
		&i.Zip,
		&i.License,
		&i.Paymentinfo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
