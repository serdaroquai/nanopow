package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/crypto/blake2b"
)

const (
	digestLength = 8 // 8 bytes
	workLength   = 8 // 8 bytes
)

var (
	prodThreshold, _ = hex.DecodeString("ffffffc000000000")
	testThreshold, _ = hex.DecodeString("ffff000000000000")
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}
	block := request.QueryStringParameters["block"]
	if len(block) == 64 {
		work := ComputePoW(block)
		return events.APIGatewayProxyResponse{Body: work, StatusCode: 200}, nil
	} else {
		return events.APIGatewayProxyResponse{Body: "Invalid block value", StatusCode: 400}, nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}

func ComputePoW(prev string) string {
	previous, err := hex.DecodeString(prev)
	if err != nil {
		panic(err)
	}

	// generate a random work []byte
	for {
		work := make([]byte, workLength)
		rand.Read(work)
		work[workLength-1] = 0

		for j := 0; j < 256; j++ {
			work[workLength-1] = byte(j)
			if Validate(work, previous) {
				return hex.EncodeToString(reverse(work))
			}
		}
	}
}

func Validate(work []byte, prev []byte) bool {
	result, err := Blake2bFromBytes(work, prev)
	if err != nil {
		return false
	}
	return (bytes.Compare(result, prodThreshold) >= 0)
}

func Blake2bFromBytes(work []byte, prev []byte) ([]byte, error) {

	hash, err := blake2b.New(digestLength, nil)
	if err != nil {
		return nil, err
	}

	hash.Write(work)
	hash.Write(prev)
	return reverse(hash.Sum(nil)), nil
}

func reverse(src []byte) []byte {
	i := 0
	j := len(src) - 1
	for i < j {
		tmp := src[i]
		src[i] = src[j]
		src[j] = tmp
		i++
		j--
	}
	return src
}
