package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoSessionStore struct {
	collection *mongo.Collection
}

type Session struct {
	Id     primitive.ObjectID `bson:"_id" json:"_id"`
	UserId string             `bson:"user_id" json:"user_id"`
}

func NewMongoSessionStore(client *mongo.Client, dbName, collectionName string) *MongoSessionStore {
	return &MongoSessionStore{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (m *MongoSessionStore) Save(userId string) (string, error) {
	ctx := context.TODO()
	session, err := m.collection.InsertOne(ctx, bson.M{"user_id": userId})
	if err != nil {
		return "", err
	}

	oid, ok := session.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("something went wrong")
	}
	return oid.Hex(), nil
}

func (m *MongoSessionStore) Load(sessionId string) (string, error) {
	ObjectId, err := primitive.ObjectIDFromHex(sessionId)
	if err != nil {
		return "", err
	}

	ctx := context.TODO()
	var result Session
	err = m.collection.FindOne(ctx, bson.M{"_id": ObjectId}).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.UserId, nil
}
