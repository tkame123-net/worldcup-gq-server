package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"os"
	"strconv"
	"tkame123-net/worldcup-gq-server/graph/generated"
	"tkame123-net/worldcup-gq-server/graph/model"
	"tkame123-net/worldcup-gq-server/infra/mongodb"
	"tkame123-net/worldcup-gq-server/infra/mongodb/competition"
)

func (r *competitionResolver) PrevCompetition(ctx context.Context, obj *model.Competition) (*model.Competition, error) {
	i, err := strconv.Atoi(obj.Year)
	if err != nil {
		return nil, nil
	} else {
		i = i - 4
	}

	// todo: to wire
	cl := mongodb.NewClient(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
	repo := competition.NewRepository(cl)

	ctx = context.Background()
	res, err := repo.GetByYear(ctx, &i)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if res == nil {
		return nil, nil
	}

	return ToCompetitionResponse(res), nil
}

func (r *competitionResolver) NextCompetition(ctx context.Context, obj *model.Competition) (*model.Competition, error) {
	i, err := strconv.Atoi(obj.Year)
	if err != nil {
		return nil, nil
	} else {
		i = i + 4
	}

	// todo: to wire
	cl := mongodb.NewClient(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
	repo := competition.NewRepository(cl)

	ctx = context.Background()
	res, err := repo.GetByYear(ctx, &i)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if res == nil {
		return nil, nil
	}

	return ToCompetitionResponse(res), nil
}

func (r *queryResolver) AllCompetition(ctx context.Context) ([]*model.Competition, error) {
	// todo: to wire
	cl := mongodb.NewClient(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
	repo := competition.NewRepository(cl)

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

// Competition returns generated.CompetitionResolver implementation.
func (r *Resolver) Competition() generated.CompetitionResolver { return &competitionResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type competitionResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
