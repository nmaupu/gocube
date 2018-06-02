package data

import "testing"

func TestToggleDebug(t *testing.T) {
	SetDebug(false)
	if ToggleDebug(); !Debug {
		t.Errorf("ToggleDebug is incorrect")
	}
}
