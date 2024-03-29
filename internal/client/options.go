package client

import (
	"encoding/json"
	"fmt"

	"ledctl3/internal/client/config"
	"ledctl3/internal/client/controller/video/capturer/bitblt"
	"ledctl3/internal/client/controller/video/capturer/dxgi"
)

type Option func(*Application) error

//func WithLedsCount(leds int) Option {
//	return func(a *Application) error {
//		a.leds = leds
//		return nil
//	}
//}

func WithConfig(cfg config.Config) Option {
	return func(a *Application) error {
		b, _ := json.Marshal(cfg)
		fmt.Println(string(b))

		err := a.validateConfig(cfg)
		if err != nil {
			return err
		}

		err = a.applyConfig(cfg)
		if err != nil {
			return err
		}

		return nil
	}
}

func validateBounds(width, height, x, y int) bool {
	if x == 0 || x == width-1 {
		if y >= 0 && y < height {
			return true
		}
	} else if y == 0 || y == height-1 {
		if x >= 0 && x < width {
			return true
		}
	}
	return false
}

func calculateOffset(width, height, x, y int) int {
	var offset int
	if x == 0 {
		offset = 2*width + height + (height - y)
	} else if x == width-1 {
		offset = width + y
	} else {
		if y == 0 {
			offset = x
		} else if y == height-1 {
			offset = width + height + (width - x)
		} else {
			return 0
		}
	}
	// offset = offset % (d.Width*2 + d.Height*2)
	return offset
}

func getPixSliceSize(width, height, from, to int) int {
	return (width*2 + height*2) - from + to
}

//func WithServerIP(ip string) Option {
//	return func(a *Application) error {
//		a.ip = ip
//		return nil
//	}
//}

func WithDisplayCapturer(capturer CapturerType) Option {
	return func(a *Application) error {
		var err error

		switch capturer {
		case DXGI:
			a.Displays, err = dxgi.New()
			if err != nil {
				return err
			}
		case BitBlt:
			a.Displays = bitblt.New()

		default:
			return fmt.Errorf("invalid capturer")
		}

		return nil
	}
}
