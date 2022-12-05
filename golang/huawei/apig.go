package huawei

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	apig "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apig/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apig/v2/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apig/v2/region"
)

func NewApigClientBuilder() *core.HcHttpClientBuilder {
	credentials := basic.NewCredentialsBuilder().WithAk(AccessKey).WithSk(AccessSecret).Build()
	httpConfig := config.DefaultHttpConfig()
	builder := apig.ApigClientBuilder().WithCredential(credentials).WithHttpConfig(httpConfig)
	return builder
}

// 新建规则
func CreateApigRoute() {
	builder := NewApigClientBuilder()
	regionID := "cn-east-3"
	apigRegion := region.ValueOf(regionID)
	client := apig.NewApigClient(builder.WithRegion(apigRegion).Build())

	apiCreateTypeEnum := model.GetApiCreateTypeEnum()
	protocol := model.GetApiCreateReqProtocolEnum()
	method := model.GetApiCreateReqMethodEnum()
	authType := model.GetApiCreateAuthTypeEnum()
	matchType := model.GetApiCreateMatchModeEnum()
	backendType := model.GetApiCreateBackendTypeEnum()

	paramType := model.GetReqParamBaseTypeEnum()
	paramBaseLocation := model.GetReqParamBaseLocationEnum()
	paramBaseRequired := model.GetReqParamBaseRequiredEnum()



	reqParams := model.ReqParamBase{
		Name: "x-oauth-token",
		Type: paramType.STRING,
		Location: paramBaseLocation.HEADER,
		Required: &paramBaseRequired.E_1,
		Enumerations: tea.String("test"),
	}
	reqParamsList := []model.ReqParamBase{reqParams}

	vpcChannelInfoBackendApi := &model.ApiBackendVpcReq{
		VpcChannelId: "28a2fcfc6b0a493ca7d5fdf0dd9a7c17",
	}

	status := model.GetBackendApiCreateVpcChannelStatusEnum().E_1
	backendApiBody := &model.BackendApiCreate{
		VpcChannelInfo: vpcChannelInfoBackendApi,
		ReqProtocol: model.GetBackendApiCreateReqProtocolEnum().HTTPS,
		ReqMethod: model.GetBackendApiCreateReqMethodEnum().ANY,
		ReqUri: "/demo/api",
		Timeout: 5000,
		VpcChannelStatus: &status,
	}

	apiCreate := &model.ApiCreate{
		Name:          "test-rule",
		Type:          apiCreateTypeEnum.E_1,
		ReqProtocol:   protocol.BOTH,
		ReqMethod:     method.ANY,
		ReqUri:        "/demo/api/v1",
		AuthType:      authType.NONE,
		MatchMode:     &matchType.SWA,
		GroupId:       "84c3ca45f6754344b16fa40e9d4f1996",
		ReqParams:     &reqParamsList,
		BackendType:   backendType.HTTP,
		BackendApi:    backendApiBody,
	}
	request := &model.CreateApiV2Request{
		InstanceId: ApigInstanceID,
		Body:       apiCreate,
	}
	resp, err := client.CreateApiV2(request)
	if err != nil {
		fmt.Printf("create route error: %s \n", err)
		return
	}
	fmt.Println("successful: ", resp.String())
	return
}

// 查询分组
func ListApigGroup() {
	builder := NewApigClientBuilder()
	regionID := "cn-east-3"
	apigRegion := region.ValueOf(regionID)
	client := apig.NewApigClient(builder.WithRegion(apigRegion).Build())

	request := &model.ListApiGroupsV2Request{
		InstanceId: ApigInstanceID,
		Limit:      tea.Int32(500),
	}
	resp, err := client.ListApiGroupsV2(request)
	if err != nil {
		fmt.Printf("list region[%s] error: %s \n", regionID, err)
		return
	}
	fmt.Println("successful: ", resp.String())

}

// 查询证书
func ListSslCert() {
}

// 创建分组
func CreateApigGroup() {
	builder := NewApigClientBuilder()
	regionID := "cn-east-3"
	apigRegion := region.ValueOf(regionID)
	client := apig.NewApigClient(builder.WithRegion(apigRegion).Build())

	apiGroupCreate := &model.ApiGroupCreate{
		Name: "devcloud-gateway-test-202205101436.dev.xgimi.com",
	}
	request := &model.CreateApiGroupV2Request{
		InstanceId: ApigInstanceID,
		Body:       apiGroupCreate,
	}
	resp, err := client.CreateApiGroupV2(request)
	if err != nil {
		fmt.Printf("create region[%s] error: %s \n", regionID, err)
		return
	}
	fmt.Println("create successful: ", resp.String())

}

func ImportMicroService() {

	builder := NewApigClientBuilder()
	regionID := "cn-east-3"
	apigRegion := region.ValueOf(regionID)
	client := apig.NewApigClient(builder.WithRegion(apigRegion).Build())

	apiList := make([]model.MicroserviceApiCreate, 0)
	reqMethod := model.GetMicroserviceApiCreateReqMethodEnum().ANY
	reqProtocol := model.GetMicroserviceImportReqProtocolEnum().HTTP
	reqMatchMode := model.GetMicroserviceApiCreateMatchModeEnum().SWA
	curApi := model.MicroserviceApiCreate{
		Name:      tea.String("apig-api-test"),
		ReqMethod: &reqMethod,
		ReqUri:    "/demo/api",
		MatchMode: &reqMatchMode,
	}
	apiList = append(apiList, curApi)
	requestBody := &model.MicroserviceImportReq{
		GroupInfo: &model.MicroserviceGroup{
			GroupId: tea.String("d1f4af2b2d184b2cb1d93e02d5cd0a47"),
		},
		ServiceType: model.GetMicroserviceImportReqServiceTypeEnum().CCE,
		Protocol:    &reqProtocol,
		Apis:        apiList,
		Cors:        tea.Bool(true),
		CceInfo: &model.MicroServiceInfoCceCreate{
			ClusterId:    "552bd312-3d77-11ed-94b1-0255ac1002cd", //pub cce 集群
			Namespace:    "monitoring",
			WorkloadType: model.GetMicroServiceInfoCceCreateWorkloadTypeEnum().STATEFULSET,
			AppName:      "",
			Port:         80,
			//Version:      "2e6f332a-dev",
		},
	}
	request := &model.ImportMicroserviceRequest{
		InstanceId: "0ab28a92697c4f1cac25e1296940c2a2", //apig  instance id
		Body:       requestBody,
	}
	resp, err := client.ImportMicroservice(request)
	if err != nil {
		fmt.Printf("err in import microservice,%s \n", err.Error())
	}
	fmt.Println(resp)
}


