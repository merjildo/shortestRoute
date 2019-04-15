package search

type Graph struct {
	vertMap     map[string]*Vertex
	numVertices int
	previous    string
}

func NewGraph() *Graph {
	g := new(Graph)
	g.vertMap = make(map[string]*Vertex)
	return g
}

func (g *Graph) AddVertex(nodeID string) *Vertex {
	g.numVertices = g.numVertices + 1
	newVertex := NewVertex(nodeID)
	g.vertMap[nodeID] = newVertex
	return newVertex
}

func (g *Graph) GetVertex(nodeID string) *Vertex {
	return g.vertMap[nodeID]
}

func (g Graph) GetVertices() []string {
	var keys = []string{}
	for k := range g.vertMap {
		keys = append(keys, k)
	}
	return keys
}

func (g *Graph) AddEdge(frm string, to string, cost int) {
	if _, ok := g.vertMap[frm]; !ok {
		g.AddVertex(frm)
	}

	if _, ok := g.vertMap[to]; !ok {
		g.AddVertex(to)
	}

	g.vertMap[frm].addNeighbor(to, cost)
	g.vertMap[to].addNeighbor(frm, cost)
}

func (g *Graph) setPrevious(current string) {
	g.previous = current
}

func (g Graph) GetPrevious(current string) string {
	return g.previous
}
