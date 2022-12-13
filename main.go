package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"reomniagent/defs"
	"reomniagent/utils"
	"time"
)

var debug = flag.Bool("debug", false, "_bool_\t"+"Enable debug functions")
var full = flag.Bool("full", false, "_bool_\t"+"Full program execution (Oauth2 + live API)")

const filenameCommodities string = "examples/commodities.txt"

func fullExec() {
	codeChan := make(chan string)

	go server(codeChan)
	client := utils.OauthToken(codeChan)
	client.Timeout = 60 * time.Second // Auction House results take a while to come back
	close(codeChan)

	reqUrl := defs.BaseUrl + defs.CommoditiesUrl + defs.UrlQueries
	rsp, err := client.Get(reqUrl)
	utils.Check(err, "No response received!")

	body, err := io.ReadAll(rsp.Body)
	utils.Check(err, "No response body could be read!")
	rsp.Body.Close()

	if *debug {
		utils.WritePrettyRspToFile(filenameCommodities, body)
	}

	auctions := defs.AuctionHouseJson{}
	err = json.Unmarshal([]byte(body), &auctions)
	utils.Check(err, "Couldn't Unmarshal JSON response body!")
}

func develExec() {
	commoditiesFile, err := os.Open(filenameCommodities)
	utils.Check(err, fmt.Sprintf("Couldn't open '%s'!", filenameCommodities))

	commoditiesContents, err := io.ReadAll(commoditiesFile)
	utils.Check(err, fmt.Sprintf("Couldn't read from opened file '%s'!", filenameCommodities))
	commoditiesFile.Close()

	commodities := defs.CommoditiesJson{}
	err = json.Unmarshal([]byte(commoditiesContents), &commodities)
	utils.Check(err, "Couldn't Unmarshal JSON file contents!")

	fmt.Println("Number of Auctions: " + fmt.Sprint(len(commodities.Auctions)))

	for i := 0; i < len(defs.Essences); i++ {
		cheapest := defs.CommodityJson{UnitPrice: -1}
		id := defs.Essences[i].Id()

		for j := 0; j < len(commodities.Auctions); j++ {
			auctionPtr := &commodities.Auctions[j]

			if (auctionPtr.Item.Id == id && cheapest.UnitPrice == -1) ||
				(auctionPtr.Item.Id == id && commodities.Auctions[j].UnitPrice < cheapest.UnitPrice) {
				if *debug {
					fmt.Printf("Commodity: %s\tOld Price: %v\tCheaper Price: %v\n", cheapest.Name, cheapest.UnitPrice, auctionPtr.UnitPrice)
				}
				cheapest = *auctionPtr
			}
		}

		fmt.Printf("%s cheapest: %v\n", defs.Essences[i].Name, cheapest.UnitPrice)
	}
}

func main() {
	flag.Parse()

	if *full {
		fullExec()
	} else {
		develExec()
	}
}
