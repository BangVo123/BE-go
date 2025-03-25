package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Content  string             `bson:"content" json:"content"`
	Like     int                `bson:"like" json:"like"`
	BelongTo primitive.ObjectID `bson:"belong_to" json:"belong_to"`
	Sender   primitive.ObjectID `bson:"sender" json:"sender"`
	Parent   primitive.ObjectID `bson:"parent" json:"parent"`
}
