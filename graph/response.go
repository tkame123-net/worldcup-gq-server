package graph

import (
	"tkame123-net/worldcup-gq-server/domain"
	"tkame123-net/worldcup-gq-server/graph/model"
)

func ToCompetitionResponse(entity *domain.Competition) *model.Competition {
	return &model.Competition{
		Year:    entity.Year,
		Country: entity.Country,
	}
}

func ToPlayerResponse(entity *domain.Player) *model.Player {
	return &model.Player{
		Name: entity.Name,
	}
}

func ToMatchResponse(entity *domain.Match) *model.Match {
	return &model.Match{
		Stage: entity.Stage,
	}
}
