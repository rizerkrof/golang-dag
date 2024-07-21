package dag

type edger interface {
	GetEdges() map[interface{}]edgerValue
	InitEdge(interface{}, interface{})
	InitVertexEdges(interface{})
	GetEdge(interface{}, interface{}) (struct{}, bool)
	GetVertexEdges(interface{}) (edgerValue, bool)
	SetVertexEdges(interface{}, edgerValue)
	DeleteEdges()
	DeleteVertexEdge(interface{}, interface{})
	DeleteVertexEdges(interface{})
	SetOptions(Options)
}

type edges struct {
	edges   map[interface{}]edgerValue
	options Options
}

func newEmptyEdges() edger {
	return &edges{
		edges:   make(map[interface{}]edgerValue),
		options: defaultOptions(),
	}
}

var _ edger = &edges{}

func (e *edges) GetEdges() map[interface{}]edgerValue {
	return e.edges
}

func (e *edges) InitEdge(from, to interface{}) {
	fromHash := e.options.VertexHashFunc(from)
	toHash := e.options.VertexHashFunc(to)
	e.edges[fromHash].SetEdgeValue(toHash, struct{}{})
}

func (e *edges) InitVertexEdges(vertex interface{}) {
	vertexHash := e.options.VertexHashFunc(vertex)
	e.edges[vertexHash] = newEmptyEdgeValue()
}

func (e *edges) GetEdge(from, to interface{}) (struct{}, bool) {
	fromHash := e.options.VertexHashFunc(from)
	toHash := e.options.VertexHashFunc(to)
	edge, exists := e.edges[fromHash].GetVertexEdgeValue(toHash)
	return edge, exists
}

func (e *edges) GetVertexEdges(vertex interface{}) (edgerValue, bool) {
	vertexHash := e.options.VertexHashFunc(vertex)
	edges, exists := e.edges[vertexHash]
	return edges, exists
}

func (e *edges) SetVertexEdges(vertex interface{}, edges edgerValue) {
	vertexHash := e.options.VertexHashFunc(vertex)
	e.edges[vertexHash] = edges
}

func (e *edges) DeleteEdges() {
	e.edges = make(map[interface{}]edgerValue)
}

func (e *edges) DeleteVertexEdge(from, to interface{}) {
	fromHash := e.options.VertexHashFunc(from)
	toHash := e.options.VertexHashFunc(to)
	delete(e.edges[fromHash].GetEdgeValue(), toHash)
}

func (e *edges) DeleteVertexEdges(vertex interface{}) {
	vertexHash := e.options.VertexHashFunc(vertex)
	delete(e.edges, vertexHash)
}

func (e *edges) SetOptions(options Options) {
	e.options = options
}
