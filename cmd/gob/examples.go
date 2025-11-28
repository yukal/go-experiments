package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"yu/golang/internal/gob"
	gobz "yu/golang/internal/gob/gz"
)

func SaveGob(filepath string, data any) error {
	var (
		buf bytes.Buffer
		err error
	)

	if buf, err = gob.Encode(data); err != nil {
		return fmt.Errorf("SaveGob.Encode: %w", err)
	}

	if err := os.WriteFile(filepath, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("SaveGob.WriteFile: %w", err)
	}

	return nil
}

func LoadGob[T any](filepath string) (T, error) {
	var (
		content []byte
		data    T
		err     error
	)

	if content, err = os.ReadFile(filepath); err != nil {
		return *new(T), fmt.Errorf("LoadGob.ReadFile: %w", err)
	}

	if data, err = gob.Decode[T](content); err != nil {
		return *new(T), fmt.Errorf("LoadGob.Decode: %w", err)
	}

	return data, nil
}

func SaveGobz(fileName string, data any) {
	gobFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer gobFile.Close()

	if err := gobz.EncodeAs(data, gobFile); err != nil {
		log.Fatalf("unable encode gob: %v", err)
	}
}

func LoadGobz[T any](fileName string) {
	gobFile, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer gobFile.Close()

	var data T

	if err := gobz.DecodeAs(gobFile, &data); err != nil {
		log.Fatalf("unable encode gob: %v", err)
	}
}
