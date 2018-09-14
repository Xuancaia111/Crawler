package config

const (
	//ElasticSearch
	ElasticIndex = "dating_profile"

	//RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
	//Parser names
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	//Rate limiting
	Qps = 20
)
