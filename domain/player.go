package domain

type PlayerID string

type Player struct {
	ID        PlayerID
	Name      string
	MatchList []Match
}

func (Player) IsNode() {}
