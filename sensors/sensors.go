package sensors

import (
	"io/ioutil"
	"path/filepath"
)

const (
	FAN  = "fan"
	IN   = "in"
	TEMP = "temp"
)

var prefixGlobs []string

type Sensor struct {
	Name   string
	Inputs map[string][]Input
}

type Input struct {
	Type  string
	Label string
	Val   float64
	Max   float64
	Crit  float64
}

func init() {
	for _, prefix := range []string{FAN, IN, TEMP} {
		prefixGlobs = append(prefixGlobs, prefix+"*")
	}
}

func Sensors() ([]Sensor, error) {
	sensors := []Sensor{}
	path := "/sys/class/hwmon/"

	devs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// Try base hwmonX, hwmonX/device directories
	for _, d := range devs {
		dirs := []string{path + d.Name(), path + d.Name() + "/device/"}
		for _, dir := range dirs {
			if isInfoDir(dir) {
				s, err := sensorFromDir(dir)
				if err != nil {
					return nil, err
				}
				sensors = append(sensors, s)
			}
		}
	}

	return sensors, nil
}

// Extract sensor information for a directory
func sensorFromDir(dir string) (Sensor, error) {
	s := Sensor{}

	// Name
	name, err := ioutil.ReadFile(dir + "/name")
	if err != nil {
		return s, err
	}
	s.Name = string(name)

	// Temperature

	return s, nil
}

// Determines if a given directory contains any of the following files:
// - fanX_
// - tempX_
// - inX_
func isInfoDir(dir string) bool {
	for _, glob := range prefixGlobs {
		if files, _ := filepath.Glob(dir + glob); len(files) > 0 {
			return true
		}
	}
	return false
}
