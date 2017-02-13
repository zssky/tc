package uuid

import (
	"testing"
)

// TestUUID - test uuid created
func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		uuid := NewUUID4()
		t.Log(uuid.String())
	}
}
