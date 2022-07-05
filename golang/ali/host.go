package ali

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateEcsClient (accessKeyId string, accessKeySecret string) (_result *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(config)
	return _result, _err
}

func DescribeInstanceStatus () (_err error) {
	client, _err := CreateEcsClient(ReadOnlyAccessKey, ReadOnlyAccessSecret)
	if _err != nil {
		return _err
	}
    region := "cn-hangzhou"
	instanceIDs := []string{"d-bp1dlbyfo4cf3tnmp30l"}
	jsonStr, _ := json.Marshal(instanceIDs)

	describeInstanceStatusRequest := &ecs20140526.DescribeDisksRequest{
		RegionId: &region,
		DiskIds:  tea.String(string(jsonStr)),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.DescribeDisks(describeInstanceStatusRequest)
	if _err != nil {
		return _err
	}
	fmt.Println(result)
	return _err
}


