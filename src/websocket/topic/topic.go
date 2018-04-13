package topic

const MARKET = "market."
const KLINE = ".kline."
const DEPTH = ".depth."
const DETAIL = ".detail"
const TRADE = ".trade" + DEPTH

/**
 *market.$symbol.kline.$period
 */
func KLine(symbol string, period string) string {
	return MARKET + symbol + KLINE + period
}

/**
 *market.$symbol.depth.$type 	
 */
func MarketDepth(symbol string, types string) string {
	return MARKET + symbol + DEPTH + types
}

/**
 *market.$symbol.trade.detail
 */
func TradeDetail(symbol string) string {
	return MARKET + symbol + TRADE
}

/**
 *market.$symbol.detail
 */
func MarketDetail(symbol string) string {
	return MARKET + symbol + DETAIL
}