package aws

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func NewMqClient() *mq.Client {
	staticCredentialsProvider := credentials.NewStaticCredentialsProvider(AccessKeyId, SecretAccessKey, SessionToken)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(staticCredentialsProvider),
		config.WithDefaultRegion("ap-northeast-1"))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := mq.NewFromConfig(cfg)
	return client
}

func QueryMqBrokers() {
	// arn := "arn:aws:elasticloadbalancing:ap-northeast-1:129850821039:loadbalancer/app/aladdin-qa/f7cee15b63a1fd64"
	elbCli := NewMqClient()
	input := &mq.ListBrokersInput{}

	output, err := elbCli.ListBrokers(context.Background(), input)
	if err != nil {
		return
	}

	jsonStr, _ := json.Marshal(output)
	fmt.Println("res: ", string(jsonStr))
}
