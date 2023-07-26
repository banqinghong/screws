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

func ListAllApis() {
	limit := 100
	offset := 0
	var apiSet []model.ApiInfoPerPage
	for {
		apis, err := ListApisV2(int32(offset), int32(limit))
		if err != nil {
			fmt.Printf("create route error: %s \n", err)
			return
		}
		apiSet = append(apiSet, apis...)
		if len(apis) < limit {
			break
		}
		offset += limit
	}

	fmt.Printf("%d api\n", len(apiSet))

	devcloudTitle := []string{"APi组名", "api名称", "请求URL", "转发url"}
	devcloudTitle = []string{"应用名", "独占成本", "分摊成本", "容器成本", "总成本"}
	var devcloudContent [][]string
	for _, api := range apiSet[:10] {
		fmt.Println("start ", api.Name)
		apiInfo, err := ShowDetailsOfApiV2(*api.Id)
		if err != nil {
			fmt.Printf("api %s 查询明细失败\n", api.Name)
			continue
		}
		if apiInfo.BackendApi == nil {
			fmt.Printf("api %s 未找到后端api转发\n", api.Name)
			continue
		}
		if api.ReqUri != apiInfo.BackendApi.ReqUri {
			content := []string{*api.GroupName, api.Name, api.ReqUri, apiInfo.BackendApi.ReqUri}
			devcloudContent = append(devcloudContent, content)
		}
	}
	devcloudExcelContent := &ExcelContent{
		Title:   devcloudTitle,
		Content: devcloudContent,
		OutFile: "1111111",
	}
	SaveExcel(devcloudExcelContent)
}

func ShowDetailsOfApiV2(apiID string) (*model.ShowDetailsOfApiV2Response, error) {
	builder := NewApigClientBuilder()
	regionID := "cn-east-3"
	apigRegion := region.ValueOf(regionID)
	client := apig.NewApigClient(builder.WithRegion(apigRegion).Build())

	request := &model.ShowDetailsOfApiV2Request{}
	request.InstanceId = ApigInstanceIDProd
	request.ApiId = apiID
	return client.ShowDetailsOfApiV2(request)
}

// 查询API
func ListApisV2(offset, limit int32) ([]model.ApiInfoPerPage, error) {
	builder := NewApigClientBuilder()
	regionID := "cn-east-3"
	apigRegion := region.ValueOf(regionID)
	client := apig.NewApigClient(builder.WithRegion(apigRegion).Build())

	request := &model.ListApisV2Request{}
	request.InstanceId = ApigInstanceIDProd
	request.Limit = &limit
	offsetInt64 := int64(offset)
	request.Offset = &offsetInt64
	response, err := client.ListApisV2(request)
	if err != nil {
		return nil, err
	}
	return *response.Apis, nil
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
		Name:         "x-oauth-token",
		Type:         paramType.STRING,
		Location:     paramBaseLocation.HEADER,
		Required:     &paramBaseRequired.E_1,
		Enumerations: tea.String("test"),
	}
	reqParamsList := []model.ReqParamBase{reqParams}

	vpcChannelInfoBackendApi := &model.ApiBackendVpcReq{
		VpcChannelId: "28a2fcfc6b0a493ca7d5fdf0dd9a7c17",
	}

	status := model.GetBackendApiCreateVpcChannelStatusEnum().E_1
	backendApiBody := &model.BackendApiCreate{
		VpcChannelInfo:   vpcChannelInfoBackendApi,
		ReqProtocol:      model.GetBackendApiCreateReqProtocolEnum().HTTPS,
		ReqMethod:        model.GetBackendApiCreateReqMethodEnum().ANY,
		ReqUri:           "/demo/api",
		Timeout:          5000,
		VpcChannelStatus: &status,
	}

	apiCreate := &model.ApiCreate{
		Name:        "test-rule",
		Type:        apiCreateTypeEnum.E_1,
		ReqProtocol: protocol.BOTH,
		ReqMethod:   method.ANY,
		ReqUri:      "/demo/api/v1",
		AuthType:    authType.NONE,
		MatchMode:   &matchType.SWA,
		GroupId:     "84c3ca45f6754344b16fa40e9d4f1996",
		ReqParams:   &reqParamsList,
		BackendType: backendType.HTTP,
		BackendApi:  backendApiBody,
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
	builder := NewApigClientBuilder()
	regionID := "cn-east-3"
	apigRegion := region.ValueOf(regionID)
	client := apig.NewApigClient(builder.WithRegion(apigRegion).Build())
	request := &model.ListCertificatesV2Request{}
	request.InstanceId = "8900209916fd497b93f3fd7fcb2f5f82"
	response, err := client.ListCertificatesV2(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}

// 创建分组
func CreateApigGroup() {
	builder := NewApigClientBuilder()
	regionID := "cn-east-3"
	apigRegion := region.ValueOf(regionID)
	client := apig.NewApigClient(builder.WithRegion(apigRegion).Build())

	apiGroupCreate := &model.ApiGroupCreate{
		Name: "devcloud-gateway-test-202205101436.dev.example.com",
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
			ClusterId:    CceClusterID, //pub cce 集群
			Namespace:    "monitoring",
			WorkloadType: model.GetMicroServiceInfoCceCreateWorkloadTypeEnum().STATEFULSET,
			AppName:      "",
			Port:         80,
			//Version:      "2e6f332a-dev",
		},
	}
	request := &model.ImportMicroserviceRequest{
		InstanceId: ApigInstanceID, //apig  instance id
		Body:       requestBody,
	}
	resp, err := client.ImportMicroservice(request)
	if err != nil {
		fmt.Printf("err in import microservice,%s \n", err.Error())
	}
	fmt.Println(resp)
}

