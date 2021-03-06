// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
}

type Competition struct {
	ID      string   `json:"id"`
	Year    string   `json:"year"`
	Country string   `json:"country"`
	Matches []*Match `json:"matches"`
}

func (Competition) IsNode() {}

type CompetitionConnection struct {
	Edges    []*CompetitionEdge `json:"edges"`
	PageInfo *PageInfo          `json:"pageInfo"`
}

type CompetitionEdge struct {
	Cursor string       `json:"cursor"`
	Node   *Competition `json:"node"`
}

type Match struct {
	ID      string `json:"id"`
	Year    int    `json:"year"`
	Stage   string `json:"stage"`
	Stadium string `json:"stadium"`
	City    string `json:"city"`
}

func (Match) IsNode() {}

type MatchConnection struct {
	Edges    []*MatchEdge `json:"edges"`
	PageInfo *PageInfo    `json:"pageInfo"`
}

type MatchEdge struct {
	Cursor string `json:"cursor"`
	Node   *Match `json:"node"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type Player struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	MatchList []*Match `json:"matchList"`
}

func (Player) IsNode() {}

type PlayerConnection struct {
	Edges    []*PlayerEdge `json:"edges"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type PlayerEdge struct {
	Cursor string  `json:"cursor"`
	Node   *Player `json:"node"`
}
