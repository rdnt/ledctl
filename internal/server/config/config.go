package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	StripType   string        `yaml:"stripType" json:"stripType"`
	GpioPin     int           `yaml:"gpioPin" json:"gpioPin"`
	Brightness  int           `yaml:"brightness" json:"brightness"`
	Segments    []Segment     `yaml:"segments" json:"segments"`
	Calibration []Calibration `yaml:"calibration" json:"calibration"`
}

type Segment struct {
	Id   int `yaml:"id" json:"id"`
	Leds int `yaml:"leds" json:"leds"`
}

type Calibration struct {
	Start int     `yaml:"start" json:"start"`
	End   int     `yaml:"end" json:"end"`
	Red   float64 `yaml:"red" json:"red"`
	Green float64 `yaml:"green" json:"green"`
	Blue  float64 `yaml:"blue" json:"blue"`
	White float64 `yaml:"white" json:"white"`
}

var name = "ledctl.json"

func (c Config) Save() error {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(name, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Load() (Config, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return createDefault()
	} else if err != nil {
		return Config{}, err
	}

	b, err := os.ReadFile(name)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func createDefault() (Config, error) {
	c := Config{
		StripType:  "rgb",
		GpioPin:    18,
		Brightness: 255,
		Segments: []Segment{
			{
				Id:   0,
				Leds: 100,
			},
		},
	}

	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return Config{}, err
	}

	err = os.WriteFile("ledctl.json", b, 0644)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
