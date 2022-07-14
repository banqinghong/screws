package huawei

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/region"
	"net/http"
)

func NewVpcClient() *vpc.VpcClient {
	client := vpc.NewVpcClient(
		vpc.VpcClientBuilder().WithRegion(region.CN_NORTH_1).
			WithCredential(
				basic.NewCredentialsBuilder().
					WithAk(AccessKey).
					WithSk(AccessSecret).
					Build()).
			WithHttpConfig(config.DefaultHttpConfig().
				WithIgnoreSSLVerification(true).
				WithHttpHandler(httphandler.
					NewHttpHandler().
					AddRequestHandler(RequestHandler).
					AddResponseHandler(ResponseHandler))).
			Build())
	return client
}

func RequestHandler(request http.Request) {
	fmt.Println("req")
	fmt.Println(request)
}

func ResponseHandler(response http.Response) {
	fmt.Println("response")
	fmt.Println(response)
}

func ListVpc() {
	client := NewVpcClient()

	limit := int32(1)
	request := &model.ListVpcsRequest{
		Limit: &limit,
	}
	response, err := client.ListVpcs(request)
	if err == nil {
		fmt.Printf("%+v\n\n", response.Vpcs)
	} else {
		fmt.Println(err)
	}
}
