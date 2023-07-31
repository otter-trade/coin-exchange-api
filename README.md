# coin-exchange-api
参考1. binance

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/otter-trade/coin-exchange-api/config"
	"github.com/otter-trade/coin-exchange-api/exchanges/binance"
	"github.com/otter-trade/coin-exchange-api/exchanges/request"
	"github.com/otter-trade/coin-exchange-api/exchanges/sharedtestvalues"
)

// getTime returns a static time for mocking endpoints, if mock is not enabled
// this will default to time now with a window size of 30 days.
// Mock details are unix seconds; start = 1577836800 and end = 1580515200
func getTime() (start, end time.Time) {

	tn := time.Now()
	offset := time.Hour * 24 * 30
	return tn.Add(-offset), tn
}

const (
	apiKey                  = ""
	apiSecret               = ""
	canManipulateRealOrders = false
	BitcoinDonationAddress  = "bc1qk0jareu4jytc0cfrhr5wgshsq8282awpavfahc"
)

var b = &binance.Binance{}

func main() {
	cfg := config.GetConfig()
	err := cfg.LoadConfig("./testdata/configtest.json", true)
	if err != nil {
		log.Fatal("Binance load config error", err)
	}
	binanceConfig, err := cfg.GetExchangeConfig("Binance")
	if err != nil {
		log.Fatal("Binance Setup() init error", err)
	}
	binanceConfig.API.AuthenticatedSupport = true
	binanceConfig.API.Credentials.Key = apiKey
	binanceConfig.API.Credentials.Secret = apiSecret
	b.SetDefaults()
	b.Websocket = sharedtestvalues.NewTestWebsocket()
	err = b.Setup(binanceConfig)
	if err != nil {
		log.Fatal("Binance setup error", err)
	}
	//b.setupOrderbookManager()
	request.MaxRequestJobs = 100
	b.Websocket.DataHandler = sharedtestvalues.GetWebsocketInterfaceChannelOverride()
	log.Printf(sharedtestvalues.LiveTesting, b.Name)
	err = b.UpdateTradablePairs(context.Background(), true)
	if err != nil {
		log.Fatal("Binance setup error", err)
	}
	resp, err := b.GetIndexPriceKlines(context.Background(), "BTCUSD", "1M", 5, time.Time{}, time.Time{})
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("open1: ", resp[0].Open)
	}
	start, end := getTime()
	_, err = b.GetIndexPriceKlines(context.Background(), "BTCUSD", "1M", 5, start, end)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("open2: ", resp[0].Open)
	}
}

