package dag

type vertexer interface {
	GetVertices() map[interface{}]string
	GetVertexID(interface{}) (string, bool)
	SetVertexID(interface{}, string)
	DeleteVertex(interface{})
	SetOptions(Options)
}

type vertices struct {
	vertices map[interface{}]string
	options  Options
}

func newEmptyVertices() vertexer {
	return &vertices{
		vertices: make(map[interface{}]string),
		options:  defaultOptions(),
	}
}

var _ vertexer = &vertices{}

func (v *vertices) GetVertices() map[interface{}]string {
	return v.vertices
}

func (v *vertices) GetVertexID(vertex interface{}) (string, bool) {
	vertexHash := v.options.VertexHashFunc(vertex)
	id, exists := v.vertices[vertexHash]
	return id, exists
}

func (v *vertices) IsVertex(vertex interface{}) bool {
	vertexHash := v.options.VertexHashFunc(vertex)
	_, exists := v.vertices[vertexHash]
	return exists
}

func (v *vertices) SetVertexID(vertex interface{}, id string) {
	vertexHash := v.options.VertexHashFunc(vertex)
	v.vertices[vertexHash] = id
}

func (v *vertices) DeleteVertex(vertex interface{}) {
	vertexHash := v.options.VertexHashFunc(vertex)
	delete(v.vertices, vertexHash)
}

func (v *vertices) SetOptions(options Options) {
	v.options = options
}
