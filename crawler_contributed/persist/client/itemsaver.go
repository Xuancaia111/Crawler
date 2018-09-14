package client

import (
	"crawler/engine"
	"log"
	"crawler/crawler_contributed/rpcsupport"
	"crawler/crawler_contributed/config"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err!=nil{
		return nil,err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out

			itemCount++
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)

			result:=""
			err:=client.Call(config.ItemSaverRpc,item,&result)
			if err != nil || result!="ok"{
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			} else {
				log.Printf("Item Saver: Save successfully")
			}
		}

	}()
	return out, nil
}
