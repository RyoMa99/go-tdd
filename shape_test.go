package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %2.f, want %2.f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}
