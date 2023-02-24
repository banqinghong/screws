package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	elb2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func NewElbClient() *elb2.Client {
	staticCredentialsProvider := credentials.NewStaticCredentialsProvider(AccessKeyId, SecretAccessKey, SessionToken)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(staticCredentialsProvider),
		config.WithDefaultRegion("ap-northeast-1"))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := elb2.NewFromConfig(cfg)
	return client
}

func QueryELB2Listeners() {
	arn := "arn:aws:elasticloadbalancing:ap-northeast-1:129850821039:loadbalancer/app/aladdin-qa/f7cee15b63a1fd64"
	elbCli := NewElbClient()
	input := &elb2.DescribeListenersInput{
		LoadBalancerArn: &arn,
	}
	region := GetAwsRegion(arn)
	output, err := elbCli.DescribeListeners(context.Background(), input, func(o *elb2.Options) {
		o.Region = region
	})
	if err != nil {
		return
	}
	for _, listenerDescription := range output.Listeners {
		QueryELB2Rules(*listenerDescription.ListenerArn)
	}
}

func QueryELB2Rules(listenerArn string) {
	elbCli := NewElbClient()
	input := &elb2.DescribeRulesInput{
		ListenerArn: &listenerArn,
	}

	region := GetAwsRegion(listenerArn)
	output, err := elbCli.DescribeRules(context.Background(), input, func(o *elb2.Options) {
		o.Region = region
	})
	if err != nil {
		return
	}
	jsonStr, _ := json.Marshal(output)
	fmt.Println("rules", string(jsonStr))
	fmt.Println("----------------------------------------")
}

func QueryELB2TargetGroup() {
	arn := "arn:aws:elasticloadbalancing:us-west-2:129850821039:targetgroup/ucenter-manager-abroad/7906e008054d34da"
	elbCli := NewElbClient()
	input := &elb2.DescribeTargetGroupsInput{
		TargetGroupArns: []string{arn},
	}
	region := GetAwsRegion(arn)
	output, err := elbCli.DescribeTargetGroups(context.Background(), input, func(o *elb2.Options) {
		o.Region = region
	})
	if err != nil {
		fmt.Errorf("query target group wrong! err is:%s", err)
		return
	}
	if len(output.TargetGroups) < 0 {
		fmt.Errorf("target group not found")
		return
	}
	jsonStr, _ := json.Marshal(output.TargetGroups)
	fmt.Println("rules", string(jsonStr))
	fmt.Println("----------------------------------------")
	return
}

func GetAwsRegion(arn string) string {
	switch {
	case strings.Contains(arn, "us-east-2"):
		return "us-east-2"
	case strings.Contains(arn, "us-east-1"):
		return "us-east-1"
	case strings.Contains(arn, "us-west-1"):
		return "us-west-1"
	case strings.Contains(arn, "us-west-2"):
		return "us-west-2"
	case strings.Contains(arn, "af-south-1"):
		return "af-south-1"
	case strings.Contains(arn, "ap-east-1"):
		return "ap-east-1"
	case strings.Contains(arn, "ap-southeast-3"):
		return "ap-southeast-3"
	case strings.Contains(arn, "ap-south-1"):
		return "ap-south-1"
	case strings.Contains(arn, "ap-northeast-3"):
		return "ap-northeast-3"
	case strings.Contains(arn, "ap-northeast-2"):
		return "ap-northeast-2"
	case strings.Contains(arn, "ap-southeast-1"):
		return "ap-southeast-1"
	case strings.Contains(arn, "ap-southeast-2"):
		return "ap-southeast-2"
	case strings.Contains(arn, "ap-northeast-1"):
		return "ap-northeast-1"
	case strings.Contains(arn, "ca-central-1"):
		return "ca-central-1"
	case strings.Contains(arn, "eu-central-1"):
		return "eu-central-1"
	case strings.Contains(arn, "eu-west-1"):
		return "eu-west-1"
	case strings.Contains(arn, "eu-west-2"):
		return "eu-west-2"
	case strings.Contains(arn, "eu-south-1"):
		return "eu-south-1"
	case strings.Contains(arn, "eu-west-3"):
		return "eu-west-3"
	case strings.Contains(arn, "eu-north-1"):
		return "eu-north-1"
	case strings.Contains(arn, "me-south-1"):
		return "me-south-1"
	case strings.Contains(arn, "me-central-1"):
		return "me-central-1"
	case strings.Contains(arn, "sa-east-1"):
		return "sa-east-1"
	case strings.Contains(arn, "me-central-1"):
		return "me-central-1"
	case strings.Contains(arn, "us-gov-east-1"):
		return "us-gov-east-1"
	case strings.Contains(arn, "us-gov-west-1"):
		return "us-gov-west-1"
	default:
		return "UnkonwnRegion"
	}
}
