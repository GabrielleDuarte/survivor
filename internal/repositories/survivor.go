package repositories

import (
	"context"
	"dataapi/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SurvivorRepository struct {
	collection *mongo.Collection
}

func NewSurvivorRepository(collection *mongo.Collection) *SurvivorRepository {
	return &SurvivorRepository{
		collection: collection,
	}
}

func (r *SurvivorRepository) Create(ctx context.Context, survivor models.Survivor) error {
	_, err := r.collection.InsertOne(ctx, survivor)
	return err
}
