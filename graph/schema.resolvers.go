package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strconv"
	"tkame123-net/worldcup-gq-server/graph/generated"
	"tkame123-net/worldcup-gq-server/graph/model"
)

func (r *competitionResolver) PrevCompetition(ctx context.Context, obj *model.Competition) (*model.Competition, error) {
	i, err := strconv.Atoi(obj.Year)
	if err != nil {
		return nil, nil
	} else {
		i = i - 4
	}

	ctx = context.Background()
	res, err := r.MongoCompetition.GetByYear(ctx, &i)
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

	ctx = context.Background()
	res, err := r.MongoCompetition.GetByYear(ctx, &i)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if res == nil {
		return nil, nil
	}

	return ToCompetitionResponse(res), nil
}

func (r *queryResolver) AllCompetition(ctx context.Context) ([]*model.Competition, error) {
	ctx = context.Background()
	res, err := r.MongoCompetition.GetAll(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	resItems := make([]*model.Competition, 0, len(res))
	for _, item := range res {
		resItems = append(resItems, ToCompetitionResponse(item))
	}

	return resItems, nil
}

func (r *queryResolver) AllMatch(ctx context.Context) ([]*model.Match, error) {
	ctx = context.Background()
	res, err := r.MongoMatch.GetAll(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	resItems := make([]*model.Match, 0, len(res))
	for _, item := range res {
		resItems = append(resItems, ToMatchResponse(item))
	}

	return resItems, nil
}

func (r *queryResolver) AllPlayer(ctx context.Context) ([]*model.Player, error) {
	ctx = context.Background()
	res, err := r.MongoPlayer.GetAll(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	resItems := make([]*model.Player, 0, len(res))
	for _, item := range res {
		resItems = append(resItems, ToPlayerResponse(item))
	}

	return resItems, nil
}

// Competition returns generated.CompetitionResolver implementation.
func (r *Resolver) Competition() generated.CompetitionResolver { return &competitionResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type competitionResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
