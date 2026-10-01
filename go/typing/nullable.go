package typing

import (
	"bytes"
	"encoding/json"
)

var nullByte = []byte("null")

// this struct's primary purpose is to accurately represent
// nullable zero values without being omitted by json's omitempty
type Nullable[T any] struct {
	Value T
	Valid bool
}

func (n Nullable[T]) IsZero() bool {
	return !n.Valid
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	res, err := json.Marshal(n.Value)
	return res, err
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		n.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &n.Value); err != nil {
		return err
	}

	n.Valid = true
	return nil
}
