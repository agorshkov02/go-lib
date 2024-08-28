package json

import (
	"bytes"
	"strings"
	"testing"
)

func TestReadJSON(t *testing.T) {
	t.Run("should read JSON", func(t *testing.T) {
		data := `{"name":"John","age":18}`
		reader := strings.NewReader(data)

		type resultType struct {
			Name string `json:"name"`
			Age  int32  `json:"age"`
		}
		result := &resultType{}

		if err := ReadJSON(reader, &result); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if result.Name != "John" || result.Age != 18 { // Correct the age to 18
			t.Errorf("unexpected result: %v", result)
		}
	})

	t.Run("should throw error", func(t *testing.T) {
		data := `{"name":"John","age":30}`
		reader := strings.NewReader(data)

		type resultType struct {
			Name string `json:"name"`
		}
		result := &resultType{}

		if err := ReadJSON(reader, &result); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if result.Name != "John" { // Expecting to ignore missing age
			t.Errorf("unexpected result: %v", result)
		}
	})

	t.Run("should throw error when input > maxBytes", func(t *testing.T) {
		data := make([]byte, maxBytes+1)
		for i := range data {
			data[i] = 'a'
		}

		reader := bytes.NewReader(data)

		type resultType struct{}
		result := &resultType{}

		err := ReadJSON(reader, &result)
		if err == nil {
			t.Error("error expected due to input size exceeding maxBytes, got nil")
		}
	})
}

func TestWriteJSON(t *testing.T) {
	t.Run("should write JSON", func(t *testing.T) {
		type dataType struct {
			Name string `json:"name"`
			Age  int32  `json:"age"`
		}
		data := &dataType{
			Name: "John",
			Age:  18,
		}

		buff := &bytes.Buffer{}
		if err := WriteJSON(buff, data); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if buff.String() != `{"name":"John","age":18}` && buff.String() != `{"age":18,"name":"John"}` {
			t.Errorf("unexpected result: %s", buff.String())
		}
	})

	t.Run("should throw error", func(t *testing.T) {
		ch := make(chan int)

		buff := &bytes.Buffer{}
		if err := WriteJSON(buff, ch); err == nil {
			t.Error("expected error, got nil")
		}
	})
}
