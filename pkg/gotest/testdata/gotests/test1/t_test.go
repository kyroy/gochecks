package test1

import "testing"

func TestMe(t *testing.T) {
	x()
}

func TestTwo(t *testing.T) {
	t.Fatalf("my fail msg")
}
