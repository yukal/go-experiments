package gob

import (
	"fmt"
	"time"
)

type LocationWrapper struct {
	LocationName string
	Location     *time.Location `gob:"-"`
}

func (t LocationWrapper) GobEncode() ([]byte, error) {
	return []byte(t.LocationName), nil
}

func (t *LocationWrapper) GobDecode(data []byte) error {
	t.LocationName = string(data)

	loc, err := time.LoadLocation(t.LocationName)
	if err != nil {
		return fmt.Errorf("unable load location %s: %w", t.LocationName, err)
	}

	t.Location = loc
	return nil
}
