package client

import (
	"crawler/engine"
	"crawler/crawler_contributed/config"
	"crawler/crawler_contributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor{
	return func(request engine.Request)(engine.ParseResult,error){
		c:=<-clientChan
		req:=worker.SerializeRequest(request)
		result:=worker.ParseResult{}
		err := c.Call(config.CrawlServiceRpc, req, &result)
		if err!=nil{
			return engine.ParseResult{},err
		}
		return worker.DeserializeResult(result),nil
	}
}


