package main

import "fmt"

var graph = make(map[string]map[string]bool)

func AddEdge(from string, to ...string) {

	for _, to := range to {
		addEdgeInternal(from, to)
		// conexão inversa
		addEdgeInternal(to, from)
	}

}

func addEdgeInternal(from, to string) {
	if graph[from] == nil {
		graph[from] = make(map[string]bool)
	}
	graph[from][to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	AddEdge("voce", "alice", "bob", "claire")
	AddEdge("bob", "anuj", "peggy")
	fmt.Println(graph)
}
