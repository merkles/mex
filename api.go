package mex

import (
	"coinx/pkg/model"
)

// api interface

type API interface {
	LimitBuy(amount, price string, symbol TradeSymbol) (int64, error)
	LimitSell(amount, price string, symbol TradeSymbol) (int64, error)
	MarketBuy(amount, price string, symbol TradeSymbol) (int64, error)
	MarketSell(amount, price string, symbol TradeSymbol) (int64, error)
	CancelOrder(orderId string) error
	GetOrderDetail(orderId string) (*TradeOrder, error)
	GetUnFinishOrders(symbol TradeSymbol) ([]model.TradeOrder, error)
	GetHistoryOrders(symbol TradeSymbol, currentPage, pageSize int) ([]model.TradeOrder, error)
	GetBalance() (*AccountBalance, error)

	GetTicker(symbol TradeSymbol) (Ticker, error)
	GetDepth(size int, order TradeOrder) (MarketDepth, error)
	GetKlineRecords(symbol TradeSymbol, period, size, since int) ([]KLineData, error)
	//非个人，整个交易所的交易记录
	GetTrades(symbol TradeSymbol) ([]TradeOrder, error)

	GetExchangeName() string
}
