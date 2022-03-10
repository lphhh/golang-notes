package main

import "time"

type Future struct {
	Id                   int64     `json:"id"`
	Event                string    `json:"event"`
	Time                 *time.Time `json:"time"`
	Symbol               string    `json:"symbol"`
	MarkPrice            string    `json:"mark_price"`
	IndexPrice           string    `json:"index_price"`
	EstimatedSettlePrice string    `json:"estimated_settle_price"`
	FundingRate          string    `json:"funding_rate"`
	NextFundingTime      *time.Time `json:"next_funding_time"`
}
