package bhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseHost = "http://api.example.com"

func HttpPostJson(path string, content string, resp interface{}) (err error) {
	jsonStr := []byte(content)
	reqUrl := baseHost + path
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
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, resp)
	fmt.Println("response: ", string(body))
	return
}

func HttpPutJson(path string, content string, resp interface{}) (err error) {
	jsonStr := []byte(content)
	reqUrl := baseHost + path

	req, err := http.NewRequest("PUT", reqUrl, bytes.NewBuffer(jsonStr))
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
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, resp)
	fmt.Println("response: ", string(body))
	return
}

func HttpDelete(path string, resp interface{}) (err error) {
	reqUrl := baseHost + path
	req, err := http.NewRequest("DELETE", reqUrl, nil)
	if err != nil {
		return err
	}
	//req.Header.Set("Authorization", apiToken)

	response, _ := http.DefaultClient.Do(req)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response: ", string(body))
	if string(body) == "" {
		return nil
	}
	err = json.Unmarshal(body, resp)
	return
}

func HttpGet(path string, resp interface{}) (err error) {
	reqUrl := baseHost + path
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	//req.Header.Set("Authorization", apiToken)

	response, _ := http.DefaultClient.Do(req)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response: ", string(body))
	err = json.Unmarshal(body, resp)
	return
}
