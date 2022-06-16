package controller

import (
	"fmt"

	"github.com/vmihailenco/msgpack/v5"

	"ledctl3/internal/client/interfaces"
	"ledctl3/internal/pkg/events"
)

type Controller struct {
	leds       int
	mode       Mode
	visualizer interfaces.Visualizer
	events     chan []byte

	displayVisualizer interfaces.Visualizer
	audioVisualizer   interfaces.Visualizer
	segmentCount      int
}

func New(opts ...Option) (*Controller, error) {
	s := &Controller{}

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}

	s.events = make(chan []byte, s.segmentCount)

	return s, nil
}

func (ctl *Controller) Start() error {
	return nil
}

func (ctl *Controller) Events() chan []byte {
	return ctl.events
}

func (ctl *Controller) Stop() error {
	if ctl.visualizer != nil {
		return ctl.visualizer.Stop()
	}

	return nil
}

type Mode string

const (
	Reset   Mode = "reset"
	Reload  Mode = "reload"
	Video   Mode = "video"
	Audio   Mode = "audio"
	Rainbow Mode = "rainbow"
	Static  Mode = "static"
)

var Modes = map[string]Mode{
	"reset":   Reset,
	"reload":  Reload,
	"video":   Video,
	"audio":   Audio,
	"rainbow": Rainbow,
	"static":  Static,
}

func (ctl *Controller) SetMode(mode Mode) error {
	if mode == ctl.mode {
		return fmt.Errorf("invalid mode")
	}

	ctl.mode = mode

	if ctl.visualizer != nil {
		err := ctl.visualizer.Stop()
		if err != nil {
			return err
		}
	}

	switch mode {
	case Video:
		ctl.visualizer = ctl.displayVisualizer
	case Audio:
		ctl.visualizer = ctl.audioVisualizer
	case Rainbow, Static:
		ctl.visualizer = nil
	}

	if ctl.visualizer != nil {
		err := ctl.visualizer.Start()
		if err != nil {
			return err
		}

		go func() {
			for evt := range ctl.visualizer.Events() {
				segs := []events.Segment{}

				for _, seg := range evt.Segments {
					segs = append(segs, events.Segment{
						Id:  seg.Id,
						Pix: seg.Pix,
					})
				}
				//fmt.Println(evt.Data)
				e := events.NewAmbilightEvent(segs)

				b, err := msgpack.Marshal(e)
				if err != nil {
					panic(err)
				}

				ctl.events <- b
			}
		}()
	}

	return nil
}
