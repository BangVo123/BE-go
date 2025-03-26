package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Url      string             `bson:"url" json:"url"`
	Size     int                `bson:"size" json:"size"`
	Duration string             `bson:"duration" json:"duration"`
	UploadAt time.Time          `bson:"upload_at" json:"upload_at"`
	BelongTo string             `bson:"belong_to" json:"belong_to"`
	Like     int                `bson:"like" json:"like"`
	Love     int                `bson:"love" json:"love"`
	Comment  int                `bson:"comment" json:"comment"`
	Content  string             `bson:"content" json:"content"`
	Share    int                `bson:"share" json:"share"`
}
