package parser

import (
	"crawler/engine"
	"regexp"
	"strconv"
	"crawler/module"
	"crawler_contributed/config"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe=regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe=regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var incomeRe=regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var genderRe=regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe=regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var educationRe=regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe=regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouRe=regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe=regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe=regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var idUrlRe=regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
//var nameRe=regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)
var guessRe=regexp.MustCompile(`href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := module.Profile{}

	if age, err := strconv.Atoi(extractString(contents,ageRe)); err == nil{
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractString(contents,heightRe)); err == nil{
		profile.Height = height
	}
	if weight, err := strconv.Atoi(extractString(contents,weightRe)); err == nil{
		profile.Weight = weight
	}

	profile.Marriage = extractString(contents,marriageRe)
	profile.Income = extractString(contents,incomeRe)
	profile.Gender = extractString(contents,genderRe)
	profile.Xinzuo = extractString(contents,xinzuoRe)
	profile.Education = extractString(contents,educationRe)
	profile.Occupation = extractString(contents,occupationRe)
	profile.Hukou = extractString(contents,hukouRe)
	profile.House = extractString(contents,houseRe)
	profile.Car = extractString(contents,carRe)
	profile.Name = name

	result:=engine.ParseResult{
		Items: []engine.Item{
			{
				Url: url,
				Type:"zhenai",
				Id: extractString([]byte(url),idUrlRe),
				Payload:profile,
			},
		},
	}
	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: NewProfileParser(string(m[2])),
			})
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{userName: name}
}