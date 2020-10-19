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

// HasNextPage
func TestHasNextPage(t *testing.T) {
	allEdges := make([]*domain.Competition, 0, 10)
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("1"),
		Year:    "1938",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("2"),
		Year:    "1942",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("3"),
		Year:    "1946",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("4"),
		Year:    "1950",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("5"),
		Year:    "1954",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("6"),
		Year:    "1958",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("7"),
		Year:    "1962",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("8"),
		Year:    "1966",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("9"),
		Year:    "1970",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("10"),
		Year:    "1974",
		Country: "",
	})

	var before string

	var first int

	// 1 firstが存在して、beforeが存在しない場合
	t.Run("case1", func(t *testing.T) {
		first = 5
		hasNextPage, err := HasNextPage(allEdges, nil, nil, &first, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case1")
	})
	// 2 firstが存在して、beforeが存在する場合 *おかしい
	t.Run("case2", func(t *testing.T) {
		first = 11
		before = "8"
		hasNextPage, err := HasNextPage(allEdges, &before, nil, &first, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case2")
	})
	// 3 firstが存在せず、beforeが存在しない場合
	t.Run("case3", func(t *testing.T) {
		hasNextPage, err := HasNextPage(allEdges, nil, nil, nil, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, false, hasNextPage, "case3")
	})
	// 4 firstが存在せず、beforeが存在する場合　*おかしい
	t.Run("case4", func(t *testing.T) {
		before = "5"
		hasNextPage, err := HasNextPage(allEdges, &before, nil, nil, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case4")
	})

}

// HasNextPage
func TestHasPreviousPage(t *testing.T) {
	allEdges := make([]*domain.Competition, 0, 10)
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("1"),
		Year:    "1938",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("2"),
		Year:    "1942",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("3"),
		Year:    "1946",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("4"),
		Year:    "1950",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("5"),
		Year:    "1954",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("6"),
		Year:    "1958",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("7"),
		Year:    "1962",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("8"),
		Year:    "1966",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("9"),
		Year:    "1970",
		Country: "",
	})
	allEdges = append(allEdges, &domain.Competition{
		ID:      domain.CompetitionID("10"),
		Year:    "1974",
		Country: "",
	})

	var after string

	var last int

	// 1 lastが存在して、afterが存在しない場合
	t.Run("case1", func(t *testing.T) {
		last = 5
		hasNextPage, err := HasPreviousPage(allEdges, nil, nil, nil, &last)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case1")
	})
	// 2 lastが存在して、afterが存在する場合 *おかしい
	t.Run("case2", func(t *testing.T) {
		last = 5
		after = "11"
		hasNextPage, err := HasPreviousPage(allEdges, nil, &after, nil, &last)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case2")
	})
	// 3 lastが存在せず、afterが存在しない場合
	t.Run("case3", func(t *testing.T) {
		hasNextPage, err := HasPreviousPage(allEdges, nil, nil, nil, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, false, hasNextPage, "case3")
	})
	// 4 lastが存在せず、afterが存在する場合　*おかしい
	t.Run("case4", func(t *testing.T) {
		after = "5"
		hasNextPage, err := HasPreviousPage(allEdges, nil, &after, nil, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case4")
	})

}
