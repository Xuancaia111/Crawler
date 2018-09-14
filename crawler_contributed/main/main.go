package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	itemsaver "crawler/crawler_contributed/persist/client"
	"crawler/crawler_contributed/config"
	worker "crawler/crawler_contributed/worker/client"
	"net/rpc"
	"crawler/crawler_contributed/rpcsupport"
	"log"
	"flag"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host","","itemsaver host")

	workerHosts=flag.String("worker_hosts","","worker hosts (comma separated)")
)
func main() {
	flag.Parse()

	//e:=engine.SimpleEngine{}
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	hosts := strings.Split(*workerHosts, ",")
	pool:=createClientPool(hosts)
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: worker.CreateProcessor(pool),
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
	/*e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})*/
}

func createClientPool(hosts[]string) chan *rpc.Client{
	var clients []*rpc.Client
	for _,host:=range hosts{
		client, err := rpcsupport.NewClient(host)
		if err!=nil{
			log.Printf("Error connecting to %s: %v",host,err)
		}else{
			clients = append(clients, client)
			log.Printf("Connected to %s",host)
		}
	}

	out:=make(chan *rpc.Client)
	go func() {
		for{
			for _,c:=range clients{
				out<-c
			}
		}
	}()
	return out
}