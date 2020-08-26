package adapter

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"tkame123-net/worldcup-gq-server/domain"
)

type MongoDBAPI struct {
	MongoDBHostURL string
	Database       string
	Client         *mongo.Client
}

type MongoClient func(ctx context.Context) (*MongoDBAPI, error)

type MongodbCompetitionRepository interface {
	GetByYear(ctx context.Context, year *int) (*domain.Competition, error)
	GetAll(ctx context.Context) ([]*domain.Competition, error)
	GetAllByCountry(ctx context.Context, country *string) ([]*domain.Competition, error)
}

type MongodbPlayerRepository interface {
	GetAll(ctx context.Context) ([]*domain.Player, error)
}

type MongodbMatchRepository interface {
	GetAll(ctx context.Context) ([]*domain.Match, error)
}
