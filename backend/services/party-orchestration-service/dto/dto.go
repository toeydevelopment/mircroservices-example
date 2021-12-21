package dto

import "time"

type CreatePartyDTO struct {
	Name        string
	Description string
	SeatLimit   int64
	UserEmail   string
}

type UpdatePartyDTO struct {
	Name        *string
	Description *string
	SeatLimit   *int64
	UserEmail   string
}

type PartyDTO struct {
	ID          string
	Name        string
	Description *string
	SeatLimit   *int64
	Seat        *int64
	ImagePath   *string
	Joined      []string
	Owner       string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeleteAt    *time.Time
}
