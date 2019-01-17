package graphmerge

import (
	"fmt"
	"sort"
)

// GraphMerge merges multiple graphes into one big
// graph. Usage:
// x, y, err := GraphMerge(x1, y1, x2, y2, ..., xN, yN)
func GraphMerge(c ...[]float64) ([]float64, []float64, error) {
	if len(c)%2 != 0 {
		return nil, nil, fmt.Errorf("expected even length of graph coordinate slices, got: %d", len(c))
	}
	coords := map[float64]float64{}
	for i := 0; i < len(c)-1; i += 2 {
		xCoords := c[i]
		yCoords := c[i+1]
		if !sort.Float64sAreSorted(xCoords) {
			return nil, nil, fmt.Errorf("graph coordinate slice %d is not sorted", i)
		}
		if len(xCoords) != len(yCoords) {
			return nil, nil, fmt.Errorf("length of x and y coordinates do not match in graph coordinate slices %d and %d, got: %d and %d", i, i+1, len(xCoords), len(yCoords))
		}
		var sumYCoords float64
		for j, xCoord := range xCoords {
			yCoord := yCoords[j] - sumYCoords
			sumYCoords += yCoord
			coords[xCoord] += yCoord
		}
	}
	xCoords := make([]float64, len(coords))
	yCoords := make([]float64, len(coords))
	var counter int
	for xCoord := range coords {
		xCoords[counter] = xCoord
		counter++
	}
	sort.Float64s(xCoords)
	var yCoord float64
	for i, xCoord := range xCoords {
		yCoord = coords[xCoord] + yCoord
		yCoords[i] = yCoord
	}
	return xCoords, yCoords, nil
}
