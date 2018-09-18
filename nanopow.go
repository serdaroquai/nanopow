package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"golang.org/x/crypto/blake2b"
)

const (
	digestLength       = 8   // 8 bytes
	workLength         = 8   // 8 bytes
	messageLength      = 32  // 32 bytes
	numWorkers         = 256 // max 256
	zeroByte      byte = 0
)

var (
	prodThreshold, _ = hex.DecodeString("ffffffc000000000")
	testThreshold, _ = hex.DecodeString("ff00000000000000")
)

func main() {

	thresholdString := flag.String("t", "ffffffc000000000", "threshold value that proof must fulfill")
	inputString := flag.String("h", "C08C7727AC85E6DCC26D13B2FB9083AF05C17616C4999B966C2BBCD1586398E6", "previous block hash to be used as input")

	flag.Parse()

	input, err := hex.DecodeString(*inputString)
	if err != nil || len(input) != messageLength {
		fmt.Println("Invalid previous block hash")
		return
	}

	threshold, _ := hex.DecodeString(*thresholdString)
	if err != nil || len(threshold) != digestLength {
		fmt.Println("Invalid threshold value")
		return
	}

	work := Solve(input, threshold, numWorkers)

	fmt.Println(hex.EncodeToString(work))
}

// Solve computes a Proof Of Work that is above given threshold for a given "previous block hash"
func Solve(input []byte, threshold []byte, numberOFWorkers int) []byte {
	done := make(chan []byte)

	for i := 0; i < numberOFWorkers; i++ {
		message := append(make([]byte, workLength), input...)
		work := message[:workLength]
		work[0] = byte((256 / numberOFWorkers) * i)

		go startWorker(message, done, threshold)
	}

	result := <-done
	return result
}

func startWorker(input []byte, done chan []byte, threshold []byte) {
	work := input[:workLength]
	hash, _ := blake2b.New(digestLength, nil)

	for i := 0; ; {

		hash.Write(input)
		result := hash.Sum(nil)

		if compare(threshold, result) < 0 {
			// found proof of work
			done <- reverse(work)
			break
		}

		hash.Reset()

		// calculate next works
		for j := workLength - 1; j >= 0; j-- {
			work[j]++
			if work[j] != zeroByte {
				break
			}
		}

		i++
	}
}

// compares a to reverse of b in place lexicographically
// assumes a and b are equal length
func compare(a, b []byte) int {

	i := 0
	j := len(b) - 1

	for j >= 0 {
		if a[i] > b[j] {
			return 1
		}
		if a[i] < b[j] {
			return -1
		}
		i++
		j--
	}
	return 0
}

func Blake2b(input []byte) []byte {
	hash, _ := blake2b.New(digestLength, nil)
	hash.Write(input)
	return hash.Sum(nil)
}

func reverse(src []byte) []byte {
	length := len(src)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[length-i-1] = src[i]
	}

	return result
}
