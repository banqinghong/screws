// This file is auto-generated, don't edit it. Thanks.
package ali

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	mse20190531 "github.com/alibabacloud-go/mse-20190531/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateGWClient (accessKeyId *string, accessKeySecret *string) (_result *mse20190531.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("mse.cn-hangzhou.aliyuncs.com")
	_result = &mse20190531.Client{}
	_result, _err = mse20190531.NewClient(config)
	return _result, _err
}

func AddGatewayRoute () error {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}
	services0 := &mse20190531.AddGatewayRouteRequestServices{
		ServiceId: tea.Int64(2618),
		Percent: tea.Int32(0),
		Name: tea.String("gateway-demo-pub-svc"),
		Namespace: tea.String("default"),
		SourceType: tea.String("K8S"),
	}
	services1 := &mse20190531.AddGatewayRouteRequestServices{
		ServiceId: tea.Int64(2618),
		Percent: tea.Int32(100),
		Name: tea.String("gateway-demo-pub-svc"),
		Namespace: tea.String("default"),
		SourceType: tea.String("K8S"),
	}
	predicatesPathPredicates := &mse20190531.AddGatewayRouteRequestPredicatesPathPredicates{
		Path: tea.String("/demo"),
		Type: tea.String("PRE"),
		IgnoreCase: tea.Bool(true),
	}
	predicates := &mse20190531.AddGatewayRouteRequestPredicates{
		PathPredicates: predicatesPathPredicates,
	}
	//domainList := []int64{1034}
	//domainListJson, _ := json.Marshal(domainList)
	addGatewayRouteRequest := &mse20190531.AddGatewayRouteRequest{
		GatewayUniqueId: tea.String(GatewayID),
		Name: tea.String("gateway-test-local-mvcc"),
		Predicates: predicates,
		Services: []*mse20190531.AddGatewayRouteRequestServices{services0, services1},
		//DomainIdListJSON: tea.String(string(domainListJson)),
		DomainId: tea.Int64(1079),
		DestinationType: tea.String("VersionOriented"),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.AddGatewayRoute(addGatewayRouteRequest)
	if _err != nil {
		return _err
	}
	fmt.Println("add: ", result.String())
	routeID := fmt.Sprintf("%d", tea.Int64Value(result.Body.Data))
	err := ApplyGateway(routeID, GatewayID)
	if err != nil {
		return err
	}
	return nil
}

func ApplyGateway (routeID, gatewayUniqueId string) (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}

	listGatewayRequest := &mse20190531.ApplyGatewayRouteRequest{
		RouteId: tea.String(routeID),
		GatewayUniqueId: tea.String(gatewayUniqueId),
	}
	fmt.Println(listGatewayRequest.String())
	// 复制代码运行请自行打印 API 的返回值
	gateway, _err := client.ApplyGatewayRoute(listGatewayRequest)
	if _err != nil {
		return _err
	}
	fmt.Println(gateway)
	return _err
}

