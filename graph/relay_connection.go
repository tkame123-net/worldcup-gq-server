package graph

import (
	"errors"
	"tkame123-net/worldcup-gq-server/domain"
)

// todo: allEdgesをinterface{}へ
func ApplyCursorsToEdges(allEdges []domain.Competition, before *string, after *string) ([]domain.Competition, error) {
	//Initialize edges to be allEdges.
	edges := allEdges
	// 2 If after is set:
	if after != nil {
		// a Let afterEdge be the edge in edges whose cursor is equal to the after argument.
		afterEdge := ""
		afterEdgeIndex := 0
		for i := range edges {
			if string(edges[i].ID) == *after {
				afterEdge = *after
				afterEdgeIndex = i
				break
			}
		}

		// b If afterEdge exists:  Remove all elements of edges before and including afterEdge.
		if afterEdge != "" && afterEdgeIndex < len(edges) {
			edges = append(edges[afterEdgeIndex+1:])
		}
	}
	// 3 If before is set:
	if before != nil {
		// a Let beforeEdge be the edge in edges whose cursor is equal to the before argument.
		beforeEdge := ""
		beforeEdgeIndex := 0
		for i := range edges {
			if string(edges[i].ID) == *before {
				beforeEdge = *before
				beforeEdgeIndex = i
				break
			}
		}
		// b If beforeEdge exists: Remove all elements of edges after and including beforeEdge.
		if beforeEdge != "" && beforeEdgeIndex < len(edges) {
			edges = append(edges[:beforeEdgeIndex+1])
		}
	}
	//Return edges.
	return edges, nil
}

func EdgesToReturn(allEdges []*domain.Competition, before *string, after *string, first *int, last *int) ([]domain.Competition, error) {
	resItems := make([]domain.Competition, 0, len(allEdges))
	for _, item := range allEdges {
		resItems = append(resItems, *item)
	}
	edges, err := ApplyCursorsToEdges(resItems, before, after)
	if err != nil {
		return nil, err
	}

	if first != nil {
		//If first is less than 0:
		//Throw an error.
		if *first < 0 {
			return nil, errors.New("first less than 0")
		}
		//	If edges has length greater than than first:
		//Slice edges to be of length first by removing edges from the end of edges.
		if len(edges) > *first {
			edges = append(edges[0:*first])
		}
	}
	if last != nil {
		//If last is less than 0:
		//Throw an error.
		if *last < 0 {
			return nil, errors.New("last less than 0")
		}
		//	If edges has length greater than than last:
		//Slice edges to be of length last by removing edges from the start of edges.
		if len(allEdges) > *last {
			edges = append(edges[len(edges)-*last-1:])
		}
	}

	return edges, nil
}
