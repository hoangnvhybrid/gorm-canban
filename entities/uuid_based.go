package entities

import (
	"database/sql/driver"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UUIDBasedAbstract struct {
	ID uint64 `json:"-"`

	IdentifierAPI UUID `gorm:"type:bytea;unique" json:"uuid"`

	CreatedAt time.Time `json:"createdAt"`
}

type UUIDBasedFromClientAbstract struct {
	UUIDBasedAbstract

	UserID           uint64 `json:"-"` // NOTE owner user
	ClientID         UUID   `gorm:"type:bytea;not null" json:"-"`
	IdentifierClient UUID   `gorm:"type:bytea;not null" json:"-"`
}

type UUID struct {
	uuid.UUID
}

func (u *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case []uint8:
		if len(src) == 16 {
			return u.UnmarshalBinary(src)
		}
		return u.UnmarshalText(src)
	default:
		return u.UUID.Scan(src)
	}
}

func (u UUID) Value() (driver.Value, error) {
	return u.Bytes(), nil
}

func NewUUIDV4() UUID {
	uudi, _ := uuid.NewV4()
	return UUID{uudi}
}

func UUIDFromString(uuidString string) (UUID, error) {
	u, err := uuid.FromString(uuidString)
	if err != nil {
		return UUID{u}, err
	}
	return UUID{u}, nil
}
