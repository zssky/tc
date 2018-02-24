package ips

import (
	"testing"
)

func TestLocalIPv4s(t *testing.T) {
	ips, err := LocalIPv4s()
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("success:%v", ips)
}

func TestGetIPv4ByInterface(t *testing.T) {
	ips, err := GetIPv4ByInterface("wlp3s0")
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("success:%v", ips)
}
