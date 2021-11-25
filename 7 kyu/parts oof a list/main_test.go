package main

import (
	"testing"
)

func TestPartList(t *testing.T) {
	data := []string{"I", "wish", "I", "hadn't", "come"}
	res := PartList(data)
	if res != "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)" {
		t.Errorf("Snils verification error: %s", res)
	}
}
