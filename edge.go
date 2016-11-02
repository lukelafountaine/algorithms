package main

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
