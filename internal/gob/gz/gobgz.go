package gobz

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io"
)

func Encode(data any) ([]byte, error) {
	const op = "gobgz.Encode"
	var encBuf bytes.Buffer

	enc := gob.NewEncoder(&encBuf)
	if err := enc.Encode(data); err != nil {
		return nil, fmt.Errorf("%s.Encode: %w", op, err)
	}

	var gzBuf bytes.Buffer
	zw := gzip.NewWriter(&gzBuf)

	if _, err := zw.Write(encBuf.Bytes()); err != nil {
		zw.Close()
		return nil, fmt.Errorf("%s.Write: %w", op, err)
	}

	if err := zw.Close(); err != nil {
		return nil, fmt.Errorf("%s.Close: %w", op, err)
	}

	return gzBuf.Bytes(), nil
}

func Decode[T any](content []byte) (T, error) {
	const op = "gobgz.Decode"

	gz, err := gzip.NewReader(bytes.NewReader(content))
	if err != nil {
		return *new(T), fmt.Errorf("%s.NewReader: %w", op, err)
	}
	defer gz.Close()

	decompressedContent, err := io.ReadAll(gz)
	if err != nil {
		return *new(T), fmt.Errorf("%s.ReadAll: %w", op, err)
	}

	var data T

	dec := gob.NewDecoder(bytes.NewReader(decompressedContent))
	if err := dec.Decode(&data); err != nil {
		return *new(T), fmt.Errorf("%s.Decode: %w", op, err)
	}

	return data, nil
}

func EncodeAs(writer io.Writer, data any) error {
	const op = "gobgz.EncodeAs"
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("%s.Encode: %w", op, err)
	}

	gzw := gzip.NewWriter(writer)

	if _, err := gzw.Write(buf.Bytes()); err != nil {
		gzw.Close()
		return fmt.Errorf("%s.Write: %w", op, err)
	}

	if err := gzw.Close(); err != nil {
		return fmt.Errorf("%s.Close: %w", op, err)
	}

	return nil
}

func DecodeAs[T any](reader io.Reader, data T) error {
	const op = "gobgz.DecodeAs"

	gz, err := gzip.NewReader(reader)
	if err != nil {
		return fmt.Errorf("%s.NewReader: %w", op, err)
	}
	defer gz.Close()

	decompressedContent, err := io.ReadAll(gz)
	if err != nil {
		return fmt.Errorf("%s.ReadAll: %w", op, err)
	}

	dec := gob.NewDecoder(bytes.NewReader(decompressedContent))
	if err := dec.Decode(&data); err != nil {
		return fmt.Errorf("%s.Decode: %w", op, err)
	}

	return nil
}
