package services

import (
	"context"
	"project/internal/models"
	"project/internal/repositories"
	"project/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeService struct {
	LikeRepo *repositories.LikeRepo
}

func NewLikeService(LikeRepo *repositories.LikeRepo) usecase.LikeUseCase {
	return &LikeService{LikeRepo: LikeRepo}
}

func (ls *LikeService) GetLike(ctx context.Context, userId primitive.ObjectID) ([]any, error) {
	return ls.LikeRepo.GetDistinct(ctx, "video_id", map[string]any{"user_id": userId})
}

func (ls *LikeService) Like(ctx context.Context, videoIdString, userIdString string) error {
	videoId, err := primitive.ObjectIDFromHex(videoIdString)
	if err != nil {
		return err
	}
	userId, err := primitive.ObjectIDFromHex(userIdString)
	if err != nil {
		return err
	}

	emotion, err := ls.LikeRepo.GetByFilter(ctx, map[string]any{"video_id": videoId, "user_id": userId})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ls.LikeRepo.Create(ctx, &models.Emotion{Id: primitive.NewObjectID(), UserId: userId, VideoId: videoId})
		} else {
			return err
		}
	}

	if emotion != nil {
		return ls.LikeRepo.Delete(ctx, emotion.Id.Hex())
	}

	return nil
}

type LoveService struct {
	LoveRepo *repositories.LoveRepo
}

func NewLoveService(LoveRepo *repositories.LoveRepo) usecase.LoveUseCase {
	return &LoveService{LoveRepo: LoveRepo}
}

func (ls *LoveService) GetLove(ctx context.Context, userId primitive.ObjectID) ([]any, error) {
	return ls.LoveRepo.GetDistinct(ctx, "video_id", map[string]any{"user_id": userId})
}

func (ls *LoveService) Love(ctx context.Context, videoIdString, userIdString string) error {
	videoId, err := primitive.ObjectIDFromHex(videoIdString)
	if err != nil {
		return err
	}
	userId, err := primitive.ObjectIDFromHex(userIdString)
	if err != nil {
		return err
	}

	emotion, err := ls.LoveRepo.GetByFilter(ctx, map[string]any{"video_id": videoId, "user_id": userId})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ls.LoveRepo.Create(ctx, &models.Emotion{Id: primitive.NewObjectID(), UserId: userId, VideoId: videoId})
		} else {
			return err
		}
	}

	if emotion != nil {
		return ls.LoveRepo.Delete(ctx, emotion.Id.Hex())
	}

	return nil
}
