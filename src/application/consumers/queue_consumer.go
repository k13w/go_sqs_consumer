package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"sync"

	"fmt"
)

var (
	sqsSvc *sqs.SQS
)

func pollMessages(chn chan<- *sqs.Message, wg *sync.WaitGroup) {

	for {
		output, err := sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String("https://sqs.us-east-1.amazonaws.com/624638737469/query_bus"),
			MaxNumberOfMessages: aws.Int64(10),
			WaitTimeSeconds:     aws.Int64(20),
		})

		if err != nil {
			fmt.Println("failed to fetch sqs message %v", err)
		}

		for _, message := range output.Messages {
			chn <- message
		}

		wg.Done()
	}
}

func handleMessage(msg *sqs.Message) {
	fmt.Println("RECEIVING MESSAGE >>> ")
	fmt.Println(*msg.Body)
}

func deleteMessage(msg *sqs.Message) {
	sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String("https://sqs.us-east-1.amazonaws.com/624638737469/query_bus"),
		ReceiptHandle: msg.ReceiptHandle,
	})
}

func main() {

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	sqsSvc = sqs.New(sess)

	chnMessages := make(chan *sqs.Message, 2)
	go pollMessages(chnMessages)

	for message := range chnMessages {
		handleMessage(message)
		deleteMessage(message)
	}
}
