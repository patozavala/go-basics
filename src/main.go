package main

import "fmt"

func rectangleArea(width float64, length float64) float64 {
	area := width * length
	return area
}

func rectanglePerimeter(width float64, length float64) float64 {
	perimeter := 2 * (width + length)
	return perimeter
}

func rectangleData(width float64, length float64) (float64, float64) {
	area := rectangleArea(width, length)
	perimeter := rectanglePerimeter(width, length)
	return area, perimeter

}

func main() {
	const width float64 = 5
	const length float64 = 4
	fmt.Println(rectangleData(width, length))
}
