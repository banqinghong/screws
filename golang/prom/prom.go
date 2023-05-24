package prom

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func QueryCpuUsage() {
	queryCPUUsageAvg := fmt.Sprintf(queryFmtCPUUsageAvg, "2d", "")
	rawQuery := url.QueryEscape(queryCPUUsageAvg)
	fmt.Println("query: ", rawQuery)

	client, err := api.NewClient(api.Config{
		Address: "http://10.63.116.213:9090",
		// RoundTripper: tr,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	// client.URL().RawQuery =

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, queryCPUUsageAvg, time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	// fmt.Printf("Result:\n%v\n", result)
	jsonStr, _ := json.Marshal(result)
	fmt.Println(string(jsonStr))
}

func HttpGet() {
	// 查询可用的k8s部署集群和命名空间，取首个可用的资源
	httpCli := http.DefaultClient
	httpCli.Timeout = 10 * time.Second

	containerDeployUrl := &url.URL{
		Scheme: "HTTP",
		Host:   "prometheus.dev.xgimi.com",
	}
	containerDeployUrl.Path = "/api/v1/query"

	queryCPUUsageAvg := fmt.Sprintf(queryFmtCPUUsageAvg, "2d", "")

	// request query
	containerDeployQuery := url.Values{}
	containerDeployQuery.Set("query", queryCPUUsageAvg)

	containerDeployUrl.RawQuery = containerDeployQuery.Encode()

	containerDeployReq, _ := http.NewRequest("POST", containerDeployUrl.String(), nil)
	clusterDeployResp, err := httpCli.Do(containerDeployReq)
	if err != nil {
		fmt.Printf("do request failed: %s\n", err)
		return
	}
	defer clusterDeployResp.Body.Close()
	respBodyBytes, _ := ioutil.ReadAll(clusterDeployResp.Body)
	if clusterDeployResp.StatusCode > 200 {
		fmt.Printf("code %d\n", clusterDeployResp.StatusCode)
		return
	}

	fmt.Println("result: ", string(respBodyBytes))
}
