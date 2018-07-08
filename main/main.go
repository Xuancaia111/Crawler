package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/persist"
	"crawler/zhenai/parser"
)

func main() {
	//e:=engine.SimpleEngine{}
	itemChan,err:=persist.ItemSaver("dating_profile")
	if err!=nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		Itemchan: itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	/*e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})*/
}
