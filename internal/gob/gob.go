package gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func Encode(data any) (bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(data); err != nil {
		return bytes.Buffer{}, fmt.Errorf("unable encode data: %w", err)
	}

	return buf, nil
}

func Decode[T any](content []byte) (T, error) {
	var data T
	dec := gob.NewDecoder(bytes.NewReader(content))

	if err := dec.Decode(&data); err != nil {
		return *new(T), fmt.Errorf("unable decode settings: %w", err)
	}

	return data, nil
}
