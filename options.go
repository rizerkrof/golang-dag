package dag

// Options is the configuration for the DAG.
type Options struct {
	// VertexHashFunc is the function that calculates the hash value of a vertex.
	// This can be useful when the vertex contains not comparable types such as maps.
	// If VertexHashFunc is nil, the defaultVertexHashFunc is used.
	VertexHashFunc func(v interface{}) interface{}
}

// Options sets the options for the DAG.
// Options must be called before any other method of the DAG is called.
func (d *DAG) Options(options Options) {
	d.muDAG.Lock()
	defer d.muDAG.Unlock()
	d.options = options
	d.vertices.SetOptions(options)
	d.inboundEdge.SetOptions(options)
	d.outboundEdge.SetOptions(options)
	d.ancestorsCache.SetOptions(options)
	d.descendantsCache.SetOptions(options)
	d.verticesLocked.SetOptions(options)
}

func defaultOptions() Options {
	return Options{
		VertexHashFunc: defaultVertexHashFunc,
	}
}

func defaultVertexHashFunc(v interface{}) interface{} {
	return v
}
