package models

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name string `gorm:"type:text" json:"name"`
}

type CreateGenreInput struct {
	Name string `json:"name"`
}

type UpdateGenreInput struct {
	Name string `json:"name"`
}
