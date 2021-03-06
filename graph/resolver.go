package graph

import (
	"tkame123-net/worldcup-gq-server/adapter"
	"tkame123-net/worldcup-gq-server/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	competitions     []*model.Competition
	players          []*model.Player
	MongoCompetition adapter.MongodbCompetitionRepository
	MongoPlayer      adapter.MongodbPlayerRepository
	MongoMatch       adapter.MongodbMatchRepository
}
