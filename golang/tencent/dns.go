package tencent

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)


func NewDnsPodClient()(*dnspod.Client, error) {
	credential := common.NewCredential(AccessKey,AccessSecret,)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 300
	cpf.Language = "en-US"

	dnsConn, _ := dnspod.NewClient(credential, regions.Beijing, cpf)
	return dnsConn, nil
}

func DescribeDomains()(){
	client, _ :=  NewDnsPodClient()
	request := dnspod.NewDescribeDomainListRequest()
	resp, err := client.DescribeDomainList(request)
	if err != nil {
		fmt.Println("query domain list error: ", err)
		return
	}
	fmt.Println("domain: ", resp.ToJsonString())
}

func DescribeDomainLog()  {
	client, _ :=  NewDnsPodClient()
	request := dnspod.NewDescribeDomainListRequest()
	domainResp, err := client.DescribeDomainList(request)
	if err != nil {
		fmt.Println("query domain list error: ", err)
		return
	}

	for _, v := range domainResp.Response.DomainList {
		describeDomainLogListRequest := dnspod.NewDescribeDomainLogListRequest()
		describeDomainLogListRequest.DomainId = v.DomainId
		describeDomainLogListRequest.Domain = v.Name
		logResp, err := client.DescribeDomainLogList(describeDomainLogListRequest)
		if err != nil {
			fmt.Printf("query domain[%s] log list error: %s\n", v.Name, err)
			continue
		}
		fmt.Println("log resp: ")
		fmt.Println(logResp.ToJsonString())
	}
}

func DescribeDomainRecords()(){
	client, _ :=  NewDnsPodClient()
	request := dnspod.NewDescribeRecordListRequest()
	limit := uint64(1000)
	domainID := uint64(3091997)
	request.Limit = &limit
	request.Domain = tea.String("xgimi.com")
	request.DomainId = &domainID
	resp, err := client.DescribeRecordList(request)
	if err != nil {
		fmt.Println("query domain list error: ", err)
		return
	}
	fmt.Println("domain: ", resp.ToJsonString())
}
