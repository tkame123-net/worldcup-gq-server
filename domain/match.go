package domain

type MatchID string

type Match struct {
	ID      MatchID
	Year    int
	Stage   string
	Stadium string
	City    string
	//RoundId      int
	//MatchId      int
	//HomeTeamName string
	//AwayTeamName string
}
