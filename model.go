/*
@Time : 2019-05-07 19:28
@Author : mocha
@File : model
*/
package mex

import "time"

type TradeOrder struct {
	Id            int64     `db:"id" json:"id"`
	Side          int       `db:"side" json:"side"`
	PriceType     int       `db:"price_type" json:"price_type"`
	BaseCurrency  string    `db:"base_currency" json:"base_currency"`
	QuoteCurrency string    `db:"quote_currency" json:"quote_currency"`
	Amount        float64   `db:"amount" json:"amount"`
	Price         float64   `db:"price" json:"price"`
	BusinessId    int64     `db:"business_id" json:"business_id"`
	MarketOrderId int64     `db:"market_order_id" json:"market_order_id"`
	FilledAmount  float64   `db:"filled_amount" json:"filled_amount"`
	FilledMoney   float64   `db:"filled_money" json:"filled_money"`
	FilledFee     float64   `db:"filled_fee" json:"filled_fee"`
	State         int       `db:"state" json:"state"`
	CreateTime    time.Time `db:"create_time" json:"create_time"`
	CanceledTime  time.Time `db:"canceled_time" json:"canceled_time"`
	FinishTime    time.Time `db:"finish_time" json:"finish_time"`
}

//交易对
type TradeSymbol struct {
	BaseCurrency  string `db:"base_currency" json:"base_currency"`   //基础币种
	QuoteCurrency string `db:"quote_currency" json:"quote_currency"` //计价币种
}

// 子账户结构
type SubAccount struct {
	Currency string `json:"currency"` // 币种
	Balance  string `json:"balance"`  // 结余
	Type     string `json:"type"`     // 类型, trade: 交易余额, frozen: 冻结余额
}

//账户余额
type AccountBalance struct {
	ID     int64        `json:"id"`    // 账户ID
	State  string       `json:"state"` // 账户状态, working: 正常, lock: 账户被锁定
	Type   string       `json:"type"`  // 账户类型, spot: 现货账户
	List   []SubAccount `json:"list"`  // 子账户数组
	UserID int64        `json:"user-id"`
}

type AccountsData struct {
	ID     int64  `json:"id"`      // Account ID
	Type   string `json:"type"`    // 账户类型, spot: 现货账户
	State  string `json:"state"`   // 账户状态, working: 正常, lock: 账户被锁定
	UserID int64  `json:"user-id"` // 用户ID
}

type KLineData struct {
	ID     int64   `json:"id"`     // K线ID
	Amount float64 `json:"amount"` // 成交量
	Count  int64   `json:"count"`  // 成交笔数
	Open   float64 `json:"open"`   // 开盘价
	Close  float64 `json:"close"`  // 收盘价, 当K线为最晚的一根时, 时最新成交价
	Low    float64 `json:"low"`    // 最低价
	High   float64 `json:"high"`   // 最高价
	Vol    float64 `json:"vol"`    // 成交额, 即SUM(每一笔成交价 * 该笔的成交数量)
}

type KLine struct {
	Ts   int64       `json:"ts"`   // 响应生成时间点, 单位毫秒
	Data []KLineData `json:"data"` // KLine数据
	Ch   string      `json:"ch"`   // 数据所属的Channel, 格式: market.$symbol.kline.$period
}

type MarketDepth struct {
	ID   int64       `json:"id"`   // 消息ID
	Ts   int64       `json:"ts"`   // 消息声称事件, 单位: 毫秒
	Bids [][]float64 `json:"bids"` // 买盘, [price(成交价), amount(成交量)], 按price降序排列
	Asks [][]float64 `json:"asks"` // 卖盘, [price(成交价), amount(成交量)], 按price升序排列
}

type Depth struct {
	Ts   int64       `json:"ts"`   // 响应生成时间点, 单位: 毫秒
	Tick MarketDepth `json:"tick"` // Depth数据
	Ch   string      `json:"ch"`   //  数据所属的Channel, 格式: market.$symbol.depth.$type
}

type MarketDetail struct {
	ID     int64   `json:"id"`     // 消息ID
	Ts     int64   `json:"ts"`     // 24小时统计时间
	Amount float64 `json:"amount"` // 24小时成交量
	Open   float64 `json:"open"`   // 前24小时成交价
	Close  float64 `json:"close"`  // 当前成交价
	High   float64 `json:"high"`   // 近24小时最高价
	Low    float64 `json:"low"`    // 近24小时最低价
	Count  int64   `json:"count"`  // 近24小时累计成交数
	Vol    float64 `json:"vol"`    // 近24小时累计成交额, 即SUM(每一笔成交价 * 该笔的成交量)
}

type OrderDetail struct {
	Id              int64  `json:"id"`
	Symbol          string `json:"symbol"`
	AccountId       int64  `json:"account-id"`
	Amount          string `json:"amount"`
	Price           string `json:"price"`
	CreateTime      int64  `json:"created-at"`
	Type            string `json:"type"`
	FiledAmount     string `json:"field-amount"`
	FiledCashAmount string `json:"field-cash-amount"`
	FiledFee        string `json:"field-fees"`
	FinishTime      int64  `json:"finished-at"`
	Api             string `json:"api"`
	CanceledTime    int64  `json:"canceled-at"`
	Exchange        string `json:"exchange"`
	Batch           string `json:"batch"`
	State           string `json:"state"`
}

type Ticker struct {
	ID     int64     `json:"id"`     // K线ID
	Amount float64   `json:"amount"` // 成交量
	Count  int64     `json:"count"`  // 成交笔数
	Open   float64   `json:"open"`   // 开盘价
	Close  float64   `json:"close"`  // 收盘价
	Low    float64   `json:"low"`    // 最低价
	High   float64   `json:"high"`   // 最高价
	Vol    float64   `json:"vol"`    // 成交额
	Bid    []float64 `json:"bid"`    // [买1价, 买1量]
	Ask    []float64 `json:"ask"`    // [卖1价, 卖1量]
}
