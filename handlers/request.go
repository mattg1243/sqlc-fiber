package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/willow-server/db"
)

type PaymentInfo struct {
	Venmo  string `json:"venmo"`
	PayPal string `json:"paypal"`
}

// user requests
type createUserRequest struct {
	User struct {
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Fname    string `json:"fname" validate:"required"`
		Lname    string `json:"lname" validate:"required"`
	} `json:"user"`
	ContactInfo struct {
		Phone string `json:"phone"`
		City string `json:"city"`
		State string `json:"state"`
		Street string `json:"street"`
		Zip string `json:"zip"`
	} `json:"contactInfo"`
}

func (r *createUserRequest) bind(c *fiber.Ctx, u *db.User, cI *db.UserContactInfo, v *Validator) error {
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	// bind user;
	u.Fname = r.User.Fname
	u.Lname = r.User.Lname
	u.Email = r.User.Email
	// hash password
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Hash = h
	// save contact info
	cI.Phone = pgtype.Text{String: r.ContactInfo.Phone, Valid: true}
	cI.City = pgtype.Text{String: r.ContactInfo.City, Valid: true}
	cI.State = pgtype.Text{String: r.ContactInfo.State, Valid: true}
	cI.Street = pgtype.Text{String: r.ContactInfo.Street, Valid: true}
	cI.Zip = pgtype.Text{String: r.ContactInfo.Zip, Valid: true}

	return nil
}

type updateUserRequest struct {
	User struct {
		Fname    string `json:"fname" validate:"required"`
		Lname    string `json:"lname" validate:"required"`
		License	 string `json:"license" validate:"required"`
		Nameforheader string `json:"nameForHeader" validate:"required"`
	} `json:"user"`
	ContactInfo struct {
		Phone string `json:"phone" validate:"required"`
		City string `json:"city" validate:"required"`
		State string `json:"state" validate:"required"`
		Street string `json:"street" validate:"required"`
		Zip string `json:"zip" validate:"required"`
		PaymentInfo struct {
			Venmo string `json:"venmo"`
			Paypal string `json:"paypal"`
		} `json:"paymentInfo"`
	} `json:"contactInfo"`
}

func (r *updateUserRequest) bind(c *fiber.Ctx, u *db.User, cI *db.UserContactInfo, v *Validator) error {
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	// bind user
	u.Fname = r.User.Fname
	u.Lname = r.User.Lname
	u.Nameforheader = r.User.Nameforheader
	u.License = pgtype.Text{String: r.User.License, Valid: true}
	// bind contact info
	cI.Phone = pgtype.Text{String: r.ContactInfo.Phone, Valid: true}
	cI.City = pgtype.Text{String: r.ContactInfo.City, Valid: true}
	cI.State = pgtype.Text{String: r.ContactInfo.State, Valid: true}
	cI.Street = pgtype.Text{String: r.ContactInfo.Street, Valid: true}
	cI.Zip = pgtype.Text{String: r.ContactInfo.Zip, Valid: true}
	paymentInfo, err := json.Marshal(r.ContactInfo.PaymentInfo)
	if err != nil {
		return err
	}
	cI.Paymentinfo = paymentInfo


	return nil
}

type loginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *loginUserRequest) bind(c *fiber.Ctx, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	return nil
}

// client requests
type createClientRequest struct {
	Client struct {
		Fname                  string `json:"fname" validate:"required"`
		Lname                  string `json:"lname"`
		Email                  string `json:"email"`
		Rate                   int16  `json:"rate" validate:"required"`
		Phone                  string `json:"phone"`
		Balancenotifythreshold int16  `json:"balanceNotifyThreshold"`
	} `json:"client"`
}

func (r *createClientRequest) bind(c *fiber.Ctx, cl *db.Client, v *Validator) error {
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	cl.Fname = r.Client.Fname
	cl.Lname = pgtype.Text{String: r.Client.Lname, Valid: true}
	cl.Email = pgtype.Text{String: r.Client.Email, Valid: true}
	cl.Rate = int32(r.Client.Rate)
	cl.Phone = pgtype.Text{String: r.Client.Phone, Valid: true}
	cl.Balancenotifythreshold = int32(r.Client.Balancenotifythreshold)

	return nil
}

