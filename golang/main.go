package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tjfoc/gmsm/x509"
	// "github.com/banqinghong/screws/golang/aws"
)

func main() {
	fmt.Println("main starting")
	//s := []float64{1, 3, 3}
	//fmt.Println("sum: ", number.GetAvgOfFloatList(s))
	//bfile.ReadFileLine("/tmp/test.txt")
	//timeStr := "2021-03-04 02:10:00.000"
	//time, _ := btime.String2Time(timeStr)
	//lastMonth := time.AddDate(0, -1, 0).Format("2006-01")
	//fmt.Println(lastMonth)
	//kafka.ConsumerTest(11)
	//_ = ali.DescribeInstanceStatus()
	//err := ali.DescribeZonesVpc()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//ali.DescribeChangeLogs()
	//huawei.ImportMicroService()
	//huawei.CreateApigRoute()
	// huawei.ListAllApis()
	//ali.DescribeChangeLogs()
	//tencent.DescribeDomainRecords()
	//aws.ListInstances()
	// aws.QueryELB2TargetGroup()
	//ali.DescribeChangeLogs()
	//ldap.ExampleSearch()
	// ldap.Run()
	// apollo.GetConfig()
	//cname, err := net.LookupCNAME("k8s-pub-tokyo-prometheus-i.xgimi.com")
	//if err != nil {
	//	log.Printf("query domain record failed: %s\n", err)
	//} else {
	//	fmt.Printf("%v\n", cname)
	//}
	//v4, _ := uuid.NewV4()
	//
	//fmt.Println(v4.String()
	// HttpPostJson()

	
	fmt.Println("main ending")
}

func getSecret() (string, error) {
	completeKey := `04C223AA7AF8D325413BF0A3EAC18535ACEF64A84D53A2856E2D0C0FE486900D9E7856415EBE686442E9BC36F857B813EF70C852D43E1978AAC02C8B2FEFA80779`
	// block, rest := pem.Decode([]byte(completeKey)) //将密钥解析成公钥实例
	// if block == nil {
	// 	fmt.Println(string(rest))
	// 	fmt.Println("0")
	// 	return
	// }

	pub, err := x509.ReadPublicKeyFromHex(completeKey) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		fmt.Println("1", err)
		return "", err
	}

	// encrtpted, err := sm2.Encrypt(pub, []byte("123"), rand.Reader, 1)
	// if err != nil {
	// 	fmt.Println("2", err)
	// 	return
	// }
	encryptStr := "7d45706d29ec9e609c3b623ad99d2b70d18a9814---1670835247000"
	encrtpted, err := pub.EncryptAsn1([]byte(encryptStr), rand.Reader)
	if err != nil {
		fmt.Println("2", err)
		return "", err
	}
	return string(encrtpted), nil
}

type LoginReq struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	UserAlias string `json:"userAlias"`
	Encrypt   string `json:"encrypt"`
}

func HttpPostJson() (err error) {
	loginReq := &LoginReq{
		AppKey:    "iiVZEtM1",
		UserAlias: "Admin",
		Encrypt:   "sm2",
	}
	appSecret, err := getSecret()
	if err != nil {
		return err
	}
	loginReq.AppSecret = appSecret
	jsonStr, _ := json.Marshal(loginReq)

	reqUrl := "http://10.62.6.221:18241/douc/api/v1/token/getToken"

	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", apiToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	body, _ := ioutil.ReadAll(response.Body)
	// err = json.Unmarshal(body, resp)
	fmt.Println("response: ", string(body))
	return
}
