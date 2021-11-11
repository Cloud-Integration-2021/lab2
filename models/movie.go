package models

type Movie struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Plot        string `json:"plot"`
	ReleaseDate string `json:"releaseDate"`
}
