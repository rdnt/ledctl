package client

import (
	"context"
	"fmt"
	"image/color"
	"sync"
	"time"

	"ledctl3/internal/client/controller"
	"ledctl3/internal/client/controller/audio"
	"ledctl3/internal/client/controller/video"

	"github.com/gorilla/websocket"
)

type Application struct {
	DefaultMode controller.Mode

	Host       string
	Port       int
	Leds       int
	StripType  StripType
	GpioPin    int
	Brightness int
	BlackPoint int
	Segments   []Segment

	connMux sync.Mutex
	conn    *websocket.Conn

	Displays       video.DisplayRepository
	DisplayConfigs [][]video.DisplayConfig

	Colors     []color.Color
	WindowSize int

	//cfg config.Config
	//ip       string
	//leds     int
	//mode     string
	//capturer string

	ctl           *controller.Controller
	ServerAddress string

	//displayVisualizer visualizer.Visualizer
	//audioVisualizer   visualizer.Visualizer
}

type Segment struct {
	Id   int
	Leds int
}

func New(opts ...Option) (*Application, error) {
	a := &Application{}

	for _, opt := range opts {
		err := opt(a)
		if err != nil {
			return nil, err
		}
	}

	var err error
	addr := fmt.Sprintf("ws://%s:%d/ws", a.Host, a.Port)

	a.ServerAddress = addr

	displayVisualizer, err := video.New(
		video.WithLedsCount(a.Leds),
		video.WithDisplayRepository(a.Displays),
		video.WithDisplayConfig(a.DisplayConfigs),
	)
	if err != nil {
		return nil, err
	}

	segs := []audio.Segment{}
	for _, seg := range a.Segments {
		segs = append(segs, audio.Segment{
			Id:   seg.Id,
			Leds: seg.Leds,
		})
	}

	audioVisualizer, err := audio.New(
		audio.WithLedsCount(a.Leds),
		audio.WithSegments(segs),
		audio.WithColors(a.Colors...),
		audio.WithWindowSize(a.WindowSize),
	)

	a.ctl, err = controller.New(
		controller.WithLedsCount(a.Leds),
		controller.WithDisplayVisualizer(displayVisualizer),
		controller.WithAudioVisualizer(audioVisualizer),
		controller.WithSegmentsCount(len(segs)),
	)
	if err != nil {
		return nil, err
	}

	go func() {
		for b := range a.ctl.Events() {
			a.connMux.Lock()
			conn := a.conn
			a.connMux.Unlock()

			if conn == nil {
				continue
			}

			err := conn.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				a.connMux.Lock()
				a.conn = nil
				a.connMux.Unlock()
			}
		}
	}()

	// go func() {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Printf("\r%s", a.ctl.Statistics().AverageProcessingTime)
	// 	}
	// }()

	return a, nil
}

func (a *Application) Start() error {
	var err error

	go func() {
		for {
			// try to re-establish connection if lost
			time.Sleep(1 * time.Second)

			a.connMux.Lock()
			conn := a.conn
			a.connMux.Unlock()

			if conn == nil {
				func() {
					ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
					defer cancel()

					conn, _, err = websocket.DefaultDialer.DialContext(ctx, a.ServerAddress, nil)
					if err != nil {
						fmt.Println(err)
						return
					}

					a.connMux.Lock()
					a.conn = conn
					a.connMux.Unlock()

					//err = a.reload()
					//if err != nil {
					//	fmt.Println(err)
					//	return
					//}
				}()

			}
		}
	}()

	err = a.ctl.SetMode(a.DefaultMode)
	if err != nil {
		panic(err)
	}

	return nil
}

//func (a *Application) reload() error {
//	segments := []event.SegmentConfig{}
//
//	for _, s := range a.Segments {
//		segments = append(
//			segments, event.SegmentConfig{
//				Id:   s.Id,
//				Leds: s.Leds,
//			},
//		)
//	}
//
//	e := event.NewUpdateEvent(a.Leds, string(a.StripType), a.GpioPin, a.Brightness, segments)
//
//	b, err := json.Marshal(e)
//	if err != nil {
//		return err
//	}
//
//	a.connMux.Lock()
//	conn := a.conn
//	a.connMux.Unlock()
//
//	if conn != nil {
//		err = a.conn.WriteMessage(websocket.TextMessage, b)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

func (a *Application) Stop() error {
	a.connMux.Lock()
	conn := a.conn
	a.connMux.Unlock()

	if conn != nil {
		conn.Close()
	}

	return a.ctl.SetMode(controller.Reset)
}