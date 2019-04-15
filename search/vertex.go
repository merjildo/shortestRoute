package search

import "math"

type Vertex struct {
	id       string
	adjacent map[string]int
	distance int
	visited  bool
	previous *Vertex
}

func NewVertex(nodeID string) *Vertex {
	newVertex := new(Vertex)
	newVertex.id = nodeID
	newVertex.adjacent = make(map[string]int)
	newVertex.visited = false
	newVertex.distance = math.MaxInt32
	newVertex.previous = nil
	return newVertex
}

func (v *Vertex) addNeighbor(neighbor string, weight int) {
	v.adjacent[neighbor] = weight
}

func (v Vertex) GetConnections() []string {
	var keys = []string{}
	for k := range v.adjacent {
		keys = append(keys, k)
	}
	return keys
}

func (v Vertex) GetID() string {
	return v.id
}

func (v Vertex) GetWeight(neighbor string) int {
	return v.adjacent[neighbor]
}

func (v *Vertex) setDistance(dist int) {
	v.distance = dist
}

func (v *Vertex) GetDistance() int {
	return v.distance
}

func (v *Vertex) setPrevious(prev *Vertex) {
	v.previous = prev
}

func (v *Vertex) setVisited() {
	v.visited = true
}

func (v *Vertex) UnsetVisited() {
	v.visited = false
}

func (v *Vertex) NilPrevious() {
	v.previous = nil
}
