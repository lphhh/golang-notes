package main

import (
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
	wsDepthHandler := func(event futures.WsAllMarkPriceEvent) {
		for _,v := range event {
			fmt.Println(v)
		}
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := futures.WsAllMarkPriceServe(wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC
}


func SocketS() {
	wsDepthHandler := func(event binance.WsAllMarketsStatEvent) {
		for _,v := range event {
			fmt.Println(v)
		}
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := binance.WsAllMarketsStatServe(wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC
}