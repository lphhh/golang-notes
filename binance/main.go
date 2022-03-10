package main

import (
	"binance/database"
	"fmt"
	_ "github.com/CodyGuo/godaemon"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"time"
)

var sMarkPrice = make(chan binance.WsAllMarketsStatEvent)
var fMarkPrice = make(chan futures.WsAllMarkPriceEvent)
var quit = make(chan bool)

func main() {
	go SocketF()
	//go SocketS()
	for {
		select {
		case event := <-fMarkPrice:
			for _, v := range event {
				if v.Symbol == "DYDXUSDT" {

					fmt.Println("future", v, time.UnixMilli(v.Time).Format("2006-01-02 15:04:05"))
					publishAt := time.Unix(v.Time/1000, 0)
					nextFundingTime := time.Unix(v.NextFundingTime/1000, 0)
					database.DB.Save(&Future{
						Event:                v.Event,
						Time:                 &publishAt,
						Symbol:               v.Symbol,
						MarkPrice:            v.MarkPrice,
						IndexPrice:           v.IndexPrice,
						EstimatedSettlePrice: v.EstimatedSettlePrice,
						FundingRate:          v.FundingRate,
						NextFundingTime:      &nextFundingTime,
					})
				}
			}
		case event := <-sMarkPrice:
			for _, v := range event {
				if v.Symbol == "DYDXUSDT" {
					fmt.Println("spot", time.UnixMilli(v.Time).Format("2006-01-02 15:04:05"))
				}
			}
		case <-quit:
			fmt.Println("accept quit signal", time.Now().Format("2006-01-02 15:04:05"))
			now := time.Now()
			database.DB.Save(&Log{
				Msg:  "accept quit signal",
				Time: &now,
			})
			go SocketF()
		}
	}
	//SocketS()
}
