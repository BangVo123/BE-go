package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sessions struct {
	Id     primitive.ObjectID `bson:"_id" json:"_id"`
	UserId primitive.ObjectID `bson:"user_id" json:"user_id"`
}
