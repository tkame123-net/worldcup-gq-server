package graph

import (
	"tkame123-net/worldcup-gq-server/domain"
	"tkame123-net/worldcup-gq-server/graph/model"
)

func ToCompetitionResponse(competition *domain.Competition) *model.Competition {
	return &model.Competition{
		Year:    competition.Year,
		Country: competition.Country,
	}
}
