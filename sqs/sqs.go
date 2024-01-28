package sqs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	queueName = "local-queue"
)

func createQueue(svc *sqs.SQS) (queueUrl string, err error) {
	params := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}

	result, err := svc.CreateQueue(params)
	if err != nil {
		return queueUrl, err
	}

	queueUrl = *result.QueueUrl
	return queueUrl, nil
}

func sendMessage(svc *sqs.SQS, queueURL string, messageBody string) (msgID string, err error) {
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    aws.String(queueURL),
	}

	result, err2 := svc.SendMessage(params)
	if err2 != nil {
		return msgID, err2
	}

	return *result.MessageId, nil
}

func receiveMessages(svc *sqs.SQS, queueURL string) ([]*sqs.Message, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(10),
	}

	resp, err := svc.ReceiveMessage(params)
	if err != nil {
		return nil, err
	}

	return resp.Messages, nil
}

func Init() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("test", "test", ""),
		Endpoint:    aws.String("http://localhost:4566"),
	})

	if err != nil {
		fmt.Println("error in creating session")
		return
	}

	svc := sqs.New(sess)

	// Creating a Queue
	queueUrl, err := createQueue(svc)
	if err != nil {
		fmt.Println("Error creating queue:", err)
		return
	}
	fmt.Println("Queue created:", queueUrl)

	// Sending a Message
	msgID, err2 := sendMessage(svc, queueUrl, "This is Body of Message")
	if err2 != nil {
		fmt.Println("Error in Sending Message to Queue:", err2)
		return
	}
	fmt.Println("Message Sent : ", msgID)

	// recieving a message
	recievedMessage, err3 := receiveMessages(svc, queueUrl)
	if err3 != nil {
		fmt.Println("Error in Recieving Message from queue", err3)
		return
	}

	for _, msg := range recievedMessage {
		fmt.Println("Message Recieved : ", *msg.Body)
	}

}
