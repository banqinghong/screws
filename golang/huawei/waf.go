package huawei

import (
	"encoding/json"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	waf "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/region"
)

func NewWafClientBuilder() *core.HcHttpClientBuilder {
	credentials := basic.NewCredentialsBuilder().WithAk(AccessKey).WithSk(AccessSecret).Build()
	httpConfig := config.DefaultHttpConfig()
	builder := waf.WafClientBuilder().WithCredential(credentials).WithHttpConfig(httpConfig)
	return builder
}

type WafHostInfo struct {
	HostID   string
	HostName string
}

func ListWafHost() {
	builder := NewWafClientBuilder()
	regionID := "cn-east-3"
	cdnRegion := region.ValueOf(regionID)
	wafCli := waf.NewWafClient(builder.WithRegion(cdnRegion).Build())

	request := &model.ListHostRequest{}

	enterpriseProjectIdRequest := "all_granted_eps"
	request.EnterpriseProjectId = &enterpriseProjectIdRequest
	ps := int32(20)
	pn := int32(1)
	request.Pagesize = &ps
	request.Page = &pn

	groupList := make([]*WafHostInfo, 0)

	for {
		items, err := wafCli.ListHost(request)
		if err != nil {
			fmt.Println("list host failed: ", err)
			return
		}

		for _, v := range *items.Items {
			hostID := tea.StringValue(v.Hostid)
			if hostID != "" {
				group := &WafHostInfo{
					HostID:   hostID,
					HostName: tea.StringValue(v.Hostname),
				}
				groupList = append(groupList, group)
			}
		}

		// 判断是否是尾页了
		if len(*items.Items) < 1 {
			break
		}

		if len(*items.Items) < 20 {
			break
		}

		// 继续读下一页数据
		*(request.Page) += 1
	}

	fmt.Println("total host: ", len(groupList))
	jsonStr, _ := json.Marshal(groupList)
	fmt.Println(string(jsonStr))
}
