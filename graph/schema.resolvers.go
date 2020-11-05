package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	"tkame123-net/worldcup-gq-server/domain"
	"tkame123-net/worldcup-gq-server/graph/generated"
	"tkame123-net/worldcup-gq-server/graph/model"
)

func (r *queryResolver) AllCompetition(ctx context.Context, first *int, last *int, after *string, before *string) (*model.CompetitionConnection, error) {
	// allEdges
	ctx = context.Background()
	allCompetitions, err := r.MongoCompetition.GetAll(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	a := make([]domain.Competition, 0, len(allCompetitions))
	for _, edge := range allCompetitions {
		e := *edge
		a = append(a, e)
	}

	// CursorsToEdge/EdgesToReturn
	edges, err := EdgesToReturn(a, before, after, first, last)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// step3: PageInfoの生成
	hasNextPage, err := HasNextPage(a, before, after, first, last)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	hasPreviousPage, err := HasPreviousPage(a, before, after, first, last)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	var startCursor, endCursor string
	if len(edges) > 0 {
		startCursor = string(edges[0].ID)
		endCursor = string(edges[len(edges)-1].ID)
	}
	pageInfo := model.PageInfo{
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
	}

	// competitionEdges
	competitionEdges := make([]*model.CompetitionEdge, 0, len(edges))
	for _, edge := range edges {
		competitionEdges = append(competitionEdges, ToCompetitionEdgeResponse(&edge))
	}

	// competitionConnection に変換
	competitionConnection := ToCompetitionConnectionResponse(competitionEdges, &pageInfo)

	return competitionConnection, nil
}

func (r *queryResolver) AllMatch(ctx context.Context, filterYear model.Filter) ([]*model.Match, error) {
	ctx = context.Background()
	if filterYear.Eq != "" {
		res, err := r.MongoMatch.GetAllByYear(ctx, filterYear.Eq, domain.FilterType_EQ)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		resItems := make([]*model.Match, 0, len(res))
		for _, item := range res {
			resItems = append(resItems, ToMatchResponse(item))
		}

		return resItems, nil
	}

	if filterYear.Regex != "" {
		res, err := r.MongoMatch.GetAllByYear(ctx, filterYear.Regex, domain.FilterType_REGEX)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		resItems := make([]*model.Match, 0, len(res))
		for _, item := range res {
			resItems = append(resItems, ToMatchResponse(item))
		}

		return resItems, nil
	}

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

func (r *queryResolver) AllPlayer(ctx context.Context, filterPlayerName model.Filter) ([]*model.Player, error) {
	ctx = context.Background()

	if filterPlayerName.Eq != "" {
		res, err := r.MongoPlayer.GetAllByPlayerName(ctx, filterPlayerName.Eq, domain.FilterType_EQ)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		resItems := make([]*model.Player, 0, len(res))
		for _, item := range res {
			resItems = append(resItems, ToPlayerResponse(item))
		}

		return resItems, nil
	}

	if filterPlayerName.Regex != "" {
		res, err := r.MongoPlayer.GetAllByPlayerName(ctx, filterPlayerName.Regex, domain.FilterType_REGEX)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		resItems := make([]*model.Player, 0, len(res))
		for _, item := range res {
			resItems = append(resItems, ToPlayerResponse(item))
		}

		return resItems, nil
	}

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

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	globalID, err := DecodeGlobalID(id)
	if err != nil {
		return nil, err
	}

	switch globalID.EntityName {
	case "Competition":
		oid := domain.CompetitionID(globalID.ID)
		b, err := r.MongoCompetition.Exists(ctx, &(oid))
		if err != nil {
			return nil, err
		}
		if b == false {
			return nil, errors.New("not found")
		}
		return model.Competition{ID: id}, nil

	case "Match":
		oid := domain.MatchID(globalID.ID)
		b, err := r.MongoMatch.Exists(ctx, &(oid))
		if err != nil {
			return nil, err
		}
		if b == false {
			return nil, errors.New("not found")
		}

		return model.Match{ID: id}, nil

	default:

	}

	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