func TestCreateVpcChannelV2() {
	ak := AccessKey
	sk := AccessSecret

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := apig.NewApigClient(
		apig.ApigClientBuilder().
			WithRegion(region.ValueOf("cn-east-3")).
			WithCredential(auth).
			Build())

	request := &model.CreateVpcChannelV2Request{}
	request.InstanceId = ApigInstanceID
	cceInfoMicroserviceInfo := &model.MicroServiceInfoCceBase{
		ClusterId:    CceDevopsPub,
		Namespace:    "qa-stage--devops-prod",
		WorkloadType: model.GetMicroServiceInfoCceBaseWorkloadTypeEnum().DEPLOYMENT,
		AppName:      "banqinghong-demo",
	}
	serviceTypeMicroserviceInfo := model.GetMicroServiceCreateServiceTypeEnum().CCE
	microserviceInfobody := &model.MicroServiceCreate{
		ServiceType: &serviceTypeMicroserviceInfo,
		CceInfo:     cceInfoMicroserviceInfo,
	}
	timeoutVpcHealthConfig := int32(5)
	vpcHealthConfigbody := &model.VpcHealthConfig{
		ThresholdNormal:   int32(2),
		Protocol:          model.GetVpcHealthConfigProtocolEnum().TCP,
		ThresholdAbnormal: int32(5),
		TimeInterval:      int32(10),
		Timeout:           &timeoutVpcHealthConfig,
	}
	var listMicroserviceLabelsMemberGroups = []model.MicroserviceLabel{
		{
			LabelName:  "appNameWithSuffix",
			LabelValue: "banqinghong-demo--cfvc0irlv24bmhb7n2mg",
		},
	}
	memberGroupWeightMemberGroups := int32(100)
	microservicePortMemberGroups := int32(80)
	var listMemberGroupsbody = []model.MemberGroupCreate{
		{
			MemberGroupName:    "banqinghong-demo--cfvc0irlv24bmhb7n2mg",
			MemberGroupWeight:  &memberGroupWeightMemberGroups,
			MicroservicePort:   &microservicePortMemberGroups,
			MicroserviceLabels: &listMicroserviceLabelsMemberGroups,
		},
	}
	typeVpcCreate := int32(3)
	request.Body = &model.VpcCreate{
		Type:             &typeVpcCreate,
		MicroserviceInfo: microserviceInfobody,
		VpcHealthConfig:  vpcHealthConfigbody,
		MemberGroups:     &listMemberGroupsbody,
		Name:             "dc-banqinghong-demo-prod-1678343536",
		Port:             int32(80),
		BalanceStrategy:  model.GetVpcCreateBalanceStrategyEnum().E_1,
		MemberType:       model.GetVpcCreateMemberTypeEnum().IP,
	}
	response, err := client.CreateVpcChannelV2(request)
	if err == nil {
		fmt.Printf("successful: %+v\n", response)
	} else {
		fmt.Println("failed: ", err)
	}
}
