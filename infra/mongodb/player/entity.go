package player

import (
	"tkame123-net/worldcup-gq-server/domain"
)

type entity struct {
	ID        string `bson:"_id,omitempty"`
	RoundList []int  `bson:"RoundList" json:"RoundList"`
}

const collection = "players"

func (e *entity) toDomain() *domain.Player {
	return &domain.Player{
		Name:      e.ID,
		RoundList: e.RoundList,
	}
}

func toEntity(from *domain.Player) *entity {
	return &entity{
		ID:        from.Name,
		RoundList: from.RoundList,
	}
}
