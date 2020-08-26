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
	list := make([]*int, 0, len(entity.RoundList))
	for _, v := range entity.RoundList {
		list = append(list, &v)
	}

	return &model.Player{
		Name:      entity.Name,
		RoundList: list,
	}
}

func ToMatchResponse(entity *domain.Match) *model.Match {
	return &model.Match{
		Stage: entity.Stage,
	}
}
