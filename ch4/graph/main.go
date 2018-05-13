package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("南京", "北京")
	addEdge("南京", "西宁")
	addEdge("南京", "昆明")
	addEdge("南京", "青岛")
	addEdge("南京", "深圳")
	fmt.Println(hasEdge("南京", "昆明"))
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
