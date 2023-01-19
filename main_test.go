package main

import "testing"

func Test(t *testing.T) {
	test := "success"

	if test != "success" {
		t.Fail()
	}
}
