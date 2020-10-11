package graph

import (
	"tkame123-net/worldcup-gq-server/domain"
)

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
