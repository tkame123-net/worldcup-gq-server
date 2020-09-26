package graph

import (
	b64 "encoding/base64"
	"errors"
	"strings"
	"tkame123-net/worldcup-gq-server/domain"
	"tkame123-net/worldcup-gq-server/graph/model"
)

func ToCompetitionResponse(entity *domain.Competition) *model.Competition {
	id := ToGlobalID("01", "Competition", string(entity.ID))
	return &model.Competition{
		ID:      id,
		Year:    entity.Year,
		Country: entity.Country,
	}
}

func ToPlayerResponse(entity *domain.Player) *model.Player {
	mList := make([]*model.Match, 0, len(entity.MatchList))
	for _, v := range entity.MatchList {
		x := v
		mList = append(mList, &model.Match{Year: x.Year, Stage: x.Stage, Stadium: x.Stadium, City: x.City})
	}

	return &model.Player{
		Name:      entity.Name,
		MatchList: mList,
	}
}

func ToMatchResponse(entity *domain.Match) *model.Match {
	id := ToGlobalID("01", "Match", string(entity.ID))
	return &model.Match{
		ID:      id,
		Year:    entity.Year,
		Stage:   entity.Stage,
		Stadium: entity.Stadium,
		City:    entity.City,
	}
}

func ToGlobalID(verStr string, entityName string, id string) string {
	idSlices := []string{verStr, entityName, id}
	idStr := strings.Join(idSlices, ":")
	idEncode := b64.StdEncoding.EncodeToString([]byte(idStr))

	return idEncode
}

func DecodeGlobalID(id string) (domain.GlobalID, error) {
	dec, _ := b64.StdEncoding.DecodeString(id)
	slice := strings.Split(string(dec), ":")
	if len(slice) < 3 {
		return domain.GlobalID{}, errors.New("invalid data")
	}
	return domain.GlobalID{VerStr: slice[0], EntityName: slice[1], ID: slice[2]}, nil

}
