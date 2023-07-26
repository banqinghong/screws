package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/banqinghong/screws/golang/huawei"
	"github.com/tjfoc/gmsm/x509"
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
	// huawei.ListSslCert()
	//ali.DescribeChangeLogs()
	//tencent.DescribeDomainRecords()
	//aws.ListInstances()
	// aws.QueryELB2TargetGroup()
	// aws.QueryMqBrokers()
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
	// prom.QueryCpuUsage()
	// prom.QueryCpuUsage()
	// stringSplit()
	// fmt.Println("hash: ", Hash("salary-backend--ce8u44rlv2468mu4htcg"))
	// rdsAddressSplit()
	// getLastMonth("2023-04")
	// huawei.ListAllApis()
	// apollo.GetApolloConfig()
	// huawei.UploadFileToObs("service_cost/2023-04/xxxxxx.xlsx", "./dist/1111111.xlsx")
	// err := newDownloadDir("./download")
	// if err != nil {
	// 	fmt.Printf("create dir %s\n", err)
	// }
	// betcd.GetKVWithPrefix("jarvis.jobcenter/cronjob/k8s-devops-hw-qa")

	huawei.ShowDomainStats()
	// btime.GetNearestTime(1 * time.Minute)

	fmt.Println("main ending")
}

func newDownloadDir(DownloadDir string) error {
	// baseDir, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }
	// path := filepath.Join(baseDir, DownloadDir)
	_, err := os.Stat(DownloadDir)
	if err != nil {
		fmt.Println("new-------------")
		return os.Mkdir(DownloadDir, 0755)
	}
	return nil
}

// 计算hash值
func Hash(s string) string {
	hs := sha256.New()
	hs.Write([]byte(s))
	bs := fmt.Sprintf("%x", hs.Sum(nil))
	return bs
}

var (
	redisProxyAddressMap = map[string]string{
		"redis-hw.dev.xgimi.com": "redis-523db391-be9d-48fb-9fb6-7c9a26f9abe8.cn-east-3.dcs.myhuaweicloud.com",
		"redis-hw.qa.xgimi.com":  "redis-0c000086-8e51-495c-bf3f-7ee404d20bc1.cn-east-3.dcs.myhuaweicloud.com",
	}
)

func stringSplit() {
	tracer := "nlu-ir-services--cblp7hblv2422tn1o1mg-7c7c8f9d48-dqxvn"
	comma := strings.LastIndex(tracer, "-")
	if comma == -1 {
		fmt.Println("not found")
		return
	}
	fmt.Println(tracer[:comma])
}

func getLastMonth(billDate string) {
	firstDayTimeStr := fmt.Sprintf("%s-01 00:00:00", billDate)
	firstDay, _ := time.ParseInLocation("2006-01-02 15:04:05", firstDayTimeStr, time.Local)
	lastMonth := firstDay.AddDate(0, -1, 0)
	fmt.Println("time: ", lastMonth.Format("2006-01"))
}

func getStartAndEndTimeFromBillDate(billDate string) (time.Time, time.Time) {
	now := time.Now()
	firstDayTimeStr := fmt.Sprintf("%s-01 00:00:00", billDate)
	firstDay, _ := time.ParseInLocation("2006-01-02 15:04:05", firstDayTimeStr, time.Local)
	lastDay := firstDay.AddDate(0, 1, -1)
	if now.Unix() < lastDay.Unix() {
		lastDay = now
	}
	return firstDay, lastDay
}
func rdsAddressSplit() {
	value := "redis://:dfdgSGGSSGJS8s8shs8hs@redis-hw.dev.xgimi.com:6379/50"
	address := ""
	valueList := strings.Split(value, ":")
	if len(valueList) < 2 {
		return
	}
	// 如果字符串中存在 @符号  dghdh*Sbs7svsys67svs6vbsvvc@r-bp16d5513cfbc1f4.redis.rds.aliyuncs.com
	// 没有 //10.40.33.182
	if strings.Contains(value, "@") {
		valueList2 := strings.Split(valueList[2], "@")
		if len(valueList2) < 2 {
			return
		}
		address = valueList2[1]
	} else {
		valueList2 := strings.Split(valueList[1], "//")
		if len(valueList2) < 2 {
			return
		}
		address = valueList2[1]
	}

	if k, ok := redisProxyAddressMap[address]; ok {
		address = k
	}

	fmt.Println("addr: ", address)
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
