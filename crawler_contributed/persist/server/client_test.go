package main

import (
	"testing"
	"crawler/crawler_contributed/rpcsupport"
	"time"
	"crawler/engine"
	"crawler/module"
	"crawler/crawler_contributed/config"
)

func TestItemSaver(t *testing.T){
	const host=":1234"
	go serveRpc(host,"test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err!=nil{
		panic(err)
	}

	item :=engine.Item{
		Url:"http://album.zhenai.com/u/107792312",
		Type:"zhenai",
		Id:"107792312",
		Payload:module.Profile{
			Name:       "推开窗子看见你",
			Age:        27,
			Height:     170,
			Income:     "3001-5000元",
			Gender:     "女",
			Xinzuo:     "金牛座",
			Occupation: "小学教师",
			Marriage:   "未婚",
			House:      "和家人同住",
			Hukou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	result:=""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err!=nil||result!="ok"{
		t.Errorf("result: %s; err: %s",result,err)
	}
}