type updateClientRequest struct {
	Client struct {
		Fname                  string `json:"fname"`
		Lname                  string `json:"lname"`
		Email                  string `json:"email"`
		Phone									 string `json:"phone"`
		Balance                int32  `json:"balance"`
		Balancenotifythreshold int32  `json:"balancenotifythreshold"`
		Rate                   int32  `json:"rate"`
		Isarchived             bool   `json:"isArchived"`
	}
}

func (r *updateClientRequest) bind(c *fiber.Ctx, cl *db.Client, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	cl.Fname = r.Client.Fname
	cl.Lname = pgtype.Text{String: r.Client.Lname, Valid: true}
	cl.Email = pgtype.Text{String: r.Client.Email, Valid: true}
	cl.Phone = pgtype.Text{String: r.Client.Phone, Valid: true}
	cl.Balance = r.Client.Balance
	cl.Balancenotifythreshold = r.Client.Balancenotifythreshold
	cl.Rate = r.Client.Rate
	cl.Isarchived = pgtype.Bool{Bool: r.Client.Isarchived, Valid: true}

	return nil
}

// event requests
type createEventRequest struct {
	Event struct {
		ClientID   	uuid.UUID `json:"client_id" validate:"required"`
		Date       	string    `json:"date" validate:"required"`
		Duration   	float64   `json:"duration" validate:"required"`
		EventTypeID	uuid.UUID `json:"event_type_id" validate:"required"`
		Detail     	string    `json:"detail"`
		Rate       	int32     `json:"rate"`
		Amount     	float64   `json:"amount"`
	} `json:"event"`
}

func Float64ToPgNumeric(f float64) pgtype.Numeric {
	var n pgtype.Numeric
	if err := n.Scan(fmt.Sprintf("%f", f)); err != nil {
		log.Error("error scanning float64 to pg numeric: ", err)
	}
	return n
}

func (r *createEventRequest) bind(c *fiber.Ctx, e *db.Event, v *Validator) error {
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	timeLayout := "2006-01-02 15:04:05"
	timeStr, err := time.Parse(timeLayout, r.Event.Date)
	if err != nil {
		log.Error("error parsing time: ", err)
		return err
	}

	e.Date = pgtype.Timestamp{Time: timeStr, Valid: true}
	e.Duration = Float64ToPgNumeric(r.Event.Duration)
	e.EventTypeID = r.Event.EventTypeID
	e.Detail = pgtype.Text{String: r.Event.Detail}
	e.Rate = r.Event.Rate
	e.Amount = Float64ToPgNumeric(r.Event.Amount)
	e.ClientID = r.Event.ClientID
	log.Info("ClientID: ", e.ClientID)

	return nil
}

// event type requests
type createEventTypeRequest struct {
	EventType struct {
		Name 		string 	`json:"name" validate:"required"`
		Charge 	bool 		`json:"charge" validate:"required"`
	} `json:"eventType" validate:"required"`
}

func (r* createEventTypeRequest) bind(c *fiber.Ctx, et *db.EventType, v *Validator) error {
	// Validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	// Bind
	et.Charge = r.EventType.Charge
	et.Name = r.EventType.Name

	return nil
}

type updateEventTypeRequest struct {
	EventType struct {
		ID 			uuid.UUID 	`json:"id" validate:"required"`
		Name 		string 			`json:"name" validate:"required"`
		Charge 	bool 				`json:"charge" validate:"required"`
	} `json:"eventType" validate:"required"`
}

func (r* updateEventTypeRequest) bind (c *fiber.Ctx, et *db.EventType, v *Validator) error {
		// Validate
		if err := c.BodyParser(r); err != nil {
			return err
		}
	
		if err := v.Validate(r); err != nil {
			return err
		}
		// Bind
		et.ID = r.EventType.ID
		et.Charge = r.EventType.Charge
		et.Name = r.EventType.Name
	
		return nil
}