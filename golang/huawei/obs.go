package huawei

import (
	"encoding/json"
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

func NewObsClient() (*obs.ObsClient, error) {
	return obs.New(AccessKey, AccessSecret, ObsEndpoint)
}

func UploadFileToObs(objectName, localFileName string) {
	input := &obs.PutFileInput{}
	input.Bucket = "devops-qa"
	input.Key = objectName
	input.SourceFile = localFileName

	obsCli, err := NewObsClient()
	if err != nil {
		fmt.Printf("new obs client failed: %s\n", err)
		return
	}

	output, err := obsCli.PutFile(input)
	if err != nil {
		fmt.Printf("put file to oss failed: %s\n", err)
		return
	}
	jsonStr, _ := json.Marshal(output)
	fmt.Println("output: ", string(jsonStr))
}
