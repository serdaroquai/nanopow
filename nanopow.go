package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	// "github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/crypto/blake2b"
)

const (
	digestLength       = 8  // 8 bytes
	workLength         = 8  // 8 bytes
	messageLength      = 32 // 32 bytes
	numWorkers         = 4  // max 256
	zeroByte      byte = 0
)

var (
	prodThreshold, _ = hex.DecodeString("ffffffc000000000")
	testThreshold, _ = hex.DecodeString("ff00000000000000")
)

func main() {

	// lambda.Start(HandleRequest)

	input, _ := hex.DecodeString("B53387DCE4553F480665E92126DB3022AF7CBD77CC1060E33659A951DB5BC2BA") //Big Endian
	result := Solve(input, prodThreshold, numWorkers)
	// result,_ := hex.DecodeString("388e72b262c2bbb3") // Little Endian
	reverse(result) // Convert to BigEndian
	// resultEncoded := hex.EncodeToString(result)

	fmt.Println(hex.EncodeToString(Blake2b(append(result, input...))))
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	block := request.QueryStringParameters["block"]
	if len(block) != 64 {
		return events.APIGatewayProxyResponse{Body: "Invalid block value", StatusCode: 400}, nil
	}
	input, _ := hex.DecodeString(block)

	result := Solve(input, prodThreshold, numWorkers)
	resultEncoded := hex.EncodeToString(result)

	return events.APIGatewayProxyResponse{Body: resultEncoded, StatusCode: 200}, nil
}

// Solve computes a Proof Of Work that is above given threshold for the given "previous block hash"
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

		reverse(result) // convert little endian to big endian for comparison
		if bytes.Compare(result, threshold) >= 0 {
			// found proof of work
			done <- work
			break
		}

		hash.Reset()

		// increment digits
		for j := workLength - 1; j >= 0; j-- {
			work[j]++
			if work[j] != zeroByte {
				break
			}
		}

		//increment i
		i++
	}
}

func Blake2b(input []byte) []byte {
	hash, _ := blake2b.New(digestLength, nil)
	hash.Write(input)
	return hash.Sum(nil)
}

func reverse(src []byte) {
	i := 0
	j := len(src) - 1
	for i < j {
		tmp := src[i]
		src[i] = src[j]
		src[j] = tmp
		i++
		j--
	}
}
