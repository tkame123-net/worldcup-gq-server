package player

import (
	"tkame123-net/worldcup-gq-server/domain"
)

type entity struct {
	ID          string        `bson:"_id,omitempty"`
	MatchIDList []int         `bson:"MatchIDList" json:"MatchIDList"`
	MatchList   []matchEntity `bson:"MatchList" json:"MatchList"`
}

type matchEntity struct {
	Stage string `bson:"Stage" json:"Stage"`
}

const collection = "players"

func (e *entity) toDomain() *domain.Player {
	list := make([]domain.Match, 0, len(e.MatchList))
	for _, v := range e.MatchList {
		list = append(list, domain.Match{Stage: v.Stage})
	}

	return &domain.Player{
		Name:        e.ID,
		MatchIDList: e.MatchIDList,
		MatchList:   list,
	}
}

func toEntity(from *domain.Player) *entity {
	return &entity{
		ID:          from.Name,
		MatchIDList: from.MatchIDList,
	}
}
