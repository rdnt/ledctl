package scrap

import (
	"context"
	"errors"
	"fmt"
	goscrap "github.com/rdnt/go-scrap"
	"golang.org/x/image/draw"
	"runtime"
	"time"
)

var framerate = 240
var ErrNoFrame = fmt.Errorf("no frame")

type display struct {
	index    int
	id       int
	width    int
	height   int
	x        int
	y        int
	scaler   draw.Scaler
	capturer *goscrap.Capturer
}

func (d *display) SyncCapture(ctx context.Context, frames chan []byte) {
	panic("implement me")
}

func (d *display) Id() int {
	return d.id
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

func (d *display) Scaler() draw.Scaler {
	return d.scaler
}

func (d *display) Resolution() string {
	return fmt.Sprintf("%dx%d", d.width, d.height)
}

func (d *display) String() string {
	return fmt.Sprintf("{id: %d, width: %d, height: %d, left: %d, top: %d}", d.id, d.width, d.height, d.x, d.y)
}

func (d *display) nextFrame() ([]byte, error) {
	img, wouldBlock, err := d.capturer.FrameImage()
	if wouldBlock {
		return nil, ErrNoFrame
	}
	if err != nil {
		return nil, d.reset()
	}

	img.Detach()

	return img.Pix, nil
}

func (d *display) reset() error {
	d.capturer = nil
	runtime.GC()

	sd, err := goscrap.GetDisplay(d.index)
	if err != nil {
		return err
	}

	d.capturer, err = goscrap.NewCapturer(sd)
	if err != nil {
		return err
	}

	return nil
}

func (d *display) Capture(ctx context.Context) chan []byte {
	frames := make(chan []byte)

	go func() {
		ticker := time.NewTicker(time.Duration(1000/framerate) * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("close chan")
				close(frames)
				return
			default:
				pix, err := d.nextFrame()
				if errors.Is(err, ErrNoFrame) {
					continue
				}
				if err != nil {
					fmt.Print("\nerr in next", err)
					continue
					//close(frames)
					//return
				}

				frames <- pix
			}
		}
	}()

	return frames
}