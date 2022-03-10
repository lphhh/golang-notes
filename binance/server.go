package main

import (
	"binance/database"
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"time"
)

var (
	apiKey    = "kuiExgYHNv78h5C2xVixs6BqcAu44O175VGrBYq14pfQcPoETauIVi6JLn8ebuAR"
	secretKey = "offKtWtkuFgBGnAfY0fATwPAzoyC7bmNSHx9g8KEnoT4WoCUIHlOF040KYikbmC8"
)

var client = binance.NewClient(apiKey, secretKey)
var futuresClient = binance.NewFuturesClient(apiKey, secretKey)   // USDT-M Futures
var deliveryClient = binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures

func ListPrices() {
	prices, err := futuresClient.NewListPricesService().Symbol("DYDXUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}
}

func SocketF() {
	fmt.Println("start at" + time.Now().String())
	now := time.Now()
	database.DB.Save(&Log{
		Msg:  "start",
		Time: &now,
	})
	wsDepthHandler := func(event futures.WsAllMarkPriceEvent) {
		fMarkPrice <- event
	}
	errHandler := func(err error) {
		fmt.Println(err)
		now := time.Now()
		database.DB.Save(&Log{
			Msg:  err.Error(),
			Time: &now,
		})
		go func() {
			quit <- true
		}()
		return
	}
	doneC, stopC, err := futures.WsAllMarkPriceServeWithRate(1*time.Second, wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		now := time.Now()
		database.DB.Save(&Log{
			Msg:  err.Error(),
			Time: &now,
		})
		go func() {
			quit <- true
		}()
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(10 * time.Second)
		quit <- true
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC
}

func SocketS() {
	wsDepthHandler := func(event binance.WsAllMarketsStatEvent) {
		sMarkPrice <- event
	}
	errHandler := func(err error) {
		fmt.Println(err)
		now := time.Now()
		database.DB.Save(&Log{
			Msg:  err.Error(),
			Time: &now,
		})
		go func() {
			quit <- true
		}()
		return
	}
	doneC, _, err := binance.WsAllMarketsStatServe(wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		now := time.Now()
		database.DB.Save(&Log{
			Msg:  err.Error(),
			Time: &now,
		})
		go func() {
			quit <- true
		}()
		return
	}
	// use stopC to exit
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	stopC <- struct{}{}
	//}()
	// remove this if you do not want to be blocked here
	<-doneC
}
