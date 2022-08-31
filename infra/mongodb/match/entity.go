package match

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tkame123-net/worldcup-gq-server/domain"
)

type entity struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Year    int                `bson:"Year" json:"Year"`
	Stage   string             `bson:"Stage" json:"Stage"`
	Stadium string             `bson:"Stadium" json:"Stadium"`
	City    string             `bson:"City" json:"City"`
}

const collection = "matches"

func (e *entity) toDomain() *domain.Match {
	return &domain.Match{
		ID:      domain.MatchID(e.ID.Hex()),
		Year:    e.Year,
		Stage:   e.Stage,
		Stadium: e.Stadium,
		City:    e.City,
	}
}

func toEntity(from *domain.Match) *entity {
	id, err := primitive.ObjectIDFromHex(string(from.ID))
	if err != nil {
		println(err)
	}
	//year, _ := strconv.Atoi(from.Year)
	return &entity{
		ID:      id,
		Year:    from.Year,
		Stage:   from.Stage,
		Stadium: from.Stadium,
		City:    from.City,
	}
}
