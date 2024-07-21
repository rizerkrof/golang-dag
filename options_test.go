package dag

import "testing"

func TestOverrideVertexHashFunOption(t *testing.T) {
	type testVertexType struct {
		comparableField    string
		notComparableField map[string]string
	}

	dag := NewDAG()
	dag.Options(Options{
		VertexHashFunc: func(v interface{}) interface{} {
			return v.(testVertexType).comparableField
		}})

	testVertex := testVertexType{
		comparableField:    "comparable",
		notComparableField: map[string]string{"not": "comparable"},
	}
	_, err := dag.addVertex(testVertex)
	if err != nil {
		t.Errorf("Should create a vertex with a not comparable field when a correct VertexHashFunc option is set")
	}
}
