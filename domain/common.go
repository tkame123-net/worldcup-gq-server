package domain

type GlobalID struct {
	VerStr     string
	EntityName string
	ID         string
}

type Node interface {
	IsNode()
}
