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
	//GetAllByPost(ctx context.Context, id *domain.PostID) ([]*domain.BlogPostComment, error)
	//Get(ctx context.Context, id *domain.BlogPostCommentID) (*domain.BlogPostComment, error)
	//Exists(ctx context.Context, id *domain.BlogPostCommentID) (bool, error)
	//Insert(ctx context.Context, comment *domain.BlogPostComment) error
	//Put(ctx context.Context, comment *domain.BlogPostComment) error
	//Delete(ctx context.Context, id *domain.BlogPostCommentID) error
}
