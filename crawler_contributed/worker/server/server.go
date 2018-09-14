package main

import (
	"crawler/crawler_contributed/rpcsupport"
	"crawler/crawler_contributed/worker"
	"fmt"
	"log"
	"flag"
)

var port = flag.Int("port", 0, "the port for me to listen")

func main() {
	flag.Parse()
	if *port==0{
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		&worker.CrawlService{}))
}
