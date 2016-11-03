package main

import (
	"strings"
	"strconv"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	//run_kruskals()
	sequence := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	fmt.Println(longestIncreasingSub(sequence))
}

func run_kruskals() {

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