package matching

import (
	"container/list"
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

// OrderBook implements standard matching algorithm
type OrderBook struct {
	orders map[string]*list.Element // orderID -> *Order (*list.Element.Value.(*Order))

	asks *OrderSide // 卖单列表
	bids *OrderSide // 买单列表
}

// NewOrderBook creates Orderbook object
func NewOrderBook() *OrderBook {
	return &OrderBook{
		orders: map[string]*list.Element{},
		bids:   NewOrderSide(),
		asks:   NewOrderSide(),
	}
}

// ProcessMarketOrder immediately gets definite quantity from the order book with market price
// Arguments:
//      side     - what do you want to do (ob.Sell or ob.Buy)
//      quantity - how much quantity you want to sell or buy
//      * to create new decimal number you should use decimal.New() func
//        read more at https://github.com/shopspring/decimal
// Return:
//      error        - not nil if price is less or equal 0
//      done         - not nil if your market order produces ends of anoter orders, this order will add to
//                     the "done" slice
//      partial      - not nil if your order has done but top order is not fully done/你的订单全部成交了，但是委托单列表中该价格订单还剩余一部分
//      partialQuantityProcessed - if partial order is not nil this result contains processed quatity from partial order/已成交部分量
//      quantityLeft - more than zero if it is not enought orders to process all quantity/你的订单未全部成交，未成交部分量
func (ob *OrderBook) ProcessMarketOrder(side Side, quantity decimal.Decimal) (done []*Order, partial *Order, partialQuantityProcessed, quantityLeft decimal.Decimal, err error) {
	if quantity.Sign() <= 0 {
		return nil, nil, decimal.Zero, decimal.Zero, ErrInvalidQuantity
	}

	var (
		iter          func() *OrderQueue
		sideToProcess *OrderSide
	)

	// 如果是市价买单，找出卖单中的最低价订单列表
	if side == Buy {
		iter = ob.asks.MinPriceQueue
		sideToProcess = ob.asks
	} else {
		// 如果是市价卖单，找出买单中的最高价订单列表
		iter = ob.bids.MaxPriceQueue
		sideToProcess = ob.bids
	}

	// 未成交完且委单列表不为空
	for quantity.Sign() > 0 && sideToProcess.Len() > 0 {
		bestPrice := iter() // 最优价
		ordersDone, partialDone, partialProcessed, quantityLeft := ob.processQueue(bestPrice, quantity)
		done = append(done, ordersDone...)
		partial = partialDone
		partialQuantityProcessed = partialProcessed
		quantity = quantityLeft
	}

	quantityLeft = quantity
	return
}

// ProcessLimitOrder places new order to the OrderBook
// Arguments:
//      side     - what do you want to do (ob.Sell or ob.Buy)
//      orderID  - unique order ID in depth
//      quantity - how much quantity you want to sell or buy
//      price    - no more expensive (or cheaper) this price
//      * to create new decimal number you should use decimal.New() func
//        read more at https://github.com/shopspring/decimal
// Return:
//      error   - not nil if quantity (or price) is less or equal 0. Or if order with given ID is exists
//      done    - not nil if your order produces ends of anoter order, this order will add to
//                the "done" slice.
//      partial - not nil if your order has done but top order is not fully done. Or if your order is
//                partial done and placed to the orderbook without full quantity - partial will contain
//                your order with quantity to left
//                1. 如果你的订单全部成交了，但是委托单部分成交，那么partial就是成交后的剩余部分
//                2. 如果你的订单部分成交了，剩下的都放入了委托单列表，那么partial就是你订单的未成交部分
//      partialQuantityProcessed - if partial order is not nil this result contains processed quatity from partial order/已处理了的partial部分单量
func (ob *OrderBook) ProcessLimitOrder(side Side, orderID string, quantity, price decimal.Decimal) (done []*Order, partial *Order, partialQuantityProcessed decimal.Decimal, err error) {
	// 订单已经存在
	if _, ok := ob.orders[orderID]; ok {
		return nil, nil, decimal.Zero, ErrOrderExists
	}

	if quantity.Sign() <= 0 {
		return nil, nil, decimal.Zero, ErrInvalidQuantity
	}

	if price.Sign() <= 0 {
		return nil, nil, decimal.Zero, ErrInvalidPrice
	}

	quantityToTrade := quantity
	var (
		sideToProcess *OrderSide // 待撮合的委单列表
		sideToAdd     *OrderSide // 如果撮合不成功就添加进的委单列表
		comparator    func(decimal.Decimal) bool
		iter          func() *OrderQueue
	)

	if side == Buy {
		sideToAdd = ob.bids
		sideToProcess = ob.asks
		comparator = price.GreaterThanOrEqual // >=
		iter = ob.asks.MinPriceQueue
	} else {
		sideToAdd = ob.asks
		sideToProcess = ob.bids
		comparator = price.LessThanOrEqual // <=
		iter = ob.bids.MaxPriceQueue
	}

	bestPrice := iter()
	for quantityToTrade.Sign() > 0 && sideToProcess.Len() > 0 && comparator(bestPrice.Price()) {
		ordersDone, partialDone, partialQty, quantityLeft := ob.processQueue(bestPrice, quantityToTrade)
		done = append(done, ordersDone...)
		partial = partialDone
		partialQuantityProcessed = partialQty
		quantityToTrade = quantityLeft
		bestPrice = iter()
	}

	// 未成交 > 0，放入委托单
	if quantityToTrade.Sign() > 0 {
		// 未成交订单
		o := NewOrder(orderID, side, quantityToTrade, price, time.Now().UTC())

		// 部分成交
		if len(done) > 0 {
			partialQuantityProcessed = quantity.Sub(quantityToTrade) // 已成交量
			partial = o
		}
		ob.orders[orderID] = sideToAdd.Append(o)
	} else {
		// 完全成交
	}
	return
}

// 按照最优价格处理订单
func (ob *OrderBook) processQueue(bestPriceOrderQueue *OrderQueue, quantityToTrade decimal.Decimal) (done []*Order, partial *Order, partialQuantityProcessed, quantityLeft decimal.Decimal) {
	quantityLeft = quantityToTrade

	// 最优价订单列表不为空且未成交完
	for bestPriceOrderQueue.Len() > 0 && quantityLeft.Sign() > 0 {
		// 同一价格 时间优先
		headOrderEl := bestPriceOrderQueue.Head()
		headOrder := headOrderEl.Value.(*Order)

		// 未成交量 < 委单量
		if quantityLeft.LessThan(headOrder.Quantity()) {
			// 剩余委单
			partial = NewOrder(headOrder.ID(), headOrder.Side(), headOrder.Quantity().Sub(quantityLeft), headOrder.Price(), headOrder.Time())
			done = append(done, NewOrder(headOrder.ID(), headOrder.Side(), quantityLeft, headOrder.Price(), headOrder.Time())) // 成交单列表
			partialQuantityProcessed = quantityLeft                                                                            // 已成交部分量
			bestPriceOrderQueue.Update(headOrderEl, partial)
			quantityLeft = decimal.Zero // 未成交剩余单量
		} else {
			quantityLeft = quantityLeft.Sub(headOrder.Quantity())
			done = append(done, ob.CancelOrder(headOrder.ID())) // 成交单列表
		}
	}

	return
}

// Order returns order by id
func (ob *OrderBook) Order(orderID string) *Order {
	e, ok := ob.orders[orderID]
	if !ok {
		return nil
	}

	return e.Value.(*Order)
}

// Depth returns price levels and volume at price level
// asks: 110 -> 5
//       100 -> 1
// --------------
// bids: 90  -> 5
//       80  -> 1
func (ob *OrderBook) Depth() (asks, bids []*PriceLevel) {
	level := ob.asks.MaxPriceQueue()
	for level != nil {
		asks = append(asks, &PriceLevel{
			Price:    level.Price(),
			Quantity: level.Volume(),
		})
		level = ob.asks.LessThan(level.Price())
	}

	level = ob.bids.MaxPriceQueue()
	for level != nil {
		bids = append(bids, &PriceLevel{
			Price:    level.Price(),
			Quantity: level.Volume(),
		})
		level = ob.bids.LessThan(level.Price())
	}
	return
}

// CancelOrder removes order with given ID from the order book
func (ob *OrderBook) CancelOrder(orderID string) *Order {
	e, ok := ob.orders[orderID]
	if !ok {
		return nil
	}

	delete(ob.orders, orderID)

	if e.Value.(*Order).Side() == Buy {
		return ob.bids.Remove(e)
	}

	return ob.asks.Remove(e)
}

// CalculateMarketPrice returns total market price for requested quantity 返回成交总价 也就是各个成交价格乘以量，最后求和
// if err is not nil price returns total price of all levels in side
func (ob *OrderBook) CalculateMarketPrice(side Side, quantity decimal.Decimal) (price decimal.Decimal, err error) {
	price = decimal.Zero

	var (
		level *OrderQueue
		iter  func(decimal.Decimal) *OrderQueue
	)

	if side == Buy {
		level = ob.asks.MinPriceQueue()
		iter = ob.asks.GreaterThan
	} else {
		level = ob.bids.MaxPriceQueue()
		iter = ob.bids.LessThan
	}

	for quantity.Sign() > 0 && level != nil {
		levelVolume := level.Volume()
		levelPrice := level.Price()
		if quantity.GreaterThanOrEqual(levelVolume) {
			price = price.Add(levelPrice.Mul(levelVolume))
			quantity = quantity.Sub(levelVolume)
			level = iter(levelPrice)
		} else {
			price = price.Add(levelPrice.Mul(quantity))
			quantity = decimal.Zero
		}
	}

	if quantity.Sign() > 0 {
		err = ErrInsufficientQuantity
	}

	return
}

// String implements fmt.Stringer interface
func (ob *OrderBook) String() string {
	return ob.asks.String() + "\r\n------------------------------------" + ob.bids.String()
}

// MarshalJSON implements json.Marshaler interface
func (ob *OrderBook) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Asks *OrderSide `json:"asks"`
			Bids *OrderSide `json:"bids"`
		}{
			Asks: ob.asks,
			Bids: ob.bids,
		},
	)
}

// UnmarshalJSON implements json.Unmarshaler interface
func (ob *OrderBook) UnmarshalJSON(data []byte) error {
	obj := struct {
		Asks *OrderSide `json:"asks"`
		Bids *OrderSide `json:"bids"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	ob.asks = obj.Asks
	ob.bids = obj.Bids
	ob.orders = map[string]*list.Element{}

	for _, order := range ob.asks.Orders() {
		ob.orders[order.Value.(*Order).ID()] = order
	}

	for _, order := range ob.bids.Orders() {
		ob.orders[order.Value.(*Order).ID()] = order
	}

	return nil
}
