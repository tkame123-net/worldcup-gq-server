package domain

type PlayerID string

type Player struct {
	Name        string
	MatchIDList []int
	MatchList   []Match
}
