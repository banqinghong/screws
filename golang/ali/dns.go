// This file is auto-generated, don't edit it. Thanks.
package ali

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	pvtz20180101 "github.com/alibabacloud-go/pvtz-20180101/client"
	"github.com/alibabacloud-go/tea/tea"
	"time"
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

func DescribeZoneRecords () (_err error) {
	client, _err := CreateDnsClient(tea.String(ReadOnlyAccessKey), tea.String(ReadOnlyAccessSecret))
	if _err != nil {
		return _err
	}

	describeZoneRecordsRequest := &pvtz20180101.DescribeZoneRecordsRequest{
		ZoneId: tea.String("6c491ff079874b54d757265fb9d2aa93"),
		SearchMode: tea.String("EXACT"),
		Keyword: tea.String("lift.pre"),
	}
	// 复制代码运行请自行打印 API 的返回值
	resp, _err := client.DescribeZoneRecords(describeZoneRecordsRequest)
	if _err != nil {
		return _err
	}
	fmt.Println(resp.String())
	return _err
}

func DescribeChangeLogs () {
	client, _err := CreateDnsClient(tea.String(ReadOnlyAccessKey), tea.String(ReadOnlyAccessSecret))
	if _err != nil {
		fmt.Println("new client error: ", _err)
		return
	}

	describeZonesRequest := &pvtz20180101.DescribeZoneVpcTreeRequest{}
	// 复制代码运行请自行打印 API 的返回值
	zoneResp, _err := client.DescribeZoneVpcTree(describeZonesRequest)
	if _err != nil {
		fmt.Println("query zones error: ", _err)
		return
	}


	for _, v := range zoneResp.Body.Zones.Zone {
		fmt.Println(tea.StringValue(v.ZoneName))
		endTime := time.Now().UnixNano() / 1e6
		startTime := time.Now().Add(-24 * time.Hour).UnixNano()/1e6
		describeRecordLogs := &pvtz20180101.DescribeChangeLogsRequest{
			ZoneId: v.ZoneId,
			EndTimestamp: tea.Int64(endTime),
			StartTimestamp: tea.Int64(startTime),
		}
		fmt.Println("req: ", describeRecordLogs.String())
		resp, _err := client.DescribeChangeLogs(describeRecordLogs)
		if _err != nil {
			fmt.Println("query error: ", _err)
			return
		}
		fmt.Println(resp.Body.String())
	}
}


