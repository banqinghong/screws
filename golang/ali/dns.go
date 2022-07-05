// This file is auto-generated, don't edit it. Thanks.
package ali

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	pvtz20180101 "github.com/alibabacloud-go/pvtz-20180101/client"
	"github.com/alibabacloud-go/tea/tea"
)


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateDnsClient (accessKeyId *string, accessKeySecret *string) (_result *pvtz20180101.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("pvtz.aliyuncs.com")
	_result = &pvtz20180101.Client{}
	_result, _err = pvtz20180101.NewClient(config)
	return _result, _err
}

func DescribeZones () (_err error) {
	client, _err := CreateDnsClient(tea.String(ReadOnlyAccessKey), tea.String(ReadOnlyAccessSecret))
	if _err != nil {
		return _err
	}

	describeZonesRequest := &pvtz20180101.DescribeZoneVpcTreeRequest{}
	// 复制代码运行请自行打印 API 的返回值
	resp, _err := client.DescribeZoneVpcTree(describeZonesRequest)
	if _err != nil {
		return _err
	}
	fmt.Println(resp.String())
	return _err
}

func DescribeZoneInfo () (_err error) {
	client, _err := CreateDnsClient(tea.String(ReadOnlyAccessKey), tea.String(ReadOnlyAccessSecret))
	if _err != nil {
		return _err
	}

	describeZoneInfoRequest := &pvtz20180101.DescribeZoneInfoRequest{
		ZoneId: tea.String("de02883774dfe7a842ed3074b55f90c3"),
	}
	// 复制代码运行请自行打印 API 的返回值
	resp, _err := client.DescribeZoneInfo(describeZoneInfoRequest)
	if _err != nil {
		return _err
	}
	fmt.Println(resp.Body.BindVpcs.String())
	return _err
}


