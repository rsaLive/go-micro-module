package utils

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func Post(url string, jsonStr []byte) (statusCode int, result string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("post url:", url)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(body))
	zap.L().Info("post", zap.Any("url", url), zap.Any("jsonStr", jsonStr), zap.Any("result", result))
	result = string(body)
	return
}

func Get(url string) (statusCode int, result string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	fmt.Println("response StatusCode:", resp.StatusCode)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	result = string(body)
	fmt.Println("response Body:", string(body))
	zap.L().Info("Get", zap.Any("url", url), zap.Any("result", result))
	return
}
