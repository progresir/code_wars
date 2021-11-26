package main

import (
	"testing"
)

func TestPartList(t *testing.T) {
	res := Thirt(8529)
	if res != 89 {
		t.Errorf("error: %v", res)
	}
}

func TestPartList2(t *testing.T) {
	res := Thirt(85299258)
	if res != 31 {
		t.Errorf("error: %v", res)
	}
}
