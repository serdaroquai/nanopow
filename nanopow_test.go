package main

import (
	"bytes"
	"testing"
)

var result string

func BenchmarkComputePoW(b *testing.B) {
	var r string
	// for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = ComputePoW("5F0BE06A322FF1334C56F55B881CAD7383FC521CFF1545F4DEEFCE05FE35EC60")
	// }
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func TestSwap(t *testing.T) {

	src := []byte{}
	if bytes.Compare(reverse(src), []byte{}) != 0 {
		t.Fail()
	}

	src = []byte{1}
	if bytes.Compare(reverse(src), []byte{1}) != 0 {
		t.Fail()
	}

	src = []byte{1, 2}
	if bytes.Compare(reverse(src), []byte{2, 1}) != 0 {
		t.Fail()
	}

	src = []byte{1, 2, 3}
	if bytes.Compare(reverse(src), []byte{3, 2, 1}) != 0 {
		t.Fail()
	}
}
