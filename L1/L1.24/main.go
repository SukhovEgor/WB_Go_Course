/* Разработать программу нахождения расстояния между двумя точками
на плоскости. Точки представлены в виде структуры Point с
инкапсулированными (приватными) полями x, y (типа float64)
и конструктором. Расстояние рассчитывается по формуле между координатами двух точек.

Подсказка: используйте функцию-конструктор NewPoint(x, y),
Point и метод Distance(other Point) float64. */

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	p1 := NewPoint(RandPoint(), RandPoint())
	p2 := NewPoint(RandPoint(), RandPoint())
	fmt.Println("Point 1 is ", p1)
	fmt.Println("Point 2 is ", p2)
	fmt.Println("Distance between point 1 and point 2 is", p1.Distance(*p2))

}

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (point1 *Point) Distance(point2 Point) float64 {
	distance := math.Sqrt(math.Pow((point2.x-point1.x), 2) + math.Pow((point2.y-point1.y), 2))
	return distance
}

func RandPoint() float64 {
	min := -10.00
	max := 10.00
	return min + rand.Float64()*(max-min)
}
