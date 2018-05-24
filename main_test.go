package log

import (
	"testing"
)

func TestConnect(t *testing.T) {
	err := Connect("", true)
	if err != nil {
		t.Fatal(err.Error())
	}
}
