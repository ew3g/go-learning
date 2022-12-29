// Using a graph where the key type is a string and the value type is map[string]bool
package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("me", "her")
	fmt.Println(hasEdge("me", "her"))
	fmt.Println(hasEdge("her", "me"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
