# go-graphmerge

Merge multiple graphs with x- and y-coordinates into one graph 

## Example

```golang
// Graph 1
...
x1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
y1 := []float64{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
// Graph 2
x2 := []float64{0, 2, 4, 6, 8, 10}
y2 := []float64{0, 1, 2, 3, 4, 5}
// Graph 3
x3 := []float64{11, 15, 18, 22}
y3 := []float64{-1, -2, -3, -4}
// Merged Graph
x, y, err := graphmerge.GraphMerge(x1, y1, x2, y2, x3, y3)
if err != nil {
    ...
}
...
```

## Result

![example.png](https://github.com/Codehardt/go-graphmerge/raw/master/example/example.png)