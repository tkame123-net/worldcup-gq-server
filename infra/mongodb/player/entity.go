package player

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tkame123-net/worldcup-gq-server/domain"
)

type entity struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"Player Name" json:"Player Name"`
}

const collection = "players"

func (e *entity) toDomain() *domain.Player {
	return &domain.Player{
		ID:   domain.PlayerID(e.ID.Hex()),
		Name: e.Name,
	}
}

func toEntity(from *domain.Player) *entity {
	id, err := primitive.ObjectIDFromHex(string(from.ID))
	if err != nil {
		println(err)
	}
	return &entity{
		ID:   id,
		Name: from.Name,
	}
}
