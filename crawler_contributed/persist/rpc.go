package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"crawler/engine"
	"crawler/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item,result *string) error{
	err := persist.Save(s.Client, s.Index, item)
	if err==nil{
		*result="ok"
		log.Printf("Item %v saved.",item)
	} else {
		log.Printf("Error saving item %v: %v",item,err)
	}
	return err
}
