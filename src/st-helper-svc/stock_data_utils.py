import helper_pb2


def create_stock_quote_response(quote, provider):
    quote_dict = quote.to_dict()
    stock_data = {key: value[0] if isinstance(value, list) else value for key, value in quote_dict.items()}
    earnings_announcement = stock_data.get('earnings_announcement', None)
    earnings_announcement_str = earnings_announcement.strftime(
        '%Y-%m-%d %H:%M:%S') if earnings_announcement else ''

    market_cap_key = 'marketCap' if provider == 'yfinance' else 'market_cap'
    market_cap = int(stock_data.get(market_cap_key, 0))

    return helper_pb2.StockQuoteResponse(
        symbol=stock_data['symbol'],
        name=stock_data['name'],
        exchange=stock_data['exchange'],
        lastPrice=stock_data['last_price'],
        open=stock_data.get('open', 0),
        high=stock_data.get('high', 0),
        low=stock_data.get('low', 0),
        volume=stock_data['volume'],
        prevClose=stock_data.get('prev_close', 0),
        change=stock_data.get('change', 0),
        changePercent=stock_data.get('change_percent', 0),
        yearHigh=stock_data.get('year_high', 0),
        yearLow=stock_data.get('year_low', 0),
        marketCap=market_cap,
        sharesOutstanding=stock_data.get('shares_outstanding', 0),
        pe=stock_data.get('pe', 0),
        earningsAnnouncement=earnings_announcement_str,
        eps=stock_data.get('eps', 0),
        sector=stock_data.get('sector', ''),
        industry=stock_data.get('industry', ''),
        beta=stock_data.get('beta', 0)
    )

def create_historical_stock_data_response(historical_prices):
    quote_dict = historical_prices.to_dict()
    if not quote_dict:
        return helper_pb2.HistoricalStockDataResponse()

    historical_data = {
        'open': quote_dict['open'],
        'high': quote_dict['high'],
        'low': quote_dict['low'],
        'close': quote_dict['close'],
        'volume': quote_dict['volume'],
        'vwap': quote_dict.get('vwap', []),
        'adj_close': quote_dict.get('adj_close', []),
        'change': quote_dict.get('change', []),
        'change_percent': quote_dict.get('change_percent', [])
    }

    return helper_pb2.HistoricalStockDataResponse(
        open=historical_data['open'],
        high=historical_data['high'],
        low=historical_data['low'],
        close=historical_data['close'],
        volume=historical_data['volume'],
        vwap=historical_data['vwap'],
        adj_close=historical_data['adj_close'],
        change=historical_data['change'],
        change_percent=historical_data['change_percent']
    )