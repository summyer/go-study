package mmm

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HttpFanYi() {
	var dataBytes, err = ioutil.ReadFile("/Users/summyer/go-workspace/study/mmm/test_fanyi.json")
	if err != nil {
		log.Fatal(err)
	}
	//类似java中的fastjson取某个key ,可以多测试看看数据结构
	dataMap := make(map[string]interface{})
	//dataMap1 := make(map[string][]MetadataOp)
	//dataMap2 := make(map[string][]map[string]string)
	err = json.Unmarshal(dataBytes, &dataMap)
	if err != nil {
		log.Fatal(err)
	}
	records := dataMap["RECORDS"]
	l := list.New()
	if r, ok := records.([]interface{}); ok {
		for _, item := range r {
			if r, ok := item.(map[string]interface{}); ok {
				//fmt.Println(r["cve_id"])
				//interface{} ->转换成string
				l.PushBack(fmt.Sprintf("%v", r["code"]))
			}

		}
	}
	fmt.Println(l.Len())
	i := 1
	for item := l.Front(); item != nil; item = item.Next() {
		OnceReqWithCookie(item.Value.(string), i)
		i++
	}

	//mo := MetadataOp{MetadataId: "CVE-2017-20092", MetadataType: "CVE", OperateType: "UPDATE"}
	//OnceReq(mo)

}

func OnceReq(url string, i int) {
	// 创建请求
	resp, err := http.Get("http://10.20.152.61:9000/api/translate/" + url) // url
	if err != nil {
		log.Printf("创建请求失败！err:%+v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取Body失败 error: %+v", err)
		return
	}
	log.Println(string(body))
}

func OnceReqWithCookie(code string, i int) {
	url := "https://mmm.dbappsecurity.com.cn/api/translate/"
	url = "http://10.20.152.192:9000/api/translate/"
	log.Println("开始翻译:", code, ",第几条:", i)
	// 创建客户端
	client := &http.Client{}
	// 创建请求
	request, err := http.NewRequest("GET", url+code, nil)
	if err != nil {
		log.Printf("创建请求失败！err:%+v", err)
		return
	}
	// 设置请求头
	request.Header.Add("Cookie", "jsessionid-sc=MjhlYzQ0MjMtYWYxMC00ODJhLTgwYWUtNWZhOTAwNGJiY2I4; token=mmm_1659319566659247888")

	// 设置1分钟的超时时间
	client.Timeout = 1 * time.Minute
	// 发起请求
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取Body失败 error: %+v", err)
		return
	}
	log.Println(string(body))
}
