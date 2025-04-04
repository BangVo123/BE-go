package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DigitCode struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Email     string             `bson:"email" json:"email"`
	Code      int                `bson:"code" json:"code"`
	Type      string             `bson:"type" json:"type"`
	ExpiredAt time.Time          `bson:"expired_at" json:"expired_at"`
}

// enum: ["auth", "reset"],
