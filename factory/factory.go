package factory

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (Circle) Draw() {
	fmt.Println("Drawing Circle")
}

type Square struct{}

func (Square) Draw() {
	fmt.Println("Drawing Square")
}

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	}

	return nil
}
