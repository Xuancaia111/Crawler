package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"crawler/persist"
)

func main() {
	//e:=engine.SimpleEngine{}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		Itemchan: persist.ItemSaver(),
	}
	/*e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
