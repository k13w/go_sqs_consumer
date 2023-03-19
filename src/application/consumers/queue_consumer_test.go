package main

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"sync"
	"testing"
)

func BenchmarkPollMessages(b *testing.B) {
	wg := new(sync.WaitGroup)
	wg.Add(b.N)

	done := make([]chan<- *sqs.Message, b.N)

	for i := 0; i < b.N; i++ {
		done[i] = make(chan<- *sqs.Message)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		go pollMessages(done[i])
	}
}
