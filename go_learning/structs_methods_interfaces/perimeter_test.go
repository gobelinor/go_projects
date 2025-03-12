package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.0, 2.0}
	got := rectangle.Perimeter()
	want := 8.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want) // %.2f is used to format the float to 2 decimal places
	}
}

func TestArea(t *testing.T) {
	checkAera := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{2.0, 2.0}
		want := 4.0
		checkAera(t, rectangle, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{2.0}
		want := 2.0 * 2.0 * math.Pi
		checkAera(t, circle, want)
	})
}

func TestAreaTableDriven(t *testing.T) {
	assertArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}
	areaTests := []struct {
		// anonymous struct areaTests slice of struct
		testname string
		shape    Shape
		hasArea  float64
	}{
		// fill the slice with struct values
		{testname: "Triangle", shape: Triangle{Base: 2.0, Height: 2.0}, hasArea: 2.0},
		{testname: "Circle", shape: Circle{Radius: 2.0}, hasArea: 2.0 * 2.0 * math.Pi},
		{testname: "Rectangle", shape: Rectangle{Width: 2.0, Height: 2.0}, hasArea: 4.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.testname, func(t *testing.T) {
			assertArea(t, tt.shape, tt.hasArea)
		})
	}
}
