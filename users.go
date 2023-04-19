package models

type User struct {
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}
type game struct {
	name        string
	platform    string
	releaseYear int
	developer   string
	publisher   string
}
type review struct {
	gameName    string
	reviewer    string
	reviewText string
	rating      int
}