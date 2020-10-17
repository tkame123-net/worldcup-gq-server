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
	GetCursorsToEdges(ctx context.Context, after *string, before *string) ([]*domain.Competition, error)
	GetMultiByRange(ctx context.Context, limit *int, cursor *string, asc *bool) ([]*domain.Competition, error)
	GetAllByCountry(ctx context.Context, country *string) ([]*domain.Competition, error)
	Exists(ctx context.Context, id *domain.CompetitionID) (bool, error)
}

type MongodbMatchRepository interface {
	GetAll(ctx context.Context) ([]*domain.Match, error)
	GetAllByYear(ctx context.Context, year string, filterType domain.FilterType) ([]*domain.Match, error)
	Exists(ctx context.Context, id *domain.MatchID) (bool, error)
}

type MongodbPlayerRepository interface {
	GetAll(ctx context.Context) ([]*domain.Player, error)
	GetAllByPlayerName(ctx context.Context, playerName string, filterType domain.FilterType) ([]*domain.Player, error)
}
