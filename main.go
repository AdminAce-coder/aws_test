package main

import (
	"log"

	"aws_test/config"
	"aws_test/services/lightsail"
)

func main() {
	sess, err := config.NewAWSSession("us-west-2")
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}

	lightsailService := lightsail.NewLightsailService(sess)

	err = lightsailService.ListResources()
	if err != nil {
		log.Fatalf("failed to list resources: %v", err)
	}
}
