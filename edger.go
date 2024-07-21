package dag

type edger interface {
	GetEdges() map[interface{}]map[interface{}]struct{}
	InitEdge(interface{}, interface{})
	InitVertexEdges(interface{})
	GetEdge(interface{}, interface{}) (struct{}, bool)
	GetVertexEdges(interface{}) (map[interface{}]struct{}, bool)
	SetVertexEdges(interface{}, map[interface{}]struct{})
	DeleteEdges()
	DeleteVertexEdge(interface{}, interface{})
	DeleteVertexEdges(interface{})
	SetOptions(Options)
}

type edges struct {
	edges   map[interface{}]map[interface{}]struct{}
	options Options
}

func newEmptyEdges() edger {
	return &edges{
		edges:   make(map[interface{}]map[interface{}]struct{}),
		options: defaultOptions(),
	}
}

var _ edger = &edges{}

func (e *edges) GetEdges() map[interface{}]map[interface{}]struct{} {
	return e.edges
}

func (e *edges) InitEdge(from, to interface{}) {
	fromHash := e.options.VertexHashFunc(from)
	toHash := e.options.VertexHashFunc(to)
	e.edges[fromHash][toHash] = struct{}{}
}

func (e *edges) InitVertexEdges(vertex interface{}) {
	vertexHash := e.options.VertexHashFunc(vertex)
	e.edges[vertexHash] = make(map[interface{}]struct{})
}

func (e *edges) GetEdge(from, to interface{}) (struct{}, bool) {
	fromHash := e.options.VertexHashFunc(from)
	toHash := e.options.VertexHashFunc(to)
	edge, exists := e.edges[fromHash][toHash]
	return edge, exists
}

func (e *edges) GetVertexEdges(vertex interface{}) (map[interface{}]struct{}, bool) {
	vertexHash := e.options.VertexHashFunc(vertex)
	edges, exists := e.edges[vertexHash]
	return edges, exists
}

func (e *edges) SetVertexEdges(vertex interface{}, edges map[interface{}]struct{}) {
	vertexHash := e.options.VertexHashFunc(vertex)
	e.edges[vertexHash] = edges
}

func (e *edges) DeleteEdges() {
	e.edges = make(map[interface{}]map[interface{}]struct{})
}

func (e *edges) DeleteVertexEdge(from, to interface{}) {
	fromHash := e.options.VertexHashFunc(from)
	toHash := e.options.VertexHashFunc(to)
	delete(e.edges[fromHash], toHash)
}

func (e *edges) DeleteVertexEdges(vertex interface{}) {
	vertexHash := e.options.VertexHashFunc(vertex)
	delete(e.edges, vertexHash)
}

func (e *edges) SetOptions(options Options) {
	e.options = options
}
