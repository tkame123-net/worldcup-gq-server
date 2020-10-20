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

	var after, before string

	// 1 afterが存在して beforeが存在しない場合
	after = "3"
	edges, err := ApplyCursorsToEdges(allEdges, nil, &after)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	assert.Equal(t, 7, len(edges), "case1")

	// 2 afterが存在して beforeが存在する場合
	after = "3"
	before = "7"
	edges, err = ApplyCursorsToEdges(allEdges, &before, &after)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	assert.Equal(t, 3, len(edges), "case2")

	// 3 afterが存在せず beforeが存在しない場合
	edges, err = ApplyCursorsToEdges(allEdges, nil, nil)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	assert.Equal(t, 10, len(edges), "case3")

	// 4 afterが存在せず beforeが存在する場合
	before = "7"
	edges, err = ApplyCursorsToEdges(allEdges, &before, nil)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	assert.Equal(t, 6, len(edges), "case4")

	// 5 存在しないafterを指定した場合
	after = "11"
	edges, err = ApplyCursorsToEdges(allEdges, nil, &after)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	assert.Equal(t, 10, len(edges), "case5")

	// 6 存在しないbeforeを指定した場合
	before = "11"
	edges, err = ApplyCursorsToEdges(allEdges, &before, nil)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	assert.Equal(t, 10, len(edges), "case6")

}

func TestEdgesToReturn(t *testing.T) {
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

	// テストケース
	// before 指定無し, after 指定無し、 first 指定無し、 last 指定無し
	// before 指定無し, after 指定無し、 first 存在する値を指定、 last 指定無し
	// before 指定無し, after 指定無し、 first 指定無し、 last 存在する値を指定
	// before 指定無し, after 指定無し、 first 存在する値を指定、 last 存在する値を指定
	var testCase1 = []struct {
		pt               string
		before           string
		after            string
		first            int
		last             int
		out              int    // 戻り値のlen値
		outCursorOfFirst string // 戻りの値の先頭の値のCursor
		outCursorOfEnd   string // 戻りの値の最後の値のCursor
	}{
		{"case1", "", "", 0, 0, 10, "1", "10"},
		{"case2", "", "", 1, 0, 1, "1", "1"},
		{"case3", "", "", 0, 1, 1, "10", "10"},
		{"case4", "", "", 4, 2, 2, "3", "4"},
	}

	t.Run("case1: before nil/ after nil", func(t *testing.T) {
		a := make([]*domain.Competition, 0, len(allEdges))
		for _, edge := range allEdges {
			e := edge
			a = append(a, &e)
		}

		for _, tt := range testCase1 {
			var before, after *string
			var first, last *int

			if tt.before != "" {
				before = &tt.before
			}
			if tt.after != "" {
				after = &tt.after
			}
			if tt.first != 0 {
				first = &tt.first
			}
			if tt.last != 0 {
				last = &tt.last
			}

			edges, err := EdgesToReturn(a, before, after, first, last)
			if err != nil {
				t.Fatalf("error: %v\n", err)
			}
			assert.Equal(t, tt.out, len(edges), tt.pt+":01")
			assert.Equal(t, tt.outCursorOfFirst, string(edges[0].ID), tt.pt+":02")
			assert.Equal(t, tt.outCursorOfEnd, string(edges[len(edges)-1].ID), tt.pt+":03")
		}
	})

	// before 存在する値を指定, after 指定無し、 first 指定無し、 last 指定無し
	// before 存在する値を指定, after 指定無し、 first 存在する値を指定、 last 指定無し
	// before 存在する値を指定, after 指定無し、 first 指定無し、 last 存在する値を指定
	// before 存在する値を指定, after 指定無し、 first 存在する値を指定、 last 存在する値を指定

	// before 指定無し, after 存在する値を指定、 first 指定無し、 last 指定無し
	// before 指定無し, after 存在する値を指定、 first 存在する値を指定、 last 指定無し
	// before 指定無し, after 存在する値を指定、 first 指定無し、 last 存在する値を指定
	// before 指定無し, after 存在する値を指定、 first 存在する値を指定、 last 存在する値を指定

	// before 存在する値を指定, after 存在する値を指定、 first 指定無し、 last 指定無し
	// before 存在する値を指定, after 存在する値を指定、 first 存在する値を指定、 last 指定無し
	// before 存在する値を指定, after 存在する値を指定、 first 指定無し、 last 存在する値を指定
	// before 存在する値を指定, after 存在する値を指定、 first 存在する値を指定、 last 存在する値を指定

}
