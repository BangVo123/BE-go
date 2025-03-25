package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Like struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id"`
	UserId  primitive.ObjectID `bson:"user_id" json:"user_id"`
	VideoId primitive.ObjectID `bson:"video_id" json:"_id"`
}
