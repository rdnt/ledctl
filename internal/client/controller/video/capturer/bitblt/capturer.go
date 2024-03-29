package bitblt

import (
	"github.com/kbinani/screenshot"

	"ledctl3/internal/client/controller/video"
)

var scaleFactor = 8

type bitbltCapturer struct {
	displays []*display
}

func (c *bitbltCapturer) All() ([]video.Display, error) {
	ds := []video.Display{}

	count := screenshot.NumActiveDisplays()
	for i := 0; i < count; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		ds = append(
			ds, &display{
				id:     i,
				width:  bounds.Dx(),
				height: bounds.Dy(),
				x:      bounds.Min.X,
				y:      bounds.Min.Y,
			},
		)
	}

	return ds, nil
}

func New() *bitbltCapturer {
	return &bitbltCapturer{}
}
