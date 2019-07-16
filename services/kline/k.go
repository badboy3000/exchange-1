package kline

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/FlowerWrong/exchange/db"
	"github.com/FlowerWrong/exchange/dtos"
	"github.com/FlowerWrong/exchange/models"
	"github.com/FlowerWrong/exchange/utils"
	"github.com/shopspring/decimal"
)

// Key ...
func Key(symbol string, period int64) string {
	return "exchange:" + symbol + ":k:" + strconv.FormatInt(period, 10)
}

func getTimestamp(symbol string, period, index int64) (int64, error) {
	dataJSON, err := db.Redis().LIndex(Key(symbol, period), index).Result()
	if err != nil {
		return 0, err
	}

	var klineDTO dtos.KlineDTO
	err = json.Unmarshal([]byte(dataJSON), &klineDTO)
	if err != nil {
		return 0, err
	}
	ts, err := utils.Str2Int64(klineDTO.Time)
	if err != nil {
		return 0, err
	}
	return ts, nil
}

// LastTimestamp ...
func LastTimestamp(symbol string, period int64) (int64, error) {
	// -1 表示列表的最后一个元素
	return getTimestamp(symbol, period, -1)
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

// K1 ...
func K1(symbol string, start time.Time) *dtos.KlineDTO {
	var orders []models.Order
	db.ORM().Where("created_at >= ? AND created_at < ?", start, start.Add(1*time.Minute)).Find(&orders)
	if len(orders) == 0 {
		return nil
	}
	totalVolume := decimal.NewFromFloat(0)
	min := orders[0].Price
	max := orders[0].Price
	for _, order := range orders {
		totalVolume = totalVolume.Add(order.Volume)
		if order.Price.Sub(max).Sign() > 0 {
			max = order.Price
		}
		if order.Price.Sub(min).Sign() < 0 {
			min = order.Price
		}
	}
	var klineDTO dtos.KlineDTO
	klineDTO.Time = utils.Int642Str(start.Unix())
	klineDTO.Open = orders[0].Price
	klineDTO.Close = orders[len(orders)-1].Price
	klineDTO.Volume = totalVolume
	klineDTO.High = max
	klineDTO.Low = min
	return &klineDTO
}

func k1Set(symbol string, start time.Time, period int64) ([]dtos.KlineDTO, error) {
	var klineDTOs []dtos.KlineDTO
	// 以 0 表示列表的第一个元素
	ts, err := getTimestamp(symbol, period, 0)
	if err != nil {
		return klineDTOs, err
	}
	offset := (start.Unix() - ts) / 60
	left := offset
	if left < 0 {
		left = 0
	}
	right := offset + period - 1
	if right < 0 {
		return klineDTOs, nil
	}
	dataJSON, err := db.Redis().LRange(Key(symbol, 1), left, right).Result()
	if err != nil {
		return klineDTOs, err
	}
	err = json.Unmarshal(utils.StrSlice2ByteSlice(dataJSON), &klineDTOs)
	if err != nil {
		return klineDTOs, err
	}
	return klineDTOs, nil
}

// Kn ...
func Kn(symbol string, start time.Time, period int64) *dtos.KlineDTO {
	klineDTOs, err := k1Set(symbol, start, period)
	if err != nil {
		return nil
	}

	totalVolume := decimal.NewFromFloat(0)
	min := klineDTOs[0].Low
	max := klineDTOs[0].High
	for _, k := range klineDTOs {
		totalVolume = totalVolume.Add(k.Volume)
		if k.High.Sub(max).Sign() > 0 {
			max = k.High
		}
		if k.Low.Sub(min).Sign() < 0 {
			min = k.Low
		}
	}
	var klineDTO dtos.KlineDTO
	klineDTO.Time = utils.Int642Str(start.Unix())
	klineDTO.Open = klineDTOs[0].Open
	klineDTO.Close = klineDTOs[len(klineDTOs)].Close
	klineDTO.Volume = totalVolume
	klineDTO.High = max
	klineDTO.Low = min
	return &klineDTO
}
