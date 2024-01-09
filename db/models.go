// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Client struct {
	ID                     uuid.UUID        `json:"id"`
	UserID                 uuid.UUID        `json:"user_id"`
	Fname                  string           `json:"fname"`
	Lname                  pgtype.Text      `json:"lname"`
	Email                  pgtype.Text      `json:"email"`
	Phone                  pgtype.Text      `json:"phone"`
	Balance                int32            `json:"balance"`
	Balancenotifythreshold int32            `json:"balancenotifythreshold"`
	Rate                   int32            `json:"rate"`
	Isarchived             pgtype.Bool      `json:"isarchived"`
	CreatedAt              pgtype.Timestamp `json:"created_at"`
	UpdatedAt              pgtype.Timestamp `json:"updated_at"`
}

type Event struct {
	ID         uuid.UUID        `json:"id"`
	ClientID   uuid.UUID        `json:"client_id"`
	Date       pgtype.Timestamp `json:"date"`
	Duration   pgtype.Numeric   `json:"duration"`
	Type       pgtype.Text      `json:"type"`
	Detail     pgtype.Text      `json:"detail"`
	Rate       int32            `json:"rate"`
	Amount     pgtype.Numeric   `json:"amount"`
	Newbalance pgtype.Numeric   `json:"newbalance"`
}

type User struct {
	ID            uuid.UUID        `json:"id"`
	Fname         string           `json:"fname"`
	Lname         string           `json:"lname"`
	Email         string           `json:"email"`
	Hash          string           `json:"hash"`
	City          pgtype.Text      `json:"city"`
	Nameforheader string           `json:"nameforheader"`
	Phone         pgtype.Text      `json:"phone"`
	State         pgtype.Text      `json:"state"`
	Street        pgtype.Text      `json:"street"`
	Zip           pgtype.Text      `json:"zip"`
	License       pgtype.Text      `json:"license"`
	Paymentinfo   []byte           `json:"paymentinfo"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}
