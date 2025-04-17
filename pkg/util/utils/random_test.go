package utils

import "testing"

func TestGetRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(GetRandom(1, 100))
	}
}
