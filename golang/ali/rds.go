package ali

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	rds "github.com/alibabacloud-go/rds-20140815/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateRdsClient (accessKeyId string, accessKeySecret string) (_result *rds.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("rds-cn-hangzhou.aliyuncs.com")
	_result = &rds.Client{}
	_result, _err = rds.NewClient(config)
	return _result, _err
}

func ListRds () (_err error) {
	client, _err := CreateRdsClient(ReadOnlyAccessKey, ReadOnlyAccessSecret)
	if _err != nil {
		return _err
	}
	region := "cn-hangzhou"
	describeInstancesRequest := &rds.DescribeDBInstancesRequest{
		RegionId: &region,
		PageSize: tea.Int32(100),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.DescribeDBInstances(describeInstancesRequest)
	if _err != nil {
		return _err
	}
	//fmt.Println(result)
	for _, instance := range result.Body.Items.DBInstance {
		instanceID := tea.StringValue(instance.DBInstanceId)
		//if tea.StringValue(instance.ConnectionMode) != "Standard" {
		//	fmt.Println(instanceID)
		//}
		err := DescribeDBProxyEndpoint(instanceID)
		if err != nil {
			continue
		}
	}
	return _err
}

// rm-bp18kk332en2rqr7d
func DescribeRds (instanceID string) (_err error) {
	client, _err := CreateRdsClient(ReadOnlyAccessKey, ReadOnlyAccessSecret)
	if _err != nil {
		return _err
	}
	//region := "cn-hangzhou"
	describeInstanceRequest := &rds.DescribeDBInstanceAttributeRequest{
		DBInstanceId: tea.String(instanceID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.DescribeDBInstanceAttribute(describeInstanceRequest)
	if _err != nil {
		return _err
	}
	fmt.Println(result)
	return _err
}

// 连接地址
func DescribeNetInfo (instanceID string) (_err error) {
	client, _err := CreateRdsClient(ReadOnlyAccessKey, ReadOnlyAccessSecret)
	if _err != nil {
		return _err
	}
	//region := "cn-hangzhou"
	describeInstanceRequest := &rds.DescribeDBInstanceNetInfoRequest{
		DBInstanceId: tea.String(instanceID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.DescribeDBInstanceNetInfo(describeInstanceRequest)
	if _err != nil {
		return _err
	}
	//fmt.Println(result)
	for _, item := range result.Body.DBInstanceNetInfos.DBInstanceNetInfo {
		fmt.Printf("%s   %s\n", instanceID, tea.StringValue(item.ConnectionString))
	}
	return _err
}

func DescribeDBProxyEndpoint (instanceID string) (_err error) {
	client, _err := CreateRdsClient(ReadOnlyAccessKey, ReadOnlyAccessSecret)
	if _err != nil {
		return _err
	}
	//region := "cn-hangzhou"
	describeInstanceRequest := &rds.DescribeDBProxyEndpointRequest{
		DBInstanceId: tea.String(instanceID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.DescribeDBProxyEndpoint(describeInstanceRequest)
	if _err != nil {
		return _err
	}
	//fmt.Println(result)
	fmt.Printf("%s   %s\n", instanceID, tea.StringValue(result.Body.DBProxyConnectString))
	return _err
}


