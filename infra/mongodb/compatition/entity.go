package compatition

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"tkame123-net/worldcup-gq-server/domain"
)

type entity struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Year    int                `bson:"Year" json:"Year"`
	Country string             `bson:"Country" json:"Country"`
}

const collection = "compatition"

func (e *entity) toDomain() *domain.Competition {
	return &domain.Competition{
		ID:      domain.CompetitionID(e.ID.Hex()),
		Year:    strconv.Itoa(e.Year),
		Country: e.Country,
	}
}

func toEntity(from *domain.Competition) *entity {
	id, err := primitive.ObjectIDFromHex(string(from.ID))
	if err != nil {
		println(err)
	}
	year, _ := strconv.Atoi(from.Year)
	return &entity{
		ID:      id,
		Year:    year,
		Country: from.Country,
	}
}
