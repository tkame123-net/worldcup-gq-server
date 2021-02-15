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
	nodes, err := EdgesToReturn(a, before, after, first, last)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// domain.Nodeからdomain.Competitionへの型キャスト
	edges := make([]domain.Competition, 0, len(nodes))
	for _, v := range nodes {
		edges = append(edges, v.(domain.Competition))
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

func (r *queryResolver) AllMatch(ctx context.Context, first *int, last *int, after *string, before *string) (*model.MatchConnection, error) {
	ctx = context.Background()
	allMatch, err := r.MongoMatch.GetAll(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	a := make([]domain.Match, 0, len(allMatch))
	for _, edge := range allMatch {
		e := *edge
		a = append(a, e)
	}

	// CursorsToEdge/EdgesToReturn
	nodes, err := EdgesToReturn(a, before, after, first, last)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// domain.Nodeからdomain.Competitionへの型キャスト
	edges := make([]domain.Match, 0, len(nodes))
	for _, v := range nodes {
		edges = append(edges, v.(domain.Match))
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

	// matchEdges
	matchEdges := make([]*model.MatchEdge, 0, len(edges))
	for _, edge := range edges {
		matchEdges = append(matchEdges, ToMatchEdgeResponse(&edge))
	}

	// matchConnection に変換
	matchConnection := ToMatchConnectionResponse(matchEdges, &pageInfo)

	return matchConnection, nil
}

func (r *queryResolver) AllPlayer(ctx context.Context, first *int, last *int, after *string, before *string) (*model.PlayerConnection, error) {
	ctx = context.Background()
	allPlayer, err := r.MongoPlayer.GetAll(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	a := make([]domain.Player, 0, len(allPlayer))
	for _, edge := range allPlayer {
		e := *edge
		a = append(a, e)
	}

	// CursorsToEdge/EdgesToReturn
	nodes, err := EdgesToReturn(a, before, after, first, last)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// domain.Nodeからdomain.Competitionへの型キャスト
	edges := make([]domain.Player, 0, len(nodes))
	for _, v := range nodes {
		edges = append(edges, v.(domain.Player))
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

	// playerEdges
	playerEdges := make([]*model.PlayerEdge, 0, len(edges))
	for _, edge := range edges {
		playerEdges = append(playerEdges, ToPlayerEdgeResponse(&edge))
	}

	// playerConnection に変換
	matchConnection := ToPlayerConnectionResponse(playerEdges, &pageInfo)

	return matchConnection, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	globalID, err := DecodeGlobalID(id)
	if err != nil {
		return nil, err
	}

	switch globalID.EntityName {
	case "Competition":
		oid := domain.CompetitionID(globalID.ID)
		b, err := r.MongoCompetition.Get(ctx, oid)
		if err != nil {
			return nil, err
		}
		return model.Competition{ID: id, Year: b.Year, Country: b.Country}, nil

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
