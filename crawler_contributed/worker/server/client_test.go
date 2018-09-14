package main

import (
	"testing"
	"crawler/crawler_contributed/rpcsupport"
	"crawler/crawler_contributed/worker"
	"time"
	"crawler/crawler_contributed/config"
	"fmt"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"

	go rpcsupport.ServeRpc(host, &worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err!=nil{
		panic(err)
	}
	request:=worker.Request{
		Url:"http://album.zhenai.com/u/107792312",
		Parser:worker.SerializedParser{
			Name:config.ParseProfile,
			Args:"推开窗子看见你",
		},
	}
	result:=worker.ParseResult{}
	err = client.Call(config.CrawlServiceRpc, request, &result)

	if err!=nil {
		t.Error(err)
	}else{
		fmt.Println(result)
	}
}
