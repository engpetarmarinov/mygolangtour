package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Implements geometrical shapes (Triangle, Square, Rectangle) all of them have Areas :)

// Triangle area -> 0.5 * triangle.base * triangle.height
// Square area -> square.length * square.length
// Rectangle area -> rect.length * rect.breadth

// Implement function CalculateArea() that takes any geometrical shape and sum the Area for all of them

// Make the CalculateArea() concurrent

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func NewTriangle(base, height float64) (*Triangle, error) {
	if base <= 0 {
		return nil, fmt.Errorf("%f base must be a positive number", base)
	}

	if height <= 0 {
		return nil, fmt.Errorf("%f height must be a positive number", height)
	}

	return &Triangle{
		Base:   base,
		Height: height,
	}, nil
}

func (t *Triangle) Area() float64 {
	//Triangle area -> 0.5 * triangle.base * triangle.height
	time.Sleep(20 * time.Millisecond)
	return 0.5 * t.Base * t.Height
}

type Square struct {
	Length float64
}

func NewSquare(length float64) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("%f length must be a positive number", length)
	}

	return &Square{
		Length: length,
	}, nil
}

func (s *Square) Area() float64 {
	time.Sleep(10 * time.Millisecond)
	return s.Length * s.Length
}

func CalculateArea(shapes ...Shape) float64 {
	var sumOfAreas float64
	areasChans := make([]<-chan interface{}, len(shapes))
	ctx := context.Background()
	defer ctx.Done()

	for _, shape := range shapes {
		areaChan := getAreaChan(ctx, shape)
		areasChans = append(areasChans, areaChan)
	}

	for areaAsInterface := range Take(ctx, FanIn(ctx, areasChans...), len(shapes)) {
		if area, ok := areaAsInterface.(float64); ok {
			sumOfAreas += area
		}
	}

	return sumOfAreas
}

func CalculateAreaWithWG(shapes ...Shape) float64 {
	var sumOfAreas float64
	wg := &sync.WaitGroup{}
	areaChan := make(chan float64, len(shapes))
	defer close(areaChan)

	wg.Add(len(shapes))
	for _, shape := range shapes {
		//shadowing of shape
		shape := shape
		go func() {
			defer wg.Done()
			areaChan <- shape.Area()
		}()
	}

	wg.Wait()

	for range shapes {
		sumOfAreas += <-areaChan
	}

	return sumOfAreas
}

func CalculateAreaWithoutChan(shapes ...Shape) float64 {
	var wg sync.WaitGroup
	areas := make([]float64, len(shapes))
	for i, shape := range shapes {
		wg.Add(1)
		go func(i int, shape Shape) {
			defer wg.Done()
			areas[i] = shape.Area()
		}(i, shape)
	}
	wg.Wait()

	var sum float64
	for _, area := range areas {
		sum += area
	}

	return sum
}

func getAreaChan(ctx context.Context, shape Shape) <-chan interface{} {
	areaChan := make(chan interface{})
	go func(shape Shape, areaChan chan interface{}) {
		defer close(areaChan)
		for {
			select {
			case <-ctx.Done():
				return
			case areaChan <- shape.Area():
				return
			}
		}
	}(shape, areaChan)

	return areaChan
}
