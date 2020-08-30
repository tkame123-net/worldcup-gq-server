package domain

type PlayerID string

type Player struct {
	Name      string
	MatchList []Match
}
