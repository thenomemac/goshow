package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Listing struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Name        string    `json:"name" db:"name"`
	Owner       string    `json:"owner" db:"owner"`
	Description string    `json:"description" db:"description"`
	Location    string    `json:"location" db:"location"`
	Startdate   time.Time `json:"startdate" db:"startdate"`
	Enddate     time.Time `json:"enddate" db:"enddate"`
}

// String is not required by pop and may be deleted
func (l Listing) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Listings is not required by pop and may be deleted
type Listings []Listing

// String is not required by pop and may be deleted
func (l Listings) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *Listing) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: l.Name, Name: "Name"},
		&validators.StringIsPresent{Field: l.Owner, Name: "Owner"},
		&validators.StringIsPresent{Field: l.Description, Name: "Description"},
		&validators.StringIsPresent{Field: l.Location, Name: "Location"},
		&validators.TimeIsPresent{Field: l.Startdate, Name: "Startdate"},
		&validators.TimeIsPresent{Field: l.Enddate, Name: "Enddate"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *Listing) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *Listing) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
