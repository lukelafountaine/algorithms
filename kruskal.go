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
	Node1, Node2, Weight int64
}

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

func kruskall(edges []Edge, dim int) []Edge {
	sort.Sort(ByWeight(edges))
	fmt.Println(edges)

	return nil
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

		start, err:= strconv.ParseInt(edge[0], 10, 64)
		end, err := strconv.ParseInt(edge[1], 10, 64)
		weight, err := strconv.ParseInt(edge[2], 10, 64)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		edges = append(edges, Edge{start, end, weight})
	}

	fmt.Println(dim, edges)
	mcst := kruskall(edges, dim)
	fmt.Println(mcst, edges)

}