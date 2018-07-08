package persist

import (
	"testing"
	"crawler/module"
	"github.com/olivere/elastic"
	"context"
	"encoding/json"
	"crawler/engine"
)

func TestSave(t *testing.T) {
	expected :=engine.Item{
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

	const index = "dating_test"
	//Save expected

	//TODO: Try to start up elastic search here using docker go client.
	client,err:=elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err!=nil{
		panic(err)
	}
	err = save(client,index,expected)
	if err!=nil{
		panic(err)
	}
	//Fetch saved item
	resp, err := client.Get().
		Index("dating_test").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err!=nil{
		panic(err)
	}

	t.Logf("%s",resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)

	actualProfile,err:=module.FromJsonObj(actual.Payload)
	actual.Payload=actualProfile
	//Verify result
	if actual!=expected {
		t.Errorf("expected: %v, but was %v",expected,actual)
	}
}