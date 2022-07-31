package main

import (
	"bytes"
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//var logger = log.Default()

type MetadataOp struct {
	MetadataId   string `json:"metadataId"`
	MetadataType string `json:"metadataType"`
	OperateType  string `json:"operateType"`
}

func main() {

	var dataBytes, err = ioutil.ReadFile("/Users/summyer/go-workspace/study/mmm_cve_update/test_cvss.json")
	if err != nil {
		log.Fatal(err)
	}
	//类似java中的fastjson取某个key ,可以多测试看看数据结构
	dataMap := make(map[string]interface{})
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
				mo := MetadataOp{MetadataId: fmt.Sprintf("%v", r["cve_id"]), MetadataType: "CVE", OperateType: "UPDATE"}
				l.PushBack(mo)
			}

		}
	}
	fmt.Println(l.Len())
	i := 1
	for item := l.Front(); item != nil; item = item.Next() {
		//fmt.Println(item.Value.(MetadataOp).MetadataId)
		OnceReq(item.Value.(MetadataOp), i)
		i++
	}

	//mo := MetadataOp{MetadataId: "CVE-2017-20092", MetadataType: "CVE", OperateType: "UPDATE"}
	//OnceReq(mo)

}

func OnceReq(mo MetadataOp, i int) {
	jsonBytes, err := json.Marshal(mo)
	if err != nil {
		log.Printf("序列化body失败！err:%+v", err)
		return
	}

	// 创建客户端
	client := &http.Client{}
	// 创建请求
	request, err := http.NewRequest("POST", "mmm://10.20.152.192:9000/api/metadata-inform", bytes.NewReader(jsonBytes))
	if err != nil {
		log.Printf("创建请求失败！err:%+v", err)
		return
	}
	// 设置请求头
	request.Header.Add("Cookie", "jsessionid-sc=Y2VhNjljNmUtN2RhMi00MmU1LTlhYzItMjAwOTljMWM4MjVh; token=mmm_1659079555676507583")
	request.Header.Add("Content-Type", "application/json;charset=utf-8")

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
	log.Println(string(body), mo.MetadataId)
}
