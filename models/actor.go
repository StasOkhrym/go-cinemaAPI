package models

import (
	"gorm.io/gorm"
)

type Actor struct {
	gorm.Model
	FirstName string `gorm:"type:text" json:"first_name"`
	LastName  string `gorm:"type:text" json:"last_name"`
}

func (a *Actor) FullName() string {
	return a.FirstName + " " + a.LastName
}
