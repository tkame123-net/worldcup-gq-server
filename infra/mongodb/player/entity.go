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
	Year    int    `bson:"Year" json:"Year"`
	Stage   string `bson:"Stage" json:"Stage"`
	Stadium string `bson:"Stadium" json:"Stadium"`
	City    string `bson:"City" json:"City"`
}

const collection = "players"

func (e *entity) toDomain() *domain.Player {
	list := make([]domain.Match, 0, len(e.MatchList))
	for _, v := range e.MatchList {
		list = append(list, domain.Match{Year: v.Year, Stage: v.Stage, Stadium: v.Stadium, City: v.City})
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
