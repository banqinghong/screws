package huawei

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/banqinghong/screws/golang/btime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	cdn "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1/region"
)

func NewCdnClientBuilder() *core.HcHttpClientBuilder {
	credentials := global.NewCredentialsBuilder().WithAk(AccessKey).WithSk(AccessSecret).Build()
	httpConfig := config.DefaultHttpConfig()
	builder := cdn.CdnClientBuilder().WithCredential(credentials).WithHttpConfig(httpConfig)
	return builder
}

func ShowDomainStats() {
	builder := NewCdnClientBuilder()
	regionID := "cn-north-1"
	cdnRegion := region.ValueOf(regionID)
	client := cdn.NewCdnClient(builder.WithRegion(cdnRegion).Build())

	endTime := btime.GetNearestTime(5 * time.Minute)
	startTime := endTime.Add(-10 * time.Minute)

	req := &model.ShowDomainStatsRequest{
		Action:              "detail",
		StartTime:           startTime.UnixMilli(),
		EndTime:             endTime.UnixMilli(),
		DomainName:          "all",
		StatType:            "flux",
		GroupBy:             tea.String("domain"),
		EnterpriseProjectId: tea.String("all"),
	}

	resp, err := client.ShowDomainStats(req)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	jsonStr, _ := json.Marshal(resp)
	fmt.Println(string(jsonStr))
}
