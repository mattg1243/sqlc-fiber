// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: events.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (
    client_id, user_id, date, duration, event_type_id, detail, rate, amount, running_balance, id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING id, user_id, client_id, date, duration, event_type_id, detail, rate, amount, running_balance
`

type CreateEventParams struct {
	ClientID       uuid.UUID        `json:"client_id"`
	UserID         uuid.UUID        `json:"user_id"`
	Date           pgtype.Timestamp `json:"date"`
	Duration       pgtype.Numeric   `json:"duration"`
	EventTypeID    uuid.UUID        `json:"event_type_id"`
	Detail         pgtype.Text      `json:"detail"`
	Rate           int32            `json:"rate"`
	Amount         int32            `json:"amount"`
	RunningBalance int32            `json:"running_balance"`
	ID             uuid.UUID        `json:"id"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, createEvent,
		arg.ClientID,
		arg.UserID,
		arg.Date,
		arg.Duration,
		arg.EventTypeID,
		arg.Detail,
		arg.Rate,
		arg.Amount,
		arg.RunningBalance,
		arg.ID,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ClientID,
		&i.Date,
		&i.Duration,
		&i.EventTypeID,
		&i.Detail,
		&i.Rate,
		&i.Amount,
		&i.RunningBalance,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteEvent, id)
	return err
}

const getEvent = `-- name: GetEvent :one
SELECT 
    e.id as id,
    e.client_id as client_id,
    e.user_id as user_id,
    e.date::timestamptz as "date",
    e.duration as duration,
    et.id as event_type_id,
    e.detail as detail,
    e.rate as rate,
    e.amount::INTEGER as amount,
    e.running_balance::INTEGER as running_balance,
    et.charge as charge
FROM events e
INNER JOIN event_types et ON e.event_type_id = et.id
WHERE e.id = $1
`

type GetEventRow struct {
	ID             uuid.UUID          `json:"id"`
	ClientID       uuid.UUID          `json:"client_id"`
	UserID         uuid.UUID          `json:"user_id"`
	Date           pgtype.Timestamptz `json:"date"`
	Duration       pgtype.Numeric     `json:"duration"`
	EventTypeID    uuid.UUID          `json:"event_type_id"`
	Detail         pgtype.Text        `json:"detail"`
	Rate           int32              `json:"rate"`
	Amount         int32              `json:"amount"`
	RunningBalance int32              `json:"running_balance"`
	Charge         bool               `json:"charge"`
}

func (q *Queries) GetEvent(ctx context.Context, id uuid.UUID) (GetEventRow, error) {
	row := q.db.QueryRow(ctx, getEvent, id)
	var i GetEventRow
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.UserID,
		&i.Date,
		&i.Duration,
		&i.EventTypeID,
		&i.Detail,
		&i.Rate,
		&i.Amount,
		&i.RunningBalance,
		&i.Charge,
	)
	return i, err
}

const getEvents = `-- name: GetEvents :many
SELECT 
    e.id as id,
    e.user_id as user_id,
    e.client_id as client_id,
    e.date::timestamptz as "date",
    e.duration as duration,
    et.id as event_type_id,
    e.detail as detail,
    e.rate as rate,
    e.amount::INTEGER as amount,
    e.running_balance::INTEGER as running_balance,
    et.charge as charge
FROM events e
INNER JOIN event_types et ON e.event_type_id = et.id
WHERE e.client_id = $1 or e.user_id = $1
ORDER BY e.date ASC
`

type GetEventsRow struct {
	ID             uuid.UUID          `json:"id"`
	UserID         uuid.UUID          `json:"user_id"`
	ClientID       uuid.UUID          `json:"client_id"`
	Date           pgtype.Timestamptz `json:"date"`
	Duration       pgtype.Numeric     `json:"duration"`
	EventTypeID    uuid.UUID          `json:"event_type_id"`
	Detail         pgtype.Text        `json:"detail"`
	Rate           int32              `json:"rate"`
	Amount         int32              `json:"amount"`
	RunningBalance int32              `json:"running_balance"`
	Charge         bool               `json:"charge"`
}

func (q *Queries) GetEvents(ctx context.Context, clientID uuid.UUID) ([]GetEventsRow, error) {
	rows, err := q.db.Query(ctx, getEvents, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEventsRow
	for rows.Next() {
		var i GetEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ClientID,
			&i.Date,
			&i.Duration,
			&i.EventTypeID,
			&i.Detail,
			&i.Rate,
			&i.Amount,
			&i.RunningBalance,
			&i.Charge,
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

const updateEvent = `-- name: UpdateEvent :one
UPDATE events
SET
    date = $2,
    duration = $3,
    event_type_id = $4,
    detail = $5,
    rate = $6,
    amount = $7,
    running_balance = $8
WHERE
    id = $1
RETURNING id, user_id, client_id, date, duration, event_type_id, detail, rate, amount, running_balance
`

type UpdateEventParams struct {
	ID             uuid.UUID        `json:"id"`
	Date           pgtype.Timestamp `json:"date"`
	Duration       pgtype.Numeric   `json:"duration"`
	EventTypeID    uuid.UUID        `json:"event_type_id"`
	Detail         pgtype.Text      `json:"detail"`
	Rate           int32            `json:"rate"`
	Amount         int32            `json:"amount"`
	RunningBalance int32            `json:"running_balance"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, updateEvent,
		arg.ID,
		arg.Date,
		arg.Duration,
		arg.EventTypeID,
		arg.Detail,
		arg.Rate,
		arg.Amount,
		arg.RunningBalance,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ClientID,
		&i.Date,
		&i.Duration,
		&i.EventTypeID,
		&i.Detail,
		&i.Rate,
		&i.Amount,
		&i.RunningBalance,
	)
	return i, err
}

const updateRunningBalance = `-- name: UpdateRunningBalance :exec
UPDATE events
SET
    running_balance = $2
WHERE
    id = $1
`

type UpdateRunningBalanceParams struct {
	ID             uuid.UUID `json:"id"`
	RunningBalance int32     `json:"running_balance"`
}

func (q *Queries) UpdateRunningBalance(ctx context.Context, arg UpdateRunningBalanceParams) error {
	_, err := q.db.Exec(ctx, updateRunningBalance, arg.ID, arg.RunningBalance)
	return err
}
