package main

import (
	"sort"
)

func kruskal(edges []Edge, dim int) []Edge {

	sort.Sort(ByWeight(edges))

	// create the slice of connected components
	components := make([]int, dim*dim)
	for i := 0; i < dim*dim; i++ {
		components[i] = i
	}

	mcst := make([]Edge, 0)

	for _, edge := range edges {
		node1, node2 := edge.Node1, edge.Node2
		if !createsCycle(components, node1, node2) {
			connectComponents(components, node1, node2)
			mcst = append(mcst, edge)
		}
	}

	return mcst
}

func createsCycle(components []int, node1, node2 int) bool {
	return components[node1-1] == components[node2-1]
}

func connectComponents(components []int, node1, node2 int) {
	for i := range components {
		if components[i] == components[node2-1] {
			components[i] = components[node1-1]
		}
	}
}