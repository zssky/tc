package conv

import (
	"testing"
)

// TestConvToString - test interface Convert to string
func TestConvToString(t *testing.T) {
	params := []interface{}{
		int(10),
		int64(1024),
		float64(128.22),
		bool(true),
		string("test"),
		[]byte("bytes"),
	}

	for _, param := range params {
		t.Logf("%s", InterfaceToString(param))
	}
}
