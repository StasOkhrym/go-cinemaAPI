package models

type Movie struct {
	Title       string  `gorm:"type:text" json:"title"`
	Description string  `gorm:"type:text" json:"description"`
	Duration    int     `gorm:"type:integer" json:"duration"` // in minutes
	Genres      []Genre `gorm:"many2many:genres" json:"genres"`
	Actors      []Actor `gorm:"many2many:actors" json:"actors"`
	Image       string  `gorm:"type:text" json:"image"`
}

type CreateMovieInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Genres      []int  `json:"genres"`
	Actors      []int  `json:"actors"`
	Image       string `json:"image"`
}

type UpdateMovieInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Genres      []int  `json:"genres"`
	Actors      []int  `json:"actors"`
	Image       string `json:"image"`
}
