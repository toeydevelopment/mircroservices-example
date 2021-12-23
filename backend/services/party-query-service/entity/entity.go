package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Party struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Joined      []string           `json:"joined" bson:"joined,omitempty"`
	ImagePath   string             `json:"image_path" bson:"image_path,omitempty"`
	SeatLimit   int64              `json:"seat_limit" bson:"seat_limit,omitempty"`
	Seat        int64              `json:"seat" bson:"seat,omitempty"`
	Owner       string             `json:"owner" bson:"owner,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt   time.Time          `json:"deleted_at" bson:"deleted_at,omitempty"`
}
