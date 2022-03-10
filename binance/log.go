package main

import "time"

type Log struct {
	Id   int64      `json:"id"`
	Msg  string     `json:"msg"`
	Time *time.Time `json:"time"`
}
