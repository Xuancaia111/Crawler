package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic"
	"crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver(index string) (chan engine.Item, error) {

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil,err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out

			itemCount++
			log.Printf("Got item #%d: %v", itemCount, item)

			err :=save(client,index,item)
			if err!=nil{
				log.Printf("Item Saver: error saving item %v: %v",item,err)
			}else{
				log.Printf("Save successfully")
			}
		}

	}()
	return out,nil
}

func save(client *elastic.Client, index string,item engine.Item) error{

	if item.Type== "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id!=""{
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	if err!=nil{
		return err
	}

	return nil
}
