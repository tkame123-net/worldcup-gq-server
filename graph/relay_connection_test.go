package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tkame123-net/worldcup-gq-server/domain"
)

// ApplyCursorsToEdges
func TestApplyCursorsToEdges(t *testing.T) {
	allEdges := make([]domain.Competition, 0, 10)
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("1"),
		Year:    "1938",
		Country: "",
	})
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("2"),
		Year:    "1942",
		Country: "",
	})
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("3"),
		Year:    "1946",
		Country: "",
	})
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("4"),
		Year:    "1950",
		Country: "",
	})
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("5"),
		Year:    "1954",
		Country: "",
	})
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("6"),
		Year:    "1958",
		Country: "",
	})
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("7"),
		Year:    "1962",
		Country: "",
	})
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("8"),
		Year:    "1966",
		Country: "",
	})
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("9"),
		Year:    "1970",
		Country: "",
	})
	allEdges = append(allEdges, domain.Competition{
		ID:      domain.CompetitionID("10"),
		Year:    "1974",
		Country: "",
	})

	after := "3"

	// 1 afterが存在して beforeが存在しない場合
	edges, err := ApplyCursorsToEdges(allEdges, nil, &after)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	assert.Equal(t, len(edges), 7)
}

// 2 afterが存在して beforeが存在する場合
// 3 afterが存在せず beforeが存在しない場合
// 4 afterが存在せず beforeが存在する場合
// 5 ありえないbeforeが入っている場合
// 6 ありえないafterが入っている場合
