package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HttpPost(url string, header map[string]string, req interface{}, rsp interface{}) error {
	client := &http.Client{Timeout: time.Second * 10}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}

	response, err := client.Do(request)
	defer response.Body.Close()

	if err != nil {
		return err
	}

	err = json.NewDecoder(response.Body).Decode(&rsp)
	return err
}

func DingDingMsg(token, title, content string) error {
	if token == "" {
		return nil
	}
	uri := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", token)

	type Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}

	type At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	}

	var req = struct {
		MsgType  string   `json:"msgtype"`
		Markdown Markdown `json:"markdown"`
		At       At       `json:"at"`
	}{
		MsgType: "markdown",
		Markdown: Markdown{
			Title: title,
			Text:  content,
		},
	}

	var rsp = struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}{}

	err := HttpPost(uri, nil, req, &rsp)
	if err != nil {
		return err
	}

	return nil
}
