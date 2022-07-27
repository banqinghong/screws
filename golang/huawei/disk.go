package huawei

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	evs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2"
	evsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/model"
	evsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/region"
)

func ListDisks() {
	client := evs.NewEvsClient(
		evs.EvsClientBuilder().WithRegion(evsRegion.CN_NORTH_1).
			WithCredential(
				basic.NewCredentialsBuilder().
					WithAk(AccessKey).
					WithSk(AccessSecret).
					Build()).
			WithHttpConfig(config.DefaultHttpConfig().
				WithIgnoreSSLVerification(true)).Build())
	request := &evsModel.ListVolumesRequest{}
	resp, err := client.ListVolumes(request)
	if err != nil {
		fmt.Println("list err: ", err)
		return
	}
	fmt.Println(resp.String())
}
