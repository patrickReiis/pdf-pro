package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type JSONB []string

type UserAccount struct {
	gorm.Model
	Email             string `gorm:"unique:true"`
	Password          string
	RequestsTimestamp JSONB  `gorm:"type:jsonb;default:'[]';not null"`
	ApiKey            string `gorm:"unique:true;size:300"`
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if ok == false {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, a)
}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}
