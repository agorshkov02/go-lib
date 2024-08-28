package json

import (
	"encoding/json"
	"io"
)

const (
	maxBytes = 1 * 1024 * 1024 // 1 MB
)

func ReadJSON(r io.Reader, v any) error {
	buff := make([]byte, uint64(maxBytes))
	n, err := r.Read(buff)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buff[:n], v); err != nil {
		return err
	}
	return nil
}

func WriteJSON(w io.Writer, v any) error {
	buff, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, err := w.Write(buff); err != nil {
		return err
	}
	return nil
}