func UpdateGatewayRoute () error {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}
	services0 := &mse20190531.UpdateGatewayRouteRequestServices{
		ServiceId: tea.Int64(3067),
		Percent: tea.Int32(80),
		Version: tea.String("7266d669-dev"),
		Name: tea.String("gateway-test--ca1kofjlv24ag4u3njsg"),
		Namespace: tea.String("devops-test"),
		SourceType: tea.String("K8S"),
	}
	services1 := &mse20190531.UpdateGatewayRouteRequestServices{
		ServiceId: tea.Int64(2975),
		Percent: tea.Int32(20),
		Version: tea.String("version.7266d669-dev"),
		Name: tea.String("gateway-test--c9t0b1blv2401kk94h2g"),
		Namespace: tea.String("devops-test"),
		SourceType: tea.String("K8S"),
	}
	predicatesPathPredicates := &mse20190531.UpdateGatewayRouteRequestPredicatesPathPredicates{
		Path: tea.String("/demo"),
		Type: tea.String("PRE"),
		IgnoreCase: tea.Bool(true),
	}
	predicates := &mse20190531.UpdateGatewayRouteRequestPredicates{
		PathPredicates: predicatesPathPredicates,
	}
	domainList := []int64{1293}
	domainListJson, _ := json.Marshal(domainList)
	addGatewayRouteRequest := &mse20190531.UpdateGatewayRouteRequest{
		Services: []*mse20190531.UpdateGatewayRouteRequestServices{services0, services1},
		Id: tea.Int64(2751),
		DomainIdListJSON: tea.String(string(domainListJson)),
		Predicates: predicates,
		DestinationType: tea.String("VersionOriented"),
		GatewayUniqueId: tea.String(GatewayID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.UpdateGatewayRoute(addGatewayRouteRequest)
	if _err != nil {
		return _err
	}
	fmt.Println("update: ", result.String())
	routeID := fmt.Sprintf("%d", tea.Int64Value(result.Body.Data))
	err := ApplyGateway(routeID, GatewayID)
	if err != nil {
		return err
	}
	return nil
}

func ListGatewayDomain () (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}

	listGatewayDomainRequest := &mse20190531.ListGatewayDomainRequest{
		GatewayUniqueId: tea.String(GatewayID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.ListGatewayDomain(listGatewayDomainRequest)
	if _err != nil {
		return _err
	}
	//domainID := tea.Int64Value(result.Body.Data)
	fmt.Println(result)
	return _err
}

func ListGatewayRoute () (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}

	filterParams := &mse20190531.ListGatewayRouteRequestFilterParams{
		GatewayUniqueId: tea.String(GatewayID),
		DomainName:      tea.String("devcloud-gateway-test-202205101436.dev.xgimi.com"),
		Status:          tea.Int32(1),
	}

	listGatewayRouteRequest := &mse20190531.ListGatewayRouteRequest{
		FilterParams: filterParams,
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.ListGatewayRoute(listGatewayRouteRequest)
	if _err != nil {
		return _err
	}
	//domainID := tea.Int64Value(result.Body.Data)
	fmt.Println(result)

	return _err
}

func ListSSLCert () (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}

	addGatewayDomainRequest := &mse20190531.ListSSLCertRequest{
		GatewayUniqueId: tea.String(GatewayID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.ListSSLCert(addGatewayDomainRequest)
	if _err != nil {
		return _err
	}
	//domainID := tea.Int64Value(result.Body.Data)
	fmt.Println(result)
	return _err
}

func UpdateGatewayDomain () (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}

	updateGatewayDomainRequest := &mse20190531.UpdateGatewayDomainRequest{
		Id: tea.Int64(1079),
		Protocol: tea.String("HTTPS"),
		CertIdentifier: tea.String("6879295-cn-hangzhou"),
		MustHttps: tea.Bool(true),
		GatewayUniqueId: tea.String(GatewayID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.UpdateGatewayDomain(updateGatewayDomainRequest)
	if _err != nil {
		return _err
	}
	//domainID := tea.Int64Value(result.Body.Data)
	fmt.Println(result)
	return _err
}

func AddGatewayDomain () (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}

	addGatewayDomainRequest := &mse20190531.AddGatewayDomainRequest{
		Name: tea.String("gateway-demo.i.xgimi.com"),
		Protocol: tea.String("HTTP"),
		GatewayUniqueId: tea.String(GatewayID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.AddGatewayDomain(addGatewayDomainRequest)
	if _err != nil {
		return _err
	}
	//domainID := tea.Int64Value(result.Body.Data)
	fmt.Println(result)
	return _err
}

func ImportServices () (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}

	serviceList0 := &mse20190531.ImportServicesRequestServiceList{
		Name: tea.String("gateway-demo-pub-svc"),
		Namespace: tea.String("default"),
	}
	serviceList1 := &mse20190531.ImportServicesRequestServiceList{
		Name: tea.String("gateway-demo-srr-svc"),
		Namespace: tea.String("default"),
	}
	importServicesRequest := &mse20190531.ImportServicesRequest{
		ServiceList: []*mse20190531.ImportServicesRequestServiceList{serviceList0, serviceList1},
		SourceType: tea.String("K8S"),
		GatewayUniqueId: tea.String(GatewayID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.ImportServices(importServicesRequest)
	if _err != nil {
		return _err
	}
	//domainID := tea.Int64Value(result.Body.Data)
	fmt.Println(result)
	return _err
}

func ListGatewayService () (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}

	filterParams := &mse20190531.ListGatewayServiceRequestFilterParams{
		//Name: tea.String("gateway-demo-pub-svc"),
		GatewayUniqueId: tea.String("gw-e324c43f475d49fa8fb81cc1c5d73c63"),
	}
	listGatewayServiceRequest := &mse20190531.ListGatewayServiceRequest{
		FilterParams: filterParams,
		PageSize: tea.Int32(200),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.ListGatewayService(listGatewayServiceRequest)
	if _err != nil {
		return _err
	}
	//domainID := tea.Int64Value(result.Body.Data)
	fmt.Println(result)
	return _err
}

func PullServices () (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}

	listGatewayServiceRequest := &mse20190531.PullServicesRequest{
		GatewayUniqueId: tea.String(GatewayID),
		SourceType: tea.String("K8S"),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.PullServices(listGatewayServiceRequest)
	if _err != nil {
		return _err
	}
	fmt.Println(result)
	return _err
}

type LabelInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ServiceVersion struct {
	Name   string       `json:"name"`
	Labels []*LabelInfo `json:"labels"`
}

func AddGatewayServiceVersion () (_err error) {
	client, _err := CreateGWClient(tea.String(AccessKey), tea.String(AccessSecret))
	if _err != nil {
		return _err
	}
    labelInfo := &LabelInfo{
    	Key: "app",
    	Value: "gateway-demo-pub",
	}
	serviceVersion := &ServiceVersion{
		Name: "v2",
		Labels: []*LabelInfo{labelInfo},
	}
	jsonStr, _ := json.Marshal(serviceVersion)
	addGatewayServiceVersionRequest := &mse20190531.AddGatewayServiceVersionRequest{
		ServiceId: tea.Int64(2618),
		ServiceVersion: tea.String(string(jsonStr)),
		GatewayUniqueId: tea.String(GatewayID),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.AddGatewayServiceVersion(addGatewayServiceVersionRequest)
	if _err != nil {
		return _err
	}
	//domainID := tea.Int64Value(result.Body.Data)
	fmt.Println(result)
	return _err
}




