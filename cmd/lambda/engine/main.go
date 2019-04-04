package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	loadsim "github.com/dave-malone/aws-iot-loadsimulator/pkg"
)

const (
	//TODO - externalize this config
	sns_topic_arn string = "arn:aws:sns:us-east-1:068311527115:iot_simulator_notifications"
	one_million   int    = 1000000
	one_thousand  int    = 1000
)

func main() {
	lambda.Start(requestHandler)
}

func requestHandler(ctx context.Context) (string, error) {
	config := &loadsim.SnsEventEngineConfig{
		TargetTotalConcurrentThings: one_thousand * 10,
		ClientsPerWorker:            one_thousand,
		MessagesToGeneratePerClient: 10,
		AwsRegion:                   os.Getenv("AWS_REGION"),
		AwsSnsTopicARN:              sns_topic_arn,
		SecondsBetweenEachEvent:     10,
	}

	engine := loadsim.NewSnsEventEngine(config)
	if err := engine.GenerateEvents(); err != nil {
		return "", fmt.Errorf("Failed to generate events: %v", err)
	}

	return "Simulation requests generated", nil
}
