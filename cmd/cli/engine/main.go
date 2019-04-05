package main

import (
	"fmt"
	"os"

	loadsim "github.com/dave-malone/aws-iot-loadsimulator/pkg"
)

const (
	one_million  int = 1000000
	one_thousand int = 1000
)

func main() {
	fmt.Println("Running aws-iot-loadsimulator engine")

	sns_topic_arn := os.Getenv("SNS_TOPIC_ARN")
	if len(sns_topic_arn) == 0 {
		fmt.Println("Environment variable SNS_TOPIC_ARN not set")
		return
	}

	config := &loadsim.SnsEventEngineConfig{
		TargetTotalConcurrentThings: one_thousand * 10,
		ClientsPerWorker:            one_thousand,
		MessagesToGeneratePerClient: 10,
		AwsRegion:                   "us-east-1",
		AwsSnsTopicARN:              sns_topic_arn,
		SecondsBetweenEachEvent:     10,
	}

	engine := loadsim.NewSnsEventEngine(config)
	if err := engine.GenerateEvents(); err != nil {
		fmt.Printf("Failed to generate events: %v", err)
		return
	}

	fmt.Println("Simulation requests generated")
}