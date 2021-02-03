package osInfo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

const AliMetaDataUrl = "http://100.100.100.200/latest/meta-data"

// 获取centos系统的版本信息
func GetCentOSVersion() (string, error) {
	// CentOS系统中/etc/system-release-cpe存在版本信息
	// cpe:/o:centos:centos:7
	cmd := exec.Command("cat", "/etc/system-release-cpe")
	osInfoStr, err := cmd.Output()
	if err != nil {
		return "", err
	}
	osInfoList := strings.Split(string(osInfoStr), ":")
	return osInfoList[4], nil
}

// 获取阿里云机器instance_id
func GetInstanceIDInAliyun() (string, error) {
	url := fmt.Sprintf("%s/instance-id", AliMetaDataUrl)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	instanceId, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(instanceId), nil
}

// 获取阿里云机器region-id
func GetRegionIDInAliyun() (string, error) {
	url := fmt.Sprintf("%s/region-id", AliMetaDataUrl)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	instanceId, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(instanceId), nil
}
