package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tkame123-net/worldcup-gq-server/domain"
)

// ApplyCursorsToEdges
func TestApplyCursorsToEdges(t *testing.T) {
	allEdges := createInstance()

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
	allEdges := createInstance()

	// 1 first 5 after 5 10件中 ５番目より後から５個取った場合 次のページは無い false
	t.Run("case1", func(t *testing.T) {
		first := 5
		after := "5"
		hasNextPage, err := HasNextPage(allEdges, nil, &after, &first, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, false, hasNextPage, "case1")
	})
	// 2 first 4 after 5 10件中 ５番目より後から４個取った場合 次のページはある　true
	t.Run("case2", func(t *testing.T) {
		first := 4
		after := "5"
		hasNextPage, err := HasNextPage(allEdges, nil, &after, &first, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case2")
	})
	// 3 first 6 after 5 10件中 ５番目より後から6個取った場合(取り切れない) 次のページはない　false
	t.Run("case3", func(t *testing.T) {
		first := 6
		after := "5"
		hasNextPage, err := HasNextPage(allEdges, nil, &after, &first, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, false, hasNextPage, "case3")
	})
	// 4 last 5 before 5 10件中 ５番目より前から後ろから５個取った場合　次のぺーじ　ある
	t.Run("case4", func(t *testing.T) {
		last := 5
		before := "5"
		hasNextPage, err := HasNextPage(allEdges, &before, nil, nil, &last)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case4")
	})

	// 5 last 4 before 5 10件中 ５番目より前から後ろから4個取った場合 次のページ　ある
	t.Run("case5", func(t *testing.T) {
		last := 4
		before := "5"
		hasNextPage, err := HasNextPage(allEdges, &before, nil, nil, &last)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case5")
	})

	// 5 last 3 before 5 10件中 ５番目より前から後ろから3個取った場合 次のページ　ある
	t.Run("case6", func(t *testing.T) {
		last := 3
		before := "5"
		hasNextPage, err := HasNextPage(allEdges, &before, nil, nil, &last)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasNextPage, "case6")
	})

}

// HasPreviousPage
func TestHasPreviousPage(t *testing.T) {
	allEdges := createInstance()

	// 1 first 5 after 5 10件中 ５番目より後から５個取った場合 前のページはある
	t.Run("case1", func(t *testing.T) {
		first := 5
		after := "5"
		hasPreviousPage, err := HasPreviousPage(allEdges, nil, &after, &first, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasPreviousPage, "case1")
	})
	// 2 first 4 after 5 10件中 ５番目より後から４個取った場合前のページはある
	t.Run("case2", func(t *testing.T) {
		first := 4
		after := "5"
		hasPreviousPage, err := HasPreviousPage(allEdges, nil, &after, &first, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasPreviousPage, "case2")
	})
	// 3 first 6 after 5 10件中 ５番目より後から6個取った場合(取り切れない) 前のページはある
	t.Run("case3", func(t *testing.T) {
		first := 6
		after := "5"
		hasPreviousPage, err := HasPreviousPage(allEdges, nil, &after, &first, nil)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasPreviousPage, "case3")
	})
	// 4 last 5 before 5 10件中 ５番目より前から後ろから５個取った場合　前のページはない
	t.Run("case4", func(t *testing.T) {
		last := 5
		before := "5"
		hasPreviousPage, err := HasPreviousPage(allEdges, &before, nil, nil, &last)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, false, hasPreviousPage, "case4")
	})
	// 5 last 4 before 5 10件中 ５番目より前から4個取った場合 前のページはない
	t.Run("case5", func(t *testing.T) {
		last := 4
		before := "5"
		hasPreviousPage, err := HasPreviousPage(allEdges, &before, nil, nil, &last)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, false, hasPreviousPage, "case5")
	})

	// 6 last 3 before 5 10件中 ５番目より前から3個取った場合 前のページはある
	t.Run("case6", func(t *testing.T) {
		last := 3
		before := "5"
		hasPreviousPage, err := HasPreviousPage(allEdges, &before, nil, nil, &last)
		if err != nil {
			t.Fatalf("error: %v\n", err)
		}
		assert.Equal(t, true, hasPreviousPage, "case6")
	})

}

