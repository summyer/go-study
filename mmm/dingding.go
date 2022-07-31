package mmm

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var accessToken = "61a43fb5f3d87fce0356b3f97128e165b4d8e9cd2af61bc9fd2fd96dbe491a29"

func buildText(msg string) string {
	content := `{
					"msgtype": "text",
      				"text": {"content": "` + msg + `"},
                	"at": {
						 "atMobiles": [
							 "18307202679"
						 ],
						 "isAtAll": false
					}
				}`
	return content
}

func SendDingMsg(msg string) {
	webHook := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", accessToken)
	content := buildText(msg)
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println("handle error")
	}
}
