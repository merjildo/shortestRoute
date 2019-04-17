package search

import (
	"math"
	"sort"
)

// Path is the struct used to abtract a vertex
type Path struct {
	distance int
	nodeID   string
}

// Dijkstra is the proper function used to
// search shortest path based on graphs
func Dijkstra(graph *Graph, start *Vertex) {
	start.setDistance(0)
	unvisitedQueue := []Path{}

	for _, v := range graph.GetVertices() {
		var vertex = graph.GetVertex(v)
		unvisitedQueue = append(unvisitedQueue, Path{vertex.GetDistance(), vertex.GetID()})
	}

	sort.Slice(unvisitedQueue, func(i, j int) bool {
		return unvisitedQueue[i].distance < unvisitedQueue[j].distance
	})

	for len(unvisitedQueue) > 0 {
		var currentPath Path
		currentPath = unvisitedQueue[0]
		unvisitedQueue = unvisitedQueue[len(unvisitedQueue)-1:]
		current := graph.vertMap[currentPath.nodeID]
		current.setVisited()
		for nextKey := range current.adjacent {
			next := graph.vertMap[nextKey]
			if next.visited {
				continue
			}
			newDist := current.GetDistance() + current.GetWeight(nextKey)
			if newDist < next.GetDistance() {
				next.setDistance(newDist)
				next.setPrevious(current)
			}
		}

		for len(unvisitedQueue) > 0 {
			unvisitedQueue = unvisitedQueue[:len(unvisitedQueue)-1]
		}

		for _, vertexKey := range graph.GetVertices() {
			vertex := graph.GetVertex(vertexKey)

			if !vertex.visited {
				unvisitedQueue = append(unvisitedQueue, Path{vertex.GetDistance(), vertex.GetID()})
			}

			sort.Slice(unvisitedQueue, func(i, j int) bool {
				return unvisitedQueue[i].distance < unvisitedQueue[j].distance
			})

		}
	}
}

// Shortest is the function used to determine the distance
// between two specific points
func Shortest(v *Vertex, shortestPath *[]string) {
	if v.previous != nil {
		*shortestPath = append(*shortestPath, v.previous.GetID())
		Shortest(v.previous, shortestPath)
	}
	return
}

func resetDistances(g *Graph) *Graph {
	for _, v := range g.GetVertices() {
		g.GetVertex(v).setDistance(math.MaxInt32)
		g.GetVertex(v).UnsetVisited()
		g.GetVertex(v).NilPrevious()
	}
	return g
}
