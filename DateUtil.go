package common

import (
	"sync"
	"time"
)

type dateUtil struct {
	mutex sync.Mutex
}

var Date = dateUtil{}

func (this dateUtil) TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func (this dateUtil) CurrentTimeString() string {
	return this.TimeToString(time.Now())
}
