package main

import "testing"

func TestSum(t *testing.T) {
	if sum(1,2) != 3 {
		t.Error("sum error")
	}

	if sum(2,4) != 5 { //わざと間違える
		t.Error("sum error 2")
	}
}