package main

type Node struct {
	InDegree  int
	Letter string
	Neighbors []Node
}

func TopologicalSort(graph []Node) []Node {

	queue := make(chan Node)
	order := make([]Node, 0)

	for _, node := range graph {
		if node.InDegree == 0 {
			queue <- node
		}
	}

	for node := range queue {

		order = append(order, node)

		for _, neighbor := range node.Neighbors {
			neighbor.InDegree -= 1

			if neighbor.InDegree == 0 {
				queue <- neighbor
			}
		}
	}

	return order
}




