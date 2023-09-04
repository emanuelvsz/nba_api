package dto

import (
	"time"

	"github.com/google/uuid"
)

type PlayerDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Number    int       `json:"number"`
	TeamID    uuid.UUID `json:"team_id"`
	BirthDate time.Time `json:"birth_date"`
	Height    float64   `json:"height"`
}

func NewPlayerDTO(id, teamID uuid.UUID, name string, number int, height float64) *PlayerDTO {
	return &PlayerDTO{
		ID:     id,
		Name:   name,
		Number: number,
		TeamID: teamID,
		Height: height,
	}
}
