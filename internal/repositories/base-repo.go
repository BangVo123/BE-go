package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepository[T any] struct {
	DB             *mongo.Database
	CollectionName string
	DbName         string
}

func NewBaseRepo[T any](db *mongo.Database, collectionName string) GenericRepository[T] {
	return &BaseRepository[T]{DB: db, CollectionName: collectionName}
}

func (r *BaseRepository[T]) GetCollection() any {
	Collection := r.DB.Collection(r.CollectionName)
	return Collection
}

func (r *BaseRepository[T]) GetById(ctx context.Context, id string) (*T, error) {
	Collection := r.DB.Collection(r.CollectionName)
	var result T

	ObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = Collection.FindOne(ctx, bson.M{"_id": ObjectId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *BaseRepository[T]) GetByField(ctx context.Context, field string, value any) (*T, error) {
	Collection := r.DB.Collection(r.CollectionName)
	var result T

	err := Collection.FindOne(ctx, bson.M{field: value, "$or": []bson.M{{"provider": nil}, {"provider": ""}}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *BaseRepository[T]) GetByFilter(ctx context.Context, filter map[string]any) (*T, error) {
	Collection := r.DB.Collection(r.CollectionName)
	var result T

	err := Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *BaseRepository[T]) Create(ctx context.Context, payload *T) error {
	Collection := r.DB.Collection(r.CollectionName)
	_, err := Collection.InsertOne(ctx, payload)
	return err
}

func (r *BaseRepository[T]) Update(ctx context.Context, id string, payload map[string]any) (*T, error) {
	Collection := r.DB.Collection(r.CollectionName)

	ObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result T
	update := bson.M{"$set": payload}
	err = Collection.FindOneAndUpdate(ctx, bson.M{"_id": ObjectId}, update).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id string) error {
	Collection := r.DB.Collection(r.CollectionName)

	ObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = Collection.DeleteOne(ctx, bson.M{"_id": ObjectId})
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) DeleteByFilter(ctx context.Context, filter map[string]any) error {
	Collection := r.DB.Collection(r.CollectionName)

	_, err := Collection.DeleteMany(ctx, filter)
	return err
}

func (r *BaseRepository[T]) GetMany(ctx context.Context, filter map[string]any, pagination map[string]int64) (*[]T, error) {
	Collection := r.DB.Collection(r.CollectionName)

	page := pagination["page"]
	limit := pagination["limit"]
	offset := (page - 1) * limit

	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := Collection.Find(ctx, bson.M(filter), findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var entities []T
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *BaseRepository[T]) GetAll(ctx context.Context, pagination map[string]int64) (*[]T, error) {
	Collection := r.DB.Collection(r.CollectionName)

	page := pagination["page"]
	limit := pagination["limit"]
	offset := (page - 1) * limit

	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := Collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var entities []T
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *BaseRepository[T]) GetWithPopulation(ctx context.Context, pagination map[string]int64, pipelineValue map[string]any, unwindAttr string) (*[]T, error) {
	Collection := r.DB.Collection(r.CollectionName)

	page := pagination["page"]
	limit := pagination["limit"]
	offset := (page - 1) * limit

	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: pipelineValue}},
		{{Key: "$skip", Value: offset}},
		{{Key: "$limit", Value: limit}},
	}

	if unwindAttr != "" {
		pipeline = append(pipeline, bson.D{{Key: "$unwind", Value: unwindAttr}})
	}

	cursor, err := Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var entities []T
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}
	return &entities, nil
}
