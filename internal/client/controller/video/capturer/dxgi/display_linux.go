package dxgi

import (
	"context"
	"fmt"
	"image"

	"ledctl3/internal/client/controller/video"
)

type display struct {
	index  int
	id     int
	width  int
	height int
	x      int
	y      int
	buf    *image.RGBA
}

func (d *display) Id() int {
	return 0
}

func (d *display) Width() int {
	return d.width
}

func (d *display) Height() int {
	return d.height
}

func (d *display) X() int {
	return d.x
}

func (d *display) Y() int {
	return d.y
}

func (d *display) Resolution() string {
	return fmt.Sprintf("%dx%d", d.width, d.height)
}

func (d *display) String() string {
	return fmt.Sprintf("{id: %d, width: %d, height: %d, left: %d, top: %d}", d.id, d.width, d.height, d.x, d.y)
}

func (d *display) Capture(ctx context.Context, framerate int) chan []byte {
	return nil
}

func (d *display) Close() error {
	return nil
}

func (d *display) reset() error {
	return nil
}

func (d *display) Orientation() video.Orientation {
	return video.Landscape
}
