package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Follower struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	FollowingId primitive.ObjectID `bson:"following_id" json:"following_id"`
}

type FollowingInfo struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	FollowingId []User             `bson:"following_id" json:"following_id"`
}

type FollowerInfo struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	UserId      []User             `bson:"user_id" json:"user_id"`
	FollowingId primitive.ObjectID `bson:"following_id" json:"following_id"`
}
