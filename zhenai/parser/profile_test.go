package parser

import (
	"crawler/module"
	"io/ioutil"
	"testing"
	"crawler/engine"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile("http://album.zhenai.com/u/107792312", contents, "推开窗子看见你")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	actual := result.Items[0]

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

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
