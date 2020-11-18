package domain

type CompetitionID string

type Competition struct {
	ID      CompetitionID
	Year    string
	Country string
}

func (Competition) IsNode() {}
