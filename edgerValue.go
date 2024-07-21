package dag

type edgerValue interface {
	GetEdgeValue() map[interface{}]struct{}
	GetVertexEdgeValue(interface{}) (struct{}, bool)
	SetEdgeValue(interface{}, struct{})
	Copy() edgerValue
	SetOptions(Options)
}

type edgeValue struct {
	edgeValue map[interface{}]struct{}
	options   Options
}

var _ edgerValue = &edgeValue{}

func newEmptyEdgeValue(options Options) *edgeValue {
	return &edgeValue{
		edgeValue: make(map[interface{}]struct{}),
		options:   options,
	}
}

func (e *edgeValue) GetEdgeValue() map[interface{}]struct{} {
	return e.edgeValue
}

func (e *edgeValue) GetVertexEdgeValue(vertex interface{}) (struct{}, bool) {
	vertexHash := e.options.VertexHashFunc(vertex)
	edge, exists := e.edgeValue[vertexHash]
	return edge, exists
}

func (e *edgeValue) SetEdgeValue(vertex interface{}, value struct{}) {
	vertexHash := e.options.VertexHashFunc(vertex)
	e.edgeValue[vertexHash] = value
}

func (e *edgeValue) Copy() edgerValue {
	return &edgeValue{
		edgeValue: copyMap(e.edgeValue),
		options:   e.options,
	}
}

func (e *edgeValue) SetOptions(options Options) {
	e.options = options
}
