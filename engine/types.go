package engine

import "crawler/crawler_contributed/config"

type ParserFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}
type NilParser struct{}

func (f *NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (f *NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}


type FuncParser struct {
	Parser ParserFunc
	Name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.Parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.Name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		Parser: p,
		Name:   name,
	}
}
