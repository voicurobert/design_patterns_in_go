package adapter

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}

// the interface you are given

// the interface we have

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// solution

type vectorToRasterAdapter struct {
	points []Point
}

var pointCache = map[[16]byte][]Point{}

func (v *vectorToRasterAdapter) addLineCached(line Line) {

	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}

	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		for _, pt := range pts {
			v.points = append(v.points, pt)
		}
		return
	}
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)

	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}

	pointCache[h] = v.points

	fmt.Println("we have", len(v.points), "points")
}

func (v *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)

	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}

	fmt.Println("we have", len(v.points), "points")
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(vi VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLineCached(line)
	}

	return adapter // as RasterImage
}

func AdapterExample() {
	rc := NewRectangle(6, 4)

	adapter := VectorToRaster(*rc)

	fmt.Print(DrawPoints(adapter))
}
