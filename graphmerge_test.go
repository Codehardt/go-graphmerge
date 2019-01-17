package graphmerge

import (
	"reflect"
	"testing"
)

func testGraphMerge(t *testing.T, eX, eY []float64, c ...[]float64) {
	x, y, err := GraphMerge(c...)
	if err != nil {
		t.Errorf("could not graph merge '%#v': %s", c, err)
		return
	}
	if !reflect.DeepEqual(&x, &eX) {
		t.Errorf("x coords not deep equal, got '%#v', want: '%#v'", x, eX)
		return
	}
	if !reflect.DeepEqual(&y, &eY) {
		t.Errorf("y coords not deep equal, got '%#v', want: '%#v'", y, eY)
		return
	}
}

func TestGraphMerge(t *testing.T) {
	testGraphMerge(t,
		[]float64{}, []float64{}, // expected
	)
	testGraphMerge(t,
		[]float64{}, []float64{}, // expected
		nil, nil, // graph 1
		nil, nil, // graph 2
	)
	testGraphMerge(t,
		[]float64{}, []float64{}, // expected
		[]float64{}, []float64{}, // graph 1
		[]float64{}, []float64{}, // graph 2
	)
	testGraphMerge(t,
		[]float64{0}, []float64{2}, // expected
		[]float64{0}, []float64{1}, // graph 1
		[]float64{0}, []float64{1}, // graph 2
	)
	testGraphMerge(t,
		[]float64{0, 1}, []float64{1, 2}, // expected
		[]float64{0}, []float64{1}, // graph 1
		[]float64{1}, []float64{1}, // graph 2
	)
	testGraphMerge(t,
		[]float64{}, []float64{}, // expected
		[]float64{}, []float64{}, // graph 1
		[]float64{}, []float64{}, // graph 2
	)
	testGraphMerge(t,
		[]float64{0, 1, 2, 3}, []float64{10, 11, 12, 13}, // expected
		[]float64{0, 1, 2, 3}, []float64{1, 2, 3, 4}, // graph 1
		[]float64{0, 1, 2, 3}, []float64{4, 3, 2, 1}, // graph 2
		[]float64{0, 1, 2, 3}, []float64{5, 6, 7, 8}, // graph 3
	)
	testGraphMerge(t,
		[]float64{0, 1, 2, 3, 4, 5, 6, 7}, []float64{1, 2, 3, 4, 9, 10, 11, 12}, // expected
		[]float64{0, 1, 2, 3}, []float64{1, 2, 3, 4}, // graph 1
		[]float64{4, 5, 6, 7}, []float64{5, 6, 7, 8}, // graph 2
	)
	testGraphMerge(t,
		[]float64{0, 1, 2, 3, 4, 5, 6}, []float64{2, 3, 4, 5, 6, 7, 8}, // expected
		[]float64{0, 1, 3, 5}, []float64{1, 2, 3, 4}, // graph 1
		[]float64{0, 2, 4, 6}, []float64{1, 2, 3, 4}, // graph 2
	)
	testGraphMerge(t,
		[]float64{0, 1, 2, 3, 4, 5, 6, 7}, []float64{3, 4, 5, 7, 8, 10, 11, 12}, // expected
		[]float64{0, 1, 3, 5}, []float64{1, 2, 3, 4}, // graph 1
		[]float64{0, 2, 4, 6}, []float64{1, 2, 3, 4}, // graph 2
		[]float64{0, 3, 5, 7}, []float64{1, 2, 3, 4}, // graph 3
	)
}
