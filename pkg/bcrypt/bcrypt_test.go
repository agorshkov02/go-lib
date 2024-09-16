package bcrypt

import (
	"testing"
)

func TestSum256(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		sum256 := Sum256("abc123")

		if len(sum256) == 0 {
			t.Errorf("unxpected result: %s", sum256)
		}
	})
}
