package ali

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	waf_openapi20190910 "github.com/alibabacloud-go/waf-openapi-20190910/client"
)

func CreateWafClient () (_result *waf_openapi20190910.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: tea.String(AccessKey),
		// 您的 AccessKey Secret
		AccessKeySecret: tea.String(AccessSecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("wafopenapi.cn-hangzhou.aliyuncs.com")
	_result = &waf_openapi20190910.Client{}
	_result, _err = waf_openapi20190910.NewClient(config)
	return _result, _err
}

func DescribeWafInfo() error{
	client, _ := CreateWafClient()
	describeInstanceInfoRequest := &waf_openapi20190910.DescribeInstanceInfoRequest{}
	result, err := client.DescribeInstanceInfo(describeInstanceInfoRequest)
	if err != nil {
		fmt.Println("describe waf info error: ", err)
		return err
	}
	fmt.Println("resp: ", result)
	return nil
}
