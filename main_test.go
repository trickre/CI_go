package main

import "testing"

func Test_subfunc(t *testing.T) {
	result := sub_func()

	if result != 0 {
		t.Fatal("failed test")
	}
}
