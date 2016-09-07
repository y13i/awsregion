package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
)

func main() {
	if os.Getenv("AWS_REGION") == "" {
		os.Setenv("AWS_REGION", "us-east-1")
	}

	client := ec2.New(session.New())

	response, err := client.DescribeRegions(nil)

	if err != nil {
		panic(err)
	}

	for _, region := range response.Regions {
		fmt.Println(*region.RegionName)
	}
}
