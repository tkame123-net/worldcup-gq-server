package match

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tkame123-net/worldcup-gq-server/domain"
)

type entity struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Stage string             `bson:"Stage" json:"Stage"`
}

const collection = "matches"

func (e *entity) toDomain() *domain.Match {
	return &domain.Match{
		ID:    domain.MatchID(e.ID.Hex()),
		Stage: e.Stage,
	}
}

func toEntity(from *domain.Match) *entity {
	id, err := primitive.ObjectIDFromHex(string(from.ID))
	if err != nil {
		println(err)
	}
	//year, _ := strconv.Atoi(from.Year)
	return &entity{
		ID:    id,
		Stage: from.Stage,
	}
}
