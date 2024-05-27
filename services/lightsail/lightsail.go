package lightsail

import (
	"fmt"
	"sync"

	"aws_test/services"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
)

type LightsailService struct {
	svc *lightsail.Lightsail
}

func NewLightsailService(sess *session.Session) services.AWSService {
	return &LightsailService{
		svc: lightsail.New(sess),
	}
}

func (ls *LightsailService) ListResources() error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	instanceCount := 0

	input := &lightsail.GetInstancesInput{}

	for {
		result, err := ls.svc.GetInstances(input)
		if err != nil {
			return fmt.Errorf("failed to get instances: %w", err)
		}

		for _, instance := range result.Instances {
			wg.Add(1)
			go func(instance *lightsail.Instance) {
				defer wg.Done()
				mu.Lock()
				instanceCount++
				mu.Unlock()
				fmt.Printf("Instance: %s\n", aws.StringValue(instance.Name))
			}(instance)
		}

		if result.NextPageToken == nil {
			break
		}

		input.PageToken = result.NextPageToken
	}

	wg.Wait()

	fmt.Printf("Total instances: %d\n", instanceCount)
	return nil
}
