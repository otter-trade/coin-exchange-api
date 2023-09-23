# coin-exchange-api
## 参考1. binance

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

 	pair := currency.NewPair(currency.BTC, currency.USDT)
	startTime := time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2021, 2, 15, 0, 0, 0, 0, time.UTC)

	resp1, err := b.GetHistoricCandles(context.Background(), pair, asset.Spot, kline.Interval(time.Hour*5), startTime, endTime)
	if !errors.Is(err, kline.ErrRequestExceedsExchangeLimits) {
		fmt.Println("received: '%v', but expected: '%v'", err, kline.ErrRequestExceedsExchangeLimits)
	} else {
		fmt.Println("open3: ", resp1)
	}

	resp1, err = b.GetHistoricCandles(context.Background(), pair, asset.Spot, kline.OneDay, startTime, endTime)
	if err != nil {
		fmt.Println("Binanceus GetHistoricCandles() error", err)
	} else {
		fmt.Println("open4: ", resp1)
	}
}

## 参考2. okx


func main() {

	cfg := config.GetConfig()
 
	err := cfg.LoadConfig("./testdata/configtest.json", true)
 
	if err != nil {
		log.Fatal("Binance load config error", err)
	}

	okxConfig, err := cfg.GetExchangeConfig("Okx")
	if err != nil {
		log.Fatal("Okx Setup() init error", err)
	}
	okxConfig.API.AuthenticatedSupport = true
	okxConfig.API.Credentials.Key = "aab214a7-2326-4cb8-91ca-f6372a6e32c6"
	okxConfig.API.Credentials.Secret = "xxxxxxxxxxxxxxxxxxxxx"
	okxConfig.API.Credentials.ClientID = "passphrase"

	//
	// Use unified api to create the exchange
	fmt.Println("Use unified api to create the exchange")
	em := engine.NewExchangeManager()
	exch, err := em.NewExchangeByName("Okx")
	if err != nil {
		fmt.Println("NewExchangeByName() error", err)
	}
	exch.SetDefaults()

	err = exch.Setup(okxConfig)
	if err != nil {
		fmt.Println("exch.Setup() error", err)
	}
	err = exch.UpdateTradablePairs(context.Background(), true)
	if err != nil {
		log.Fatal("okx setup error", err)
	}

	pair := currency.NewPair(currency.BTC, currency.USDT)
	fmt.Println("pair: ", pair)
	startTime := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC)

	//resp1, err := exch.GetHistoricCandles(context.Background(), pair, asset.Spot, kline.Interval(time.Hour), startTime, endTime)
	resp1, err := exch.GetHistoricCandles(context.Background(), pair, asset.Spot, kline.OneDay, startTime, endTime)
	if !errors.Is(err, kline.ErrRequestExceedsExchangeLimits) {
		fmt.Println("received: '%v', but expected: '%v'", err, kline.ErrRequestExceedsExchangeLimits)
	} else {
		fmt.Println("resp1: ", resp1)
	}
}


