package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func NewClient() *ec2.Client {
	staticCredentialsProvider := credentials.NewStaticCredentialsProvider(AccessKeyId, SecretAccessKey, SessionToken)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(staticCredentialsProvider),
		config.WithDefaultRegion("ap-northeast-1"))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ec2.NewFromConfig(cfg)
	return client
}

// EC2DescribeInstancesAPI defines the interface for the DescribeInstances function.
// We use this interface to test the function using a mocked service.
type EC2DescribeInstancesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

// GetInstances retrieves information about your Amazon Elastic Compute Cloud (Amazon EC2) instances.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a DescribeInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to DescribeInstances.
func GetInstances(c context.Context, api EC2DescribeInstancesAPI, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return api.DescribeInstances(c, input)
}

func withRegion(region string) func(options *ec2.Options) {
	return func(options *ec2.Options) {
		options.Region = region
	}
}

func limit(num int32) func(options *ec2.DescribeInstancesPaginatorOptions) {
	return func(options *ec2.DescribeInstancesPaginatorOptions) {
		options.Limit = num
	}
}

func limitHost(num int32) func(options *ec2.DescribeHostsPaginatorOptions) {
	return func(options *ec2.DescribeHostsPaginatorOptions) {
		options.Limit = num
	}
}

func ListInstances() {
	client := NewClient()
	input := &ec2.DescribeInstancesInput{}

	inputPage := ec2.NewDescribeInstancesPaginator(client, input, limit(50))
	i := 0
	count := 0
	for {
		result, err := inputPage.NextPage(context.TODO(), withRegion("ap-northeast-1"))
		if err != nil {
			fmt.Println("query error: ", err)
			return
		}

		//jsonStr, _ := json.Marshal(result)
		//fmt.Println("result: ", string(jsonStr))
		for _, item := range result.Reservations {
			count += len(item.Instances)
		}
		fmt.Printf("第%d次查询，结果长度：%d\n", i, count)
		i++

	}


	//result, err := client.DescribeInstances(context.TODO(), input, withRegion("ap-northeast-1"))
	//if err != nil {
	//	fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
	//	fmt.Println(err)
	//	return
	//}
	//
	//for _, r := range result.Reservations {
	//	fmt.Println("Reservation ID: " + *r.ReservationId)
	//	fmt.Println("Instance IDs:")
	//	for _, i := range r.Instances {
	//		fmt.Println("   " + *i.InstanceId)
	//	}
	//
	//	fmt.Println("")
	//}
}

func ListRegions(){
	client := NewClient()
	input := &ec2.DescribeRegionsInput{
		AllRegions: tea.Bool(true),
	}
	output, err := client.DescribeRegions(context.TODO(), input)
	if err != nil {
		fmt.Println(err)
	}
	jsonStr, _ := json.Marshal(output)
	fmt.Println(string(jsonStr))
}

func ListHosts(){
	client := NewClient()
	input := &ec2.DescribeHostsInput{}

	inputPage := ec2.NewDescribeHostsPaginator(client, input, limitHost(5))
	i := 0
	count := 0
	for {
		result, err := inputPage.NextPage(context.TODO(), withRegion("ap-northeast-1"))
		if err != nil {
			fmt.Println("query error: ", err)
			return
		}

		jsonStr, _ := json.Marshal(result)
		fmt.Println("result: ", string(jsonStr))
		for _, item := range result.Hosts {
			count += len(item.Instances)
		}
		fmt.Printf("第%d次查询，结果长度：%d\n", i, count)
		i++

	}
}
