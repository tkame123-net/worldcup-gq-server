package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"os"
	"tkame123-net/worldcup-gq-server/graph/generated"
	"tkame123-net/worldcup-gq-server/graph/model"
	"tkame123-net/worldcup-gq-server/infra/mongodb"
	"tkame123-net/worldcup-gq-server/infra/mongodb/compatition"
)

func (r *queryResolver) AllCompetition(ctx context.Context) ([]*model.Competition, error) {
	// todo: to wire
	cl := mongodb.NewClient(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
	repo := compatition.NewRepository(cl)

	ctx = context.Background()
	res, err := repo.GetAll(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	resItems := make([]*model.Competition, 0, len(res))
	for _, item := range res {
		resItems = append(resItems, ToCompetitionResponse(item))
	}

	return resItems, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
