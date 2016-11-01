package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

type Edge struct {
	Node1, Node2, Weight int
}

// type to sort by edge weight
type ByWeight []Edge

func (a ByWeight) Len() int {
	return len(a)
}
func (a ByWeight) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByWeight) Less(i, j int) bool {
	return a[i].Weight < a[j].Weight
}

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
	new := components[node1-1]
	old := components[node2-1]

	for i := range components {
		if components[i] == old {
			components[i] = new
		}
	}
}

func main() {

	// verify number of command line args
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	filename := os.Args[1]
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := string(contents)

	// extract the n-grid dimension and edges
	parts := strings.Split(data, ":")
	raw_dim, raw_edges := parts[0], parts[1]
	dim64, err := strconv.ParseInt(raw_dim, 10, 32)
	dim := int(dim64)

	var edges []Edge

	for _, x := range strings.Split(strings.Trim(raw_edges, "()"), "),(") {

		edge := strings.Split(x, ",")

		start, err:= strconv.Atoi(edge[0])
		end, err := strconv.Atoi(edge[1])
		weight, err := strconv.Atoi(edge[2])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		edges = append(edges, Edge{start, end, weight})
	}

	mcst := kruskal(edges, dim)
	fmt.Println(mcst)
}