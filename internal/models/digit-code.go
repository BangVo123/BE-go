package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DigitCode struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Email      string             `bson:"email" json:"email"`
	Code       int                `bson:"code" json:"code"`
	Type       string             `bson:"type" json:"type"`
	ExpiriedAt time.Time          `bson:"expired_at" json:"expired_at"`
}
