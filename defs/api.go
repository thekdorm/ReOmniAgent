package defs

import "fmt"

const BaseUrl string = "https://us.api.blizzard.com"
const Locale string = "en_US"
const Namespace string = "dynamic-us"
const ProudmooreRealmId int = 5

var AuctionHouseUrl string = fmt.Sprintf("/data/wow/connected-realm/%d/auctions", ProudmooreRealmId)
var CommoditiesUrl string = "/data/wow/auctions/commodities"
var UrlQueries string = fmt.Sprintf("?namespace=%s&locale=%s", Namespace, Locale)

type AuctionHouseJson struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`

	ConnectedRealm struct {
		Href string `json:"href"`
	} `json:"connected_realm"`

	Auctions []struct {
		Id   int `json:"id"`
		Item struct {
			Id         int32 `json:"id"`
			Context    int32 `json:"context"`
			BonusLists []int32
			Modifiers  []struct {
				Type  int32 `json:"type"`
				Value int32 `json:"value"`
			} `json:"modifiers"`
		}
		Buyout   int64  `json:"buyout"`
		Quantity int64  `json:"quantity"`
		TimeLeft string `json:"time_left"`
	} `json:"auctions"`

	CommoditiesHref struct {
		Href string `json:"href"`
	} `json:"commodities"`
}

type CommodityJson struct {
	Id   int32 `json:"id"`
	Item struct {
		Id int32 `json:"id"`
	} `json:"item"`
	Quantity  int32  `json:"quantity"`
	UnitPrice int64  `json:"unit_price"`
	TimeLeft  string `json:"time_left"`

	Name string
}

type CommoditiesJson struct {
	Auctions []CommodityJson `json:"auctions"`
}

type Commodity struct {
	id int32

	Name string
}

type QualityCommodity struct {
	Tier1 Commodity
	Tier2 Commodity
	Tier3 Commodity
}

func (commodity Commodity) Id() int32 {
	return commodity.id
}
