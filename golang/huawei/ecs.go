package huawei

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"

	bss "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2"
	bssModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	bssRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/region"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
)

func NewEcsClientBuilder() *core.HcHttpClientBuilder {
	credentials := basic.NewCredentialsBuilder().WithAk(AccessKey).WithSk(AccessSecret).Build()
	httpConfig := config.DefaultHttpConfig().WithHttpHandler(httphandler.
		NewHttpHandler().
		AddRequestHandler(RequestHandler).
		AddResponseHandler(ResponseHandler))
	builder := ecs.EcsClientBuilder().WithCredential(credentials).WithHttpConfig(httpConfig)
	//client := ecs.NewEcsClient(
	//	ecs.EcsClientBuilder().WithRegion(region.CN_NORTH_1).
	//		WithCredential(
	//			basic.NewCredentialsBuilder().
	//				WithAk(AccessKey).
	//				WithSk(AccessSecret).
	//				Build()).
	//		WithHttpConfig(config.DefaultHttpConfig().
	//			WithIgnoreSSLVerification(true)).Build())
	return builder
}

func ListEcs() {
	client := ecs.NewEcsClient(NewEcsClientBuilder().WithRegion(region.CN_NORTH_1).Build())
	//client.HcClient.WithEndpoint(region.CN_NORTH_1.Endpoint)

	request := &model.ListServersDetailsRequest{}
	resp, err := client.ListServersDetails(request)
	if err != nil {
		fmt.Println("list err: ", err)
		return
	}
	fmt.Println(resp.String())
}

//func DescribeEcs() {
//	client := NewEcsClient()
//	request := &model.ShowServerRequest{
//		ServerId: "355322cd-354b-44ea-8c6d-91bcd13e1ef4",
//	}
//	resp, err := client.ShowServer(request)
//	if err != nil {
//		fmt.Println("list err: ", err)
//		return
//	}
//	fmt.Println(resp.String())
//}
//
func ListRegions()  {
	auth := basic.NewCredentialsBuilder().
		WithAk(AccessKey).
		WithSk(AccessSecret).
		Build()

	client := ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(region.ValueOf("cn-north-1")).
			WithCredential(auth).
			Build())

	request := &model.NovaListAvailabilityZonesRequest{}
	resp, err := client.NovaListAvailabilityZones(request)
	if err != nil {
		fmt.Println("list err: ", err)
		return
	}
	fmt.Println(resp.String())
}

func GetOrderInfo() {
	client := bss.NewBssClient(
		bss.BssClientBuilder().WithRegion(bssRegion.CN_NORTH_1).
			WithCredential(
				global.NewCredentialsBuilder().
					WithAk(AccessKey).
					WithSk(AccessSecret).
					Build()).
			WithHttpConfig(config.DefaultHttpConfig().
				WithIgnoreSSLVerification(true)).Build())
    resourceIDs := []string{"df46de1b-97bc-493f-95e8-8157e93c7608"}
	request := &bssModel.ListPayPerUseCustomerResourcesRequest{
		Body: &bssModel.QueryResourcesReq{
			ResourceIds: &resourceIDs,
			OnlyMainResource: tea.Int32(1),
		},
	}
	resp, err := client.ListPayPerUseCustomerResources(request)
	if err != nil {
		fmt.Println("list err: ", err)
		return
	}
	fmt.Println(resp.String())
}

