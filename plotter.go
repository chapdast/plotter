package plotter

import (
	"io"
	"gonum.org/v1/plot"

)

type Point interface {
	Langitude() int
	Latitude() int
	Altitude() int
}

type RoutePlotter interface {
	Route(w *io.Writer, points []Point) (int, error)
}

type point struct {
	lat int
	lng int
	alt int
}

func (p point) Langitude() int {
	return p.lng
}
func (p point) Latitude() int {
	return p.lat
}
func (p point) Altitude() int {
	return p.alt
}

func NewPoint(lat, lng, alt int) Point {
	return point{lat: lat, lng: lng, alt: alt}
}

func NewRoutePlotter() RoutePlotter {
	return routePlotter{}
}

type routePlotter struct{
	plotter *plot.Plot
}


func (r routePlotter) Route(w *io.Writer, points []Point) (int, error) {
	panic("Implement me")
}
