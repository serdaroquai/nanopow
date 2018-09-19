package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"testing"
)

var result []byte

func TestWork(t *testing.T) {
	threshold, _ := hex.DecodeString("ffffffc000000000")
	input, _ := hex.DecodeString("C08C7727AC85E6DCC26D13B2FB9083AF05C17616C4999B966C2BBCD1586398E6") 
	work, _ := hex.DecodeString("ebd042008df3b2be")                                                  
	hash := Blake2b(append(reverse(work), input...))
	if compare(threshold, hash) >= 0 {
		t.Fail()
	}
}

func TestCompare(t *testing.T) {
	m := []byte{255, 2, 1}
	n := []byte{1, 1, 130}

	if compare(m, n) != 1 {
		t.Fail()
	}

	n = []byte{1, 2, 255}

	if compare(m, n) != 0 {
		t.Fail()
	}

	n = []byte{2, 2, 255}

	if compare(m, n) != -1 {
		t.Fail()
	}
}

func BenchmarkSolve(b *testing.B) {
	var r []byte
	threshold, _ := hex.DecodeString("ff00000000000000")
	for n := 0; n < b.N; n++ {

		bytes := make([]byte, 32)
		if _, err := rand.Read(bytes); err != nil {
			b.Fail()
		}

		// always record the result to prevent
		// the compiler eliminating the function call.
		r = Solve(bytes, threshold, 256)
	}
	
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
