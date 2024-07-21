package dag

import "testing"

func TestOverrideVertexHashFunOption(t *testing.T) {
	type testVertexType struct {
		idField            string
		notComparableField map[string]string
	}

	dag := NewDAG()
	dag.Options(Options{
		VertexHashFunc: func(v interface{}) interface{} {
			return v.(testVertexType).idField
		}})

	testVertex1 := testVertexType{
		idField:            "comparable",
		notComparableField: map[string]string{"not": "comparable"},
	}
	vertexId1, err := dag.addVertex(testVertex1)
	if err != nil {
		t.Errorf("Should create a vertex with a not comparable field when a correct VertexHashFunc option is set")
	}

	testVertex2 := testVertexType{
		idField:            "stillComparable",
		notComparableField: map[string]string{"stillNot": "comparable"},
	}
	vertexId2, err := dag.addVertex(testVertex2)
	if err != nil {
		t.Errorf("Should create a vertex with a not comparable field when a correct VertexHashFunc option is set")
	}

	err = dag.AddEdge(vertexId1, vertexId2)
	if err != nil {
		t.Errorf("Should create an edge between vertices with not comparable fields when a correct VertexHashFunc option is set")
	}
}
