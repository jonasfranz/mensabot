package openmensa

import (
	"encoding/json"
	"errors"
)

// A Coordinate object
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

// UnmarshalJSON override for Coordinates
func (c *Coordinate) UnmarshalJSON(data []byte) error {
	tmp := []interface{}{&c.Latitude, &c.Longitude}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if len(tmp) != 2 {
		return errors.New("Coordinate contains more than two values.")
	}

	return nil
}
