package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Address struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Address   string    `json:"address" db:"address"`
	City      string    `json:"city" db:"city"`
	State     string    `json:"state" db:"state"`
	Zipcode   string    `json:"zipcode" db:"zipcode"`
	ListingID uuid.UUID `json:"listing_id" db:"listing_id"`
}

// String is not required by pop and may be deleted
func (a Address) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Addresses is not required by pop and may be deleted
type Addresses []Address

// String is not required by pop and may be deleted
func (a Addresses) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Address) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Address, Name: "Address"},
		&validators.StringIsPresent{Field: a.City, Name: "City"},
		&validators.StringIsPresent{Field: a.State, Name: "State"},
		&validators.StringIsPresent{Field: a.Zipcode, Name: "Zipcode"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Address) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Address) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
