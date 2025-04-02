package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id              primitive.ObjectID `bson:"_id" json:"_id"`
	FullName        string             `bson:"full_name" json:"full_name"`
	NickName        string             `bson:"nick_name" json:"nick_name"`
	Email           string             `bson:"email" json:"email"`
	PhoneNumber     string             `bson:"phone_nums" json:"phone_nums"`
	Password        string             `bson:"password" json:"password"`
	Avatar          string             `bson:"avatar" json:"avatar"`
	Bio             string             `bson:"bio" json:"bio"`
	AccountId       string             `bson:"account_id" json:"account_id"`
	Provider        string             `bson:"provider" json:"provider"`
	Tick            bool               `bson:"tick" json:"tick"`
	FollowingCounts int                `bson:"followings_count" json:"followings_count"`
	FollowerCounts  int                `bson:"followers_count" json:"followers_count"`
	LikeCounts      int                `bson:"likes_count" json:"likes_count"`
	WebsiteUrl      string             `bson:"website_URL" json:"website_URL"`
	FacebookUrl     string             `bson:"facebook_URL" json:"facebook_URL"`
	YoutubeUrl      string             `bson:"youtube_URL" json:"youtube_URL"`
	TwitterUrl      string             `bson:"twitter_URL" json:"twitter_URL"`
	InstagramUrl    string             `bson:"instagram_URL" json:"instagram_URL"`
}

type UserSummary struct {
	Id             primitive.ObjectID `bson:"_id" json:"_id"`
	FullName       string             `bson:"full_name" json:"full_name"`
	NickName       string             `bson:"nick_name" json:"nick_name"`
	Avatar         string             `bson:"avatar" json:"avatar"`
	Bio            string             `bson:"bio" json:"bio"`
	Tick           bool               `bson:"tick" json:"tick"`
	FollowerCounts int                `bson:"followers_count" json:"followers_count"`
	LikeCounts     int                `bson:"likes_count" json:"likes_count"`
}

func (us *User) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(us.Password), 10)

	if err != nil {
		return err
	}

	us.Password = string(hashPassword)
	return nil
}

func (us *User) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(password)); err != nil {
		return false
	}

	return true
}
