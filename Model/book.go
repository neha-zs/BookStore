package Model

import "github.com/google/uuid"

// Book model
type Book struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Author           string    `json:"author"`
	PublishedOn      string    `json:"publishedOn"`
	Publisher        string    `json:"publisher"`
	NumberOfBookSold int       `json:"numberofBookSold"`
}