func TestEdgesToReturn(t *testing.T) {
	allEdges := createInstance()

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
		{"step1", "", "", 0, 0, 10, "1", "10"},
		{"step2", "", "", 1, 0, 1, "1", "1"},
		{"step3", "", "", 0, 1, 1, "10", "10"},
		{"step4", "", "", 4, 2, 2, "3", "4"},
	}

	t.Run("case1: before nil/ after nil", func(t *testing.T) {
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

			edges, err := EdgesToReturn(allEdges, before, after, first, last)
			if err != nil {
				t.Fatalf("error: %v\n", err)
			}

			// domain.Nodeからdomain.Competitionへの型キャスト
			items := make([]domain.Competition, 0, len(edges))
			for _, v := range edges {
				f := v.(domain.Competition)
				items = append(items, f)
			}
			//fmt.Printf("%v\n", edges)

			assert.Equal(t, tt.out, len(items), tt.pt+":01")
			assert.Equal(t, tt.outCursorOfFirst, string(items[0].ID), tt.pt+":02")
			assert.Equal(t, tt.outCursorOfEnd, string(items[len(items)-1].ID), tt.pt+":03")
		}
	})

	// before 存在する値を指定, after 指定無し、 first 指定無し、 last 指定無し
	// before 存在する値を指定, after 指定無し、 first 存在する値を指定、 last 指定無し
	// before 存在する値を指定, after 指定無し、 first 指定無し、 last 存在する値を指定
	// before 存在する値を指定, after 指定無し、 first 存在する値を指定、 last 存在する値を指定
	var testCase2 = []struct {
		pt               string
		before           string
		after            string
		first            int
		last             int
		out              int    // 戻り値のlen値
		outCursorOfFirst string // 戻りの値の先頭の値のCursor
		outCursorOfEnd   string // 戻りの値の最後の値のCursor
	}{
		{"step1", "4", "", 0, 0, 3, "1", "3"},
		{"step2", "4", "", 1, 0, 1, "1", "1"},
		{"step3", "4", "", 0, 1, 1, "3", "3"},
		{"step4", "4", "", 4, 2, 2, "2", "3"},
	}

	t.Run("case2: before any/ after nil", func(t *testing.T) {
		for _, tt := range testCase2 {
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

			edges, err := EdgesToReturn(allEdges, before, after, first, last)
			if err != nil {
				t.Fatalf("error: %v\n", err)
			}

			// domain.Nodeからdomain.Competitionへの型キャスト
			items := make([]domain.Competition, 0, len(edges))
			for _, v := range edges {
				items = append(items, v.(domain.Competition))
			}

			assert.Equal(t, tt.out, len(items), tt.pt+":01")
			assert.Equal(t, tt.outCursorOfFirst, string(items[0].ID), tt.pt+":02")
			assert.Equal(t, tt.outCursorOfEnd, string(items[len(items)-1].ID), tt.pt+":03")
		}
	})

	// before 指定無し, after 存在する値を指定、 first 指定無し、 last 指定無し
	// before 指定無し, after 存在する値を指定、 first 存在する値を指定、 last 指定無し
	// before 指定無し, after 存在する値を指定、 first 指定無し、 last 存在する値を指定
	// before 指定無し, after 存在する値を指定、 first 存在する値を指定、 last 存在する値を指定
	var testCase3 = []struct {
		pt               string
		before           string
		after            string
		first            int
		last             int
		out              int    // 戻り値のlen値
		outCursorOfFirst string // 戻りの値の先頭の値のCursor
		outCursorOfEnd   string // 戻りの値の最後の値のCursor
	}{
		{"step1", "", "4", 0, 0, 6, "5", "10"},
		{"step2", "", "4", 1, 0, 1, "5", "5"},
		{"step3", "", "4", 0, 1, 1, "10", "10"},
		{"step4", "", "4", 4, 2, 2, "7", "8"},
	}

	t.Run("case3: before nil/ after any", func(t *testing.T) {
		for _, tt := range testCase3 {
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

			edges, err := EdgesToReturn(allEdges, before, after, first, last)
			if err != nil {
				t.Fatalf("error: %v\n", err)
			}

			// domain.Nodeからdomain.Competitionへの型キャスト
			items := make([]domain.Competition, 0, len(edges))
			for _, v := range edges {
				items = append(items, v.(domain.Competition))
			}

			assert.Equal(t, tt.out, len(items), tt.pt+":01")
			assert.Equal(t, tt.outCursorOfFirst, string(items[0].ID), tt.pt+":02")
			assert.Equal(t, tt.outCursorOfEnd, string(items[len(items)-1].ID), tt.pt+":03")
		}
	})

	// before 存在する値を指定, after 存在する値を指定、 first 指定無し、 last 指定無し
	// before 存在する値を指定, after 存在する値を指定、 first 存在する値を指定、 last 指定無し
	// before 存在する値を指定, after 存在する値を指定、 first 指定無し、 last 存在する値を指定
	// before 存在する値を指定, after 存在する値を指定、 first 存在する値を指定、 last 存在する値を指定
	var testCase4 = []struct {
		pt               string
		before           string
		after            string
		first            int
		last             int
		out              int    // 戻り値のlen値
		outCursorOfFirst string // 戻りの値の先頭の値のCursor
		outCursorOfEnd   string // 戻りの値の最後の値のCursor
	}{
		{"step1", "6", "2", 0, 0, 3, "3", "5"},
		{"step2", "6", "2", 1, 0, 1, "3", "3"},
		{"step3", "6", "2", 0, 1, 1, "5", "5"},
		{"step4", "6", "2", 4, 2, 2, "4", "5"},
	}

	t.Run("case4: before any/ after any", func(t *testing.T) {
		for _, tt := range testCase4 {
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

			edges, err := EdgesToReturn(allEdges, before, after, first, last)
			if err != nil {
				t.Fatalf("error: %v\n", err)
			}

			// domain.Nodeからdomain.Competitionへの型キャスト
			items := make([]domain.Competition, 0, len(edges))
			for _, v := range edges {
				items = append(items, v.(domain.Competition))
			}

			assert.Equal(t, tt.out, len(items), tt.pt+":01")
			assert.Equal(t, tt.outCursorOfFirst, string(items[0].ID), tt.pt+":02")
			assert.Equal(t, tt.outCursorOfEnd, string(items[len(items)-1].ID), tt.pt+":03")
		}
	})

}

func createInstance() []domain.Competition {
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

	return allEdges
}
