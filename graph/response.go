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
	idList := make([]*int, 0, len(entity.MatchIDList))
	for _, v := range entity.MatchIDList {
		x := v
		idList = append(idList, &x)
	}

	mList := make([]*model.Match, 0, len(entity.MatchList))
	for _, v := range entity.MatchList {
		x := v
		mList = append(mList, &model.Match{Year: x.Year, Stage: x.Stage, Stadium: x.Stadium, City: x.City})
	}

	return &model.Player{
		Name:        entity.Name,
		MatchIDList: idList,
		MatchList:   mList,
	}
}

func ToMatchResponse(entity *domain.Match) *model.Match {
	return &model.Match{
		Stage: entity.Stage,
	}
}
