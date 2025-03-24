package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Follower struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	FollowingId primitive.ObjectID `bson:"following_id" json:"following_id"`
}
