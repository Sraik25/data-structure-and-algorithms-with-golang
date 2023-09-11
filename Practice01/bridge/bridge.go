package bridge

import "fmt"

// IDrawShape interface
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

// DrawShape struct
type DrawShape struct {
}

func (d DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

func main() {
	x := [5]float32{1, 2, 3, 4, 5}
	y := [5]float32{1, 2, 3, 4, 5}

	contour := DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
