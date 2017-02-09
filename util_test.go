package tc

import (
	"testing"
)

// TestTrimSpace - test trim space
func TestTrimSpace(t *testing.T) {
	ss := []string{
		"    test1",
		"test2    ",
		" test3   ",
		" test 4  ",
	}

	for _, s := range ss {
		t.Logf("|%v|", TrimSpace(s))
	}

}

// TestTrimSplit - test trim split
func TestTrimSplit(t *testing.T) {
	ss := []struct {
		sep string
		raw string
	}{
		{";", " test1;test2 ; test3 ;test 4"},
		{",", " test1,test2 , test3 ,test 4"},
		{"#", " test1#test2 # test3 #test 4"},
	}

	for _, s := range ss {
		t.Logf("%v", TrimSplit(s.raw, s.sep))
	}
}

// TestGenUID - test generate unique id
func TestGenUID(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Logf("%v", GenUID())
	}

}
