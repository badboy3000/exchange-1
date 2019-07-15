package kline

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/FlowerWrong/exchange/db"
)

// Key ...
func Key(symbol string, period int64) string {
	return "exchange:" + symbol + ":k:" + strconv.FormatInt(period, 10)
}

// LastTimestamp ...
func LastTimestamp(symbol string, period int64) (int64, error) {
	// -1 表示列表的最后一个元素
	dataJSON, err := db.Redis().LIndex(Key(symbol, period), -1).Result()
	if err != nil {
		return 0, err
	}

	var arr []string
	err = json.Unmarshal([]byte(dataJSON), &arr)
	if err != nil {
		return 0, err
	}
	ts, err := strconv.ParseInt(arr[0], 10, 64)
	if err != nil {
		return 0, err
	}
	return ts, nil
}

// NextTimestamp ...
func NextTimestamp(symbol string, period int64) (int64, error) {
	ts, err := LastTimestamp(symbol, period)
	if err != nil {
		// TODO
		return 0, err
	}
	tm := time.Unix(ts, 0)
	tm.Add(time.Duration(period) * time.Second)
	return tm.Unix(), err
}
