package apollo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetApolloConfig() {
	reqUrl := fmt.Sprintf("%s/apps/%s/envs/%s/clusters/%s/namespaces", ApolloHost, "xgimi-bff-overseas", "PRO", "default")

	request, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		fmt.Printf("new request failed: %s\n", err)
		return
	}

	cookies, _ := getCookie(ApolloHost, ApolloUser, ApolloPassword)
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("do request url[%s] failed: %s\n", reqUrl, err)
		return
	}

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("body: ", string(body))
	var configSet []*ApplicationConfig
	err = json.Unmarshal(body, &configSet)
	if err != nil {
		fmt.Printf("parse apollo config failed: %s\n", err)
		return
	}
	defer response.Body.Close()

	jsonStr, _ := json.Marshal(configSet)
	fmt.Println("result: ", string(jsonStr))
}

func getCookie(host, user, passwd string) ([]*http.Cookie, error) {
	req, err := http.NewRequest("GET", host, nil)
	if err != nil {
		fmt.Printf("new request failed: %s\n", err)
		return nil, err
	}

	req.SetBasicAuth(user, passwd)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http request failed: %s\n", err)
		return nil, err
	}

	defer response.Body.Close()
	return response.Cookies(), nil
}
