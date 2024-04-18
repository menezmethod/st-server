import configparser

import grpc
import logging
from openbb import obb
from crewai import Crew, Agent, Task
from langchain_community.llms import Ollama

import helper_pb2
import helper_pb2_grpc

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

config = configparser.ConfigParser()
config.read('config.ini')
obb_pat = config['DEFAULT']['OPENBB_PAT']

obb.account.login(pat=obb_pat)


class OllamaModelLoader:
    def __init__(self):
        try:
            self.ollama_mistral = Ollama(model="mistral")
            logger.info("Ollama model loaded successfully")
        except Exception as e:
            logger.error(f"Failed to load Ollama model: {str(e)}")
            raise


class FinancialAnalysisService(helper_pb2_grpc.STHelperServicer):
    def __init__(self):
        self.ollama_model_loader = OllamaModelLoader()
        self.local_expert = self.create_local_expert()
        logger.info("FinancialAnalysisService initialized")

    def create_local_expert(self):
        logger.info("Creating local expert")
        return Agent(
            role='The Best Financial Analyst',
            goal="Impress all customers with your financial data and market trends analysis",
            backstory=(
                "As a highly experienced and knowledgeable financial analyst, you possess extensive expertise in "
                "overall financial analysis, investment strategies, and market trends. You are working for a "
                "prestigious firm and your goal is to provide top-notch financial advice and guidance to your "
                "valued clients. Your deep understanding of financial markets, coupled with your analytical skills "
                "and ability to interpret complex data, allows you to deliver comprehensive and actionable insights "
                "to help your clients make informed investment decisions and achieve their financial objectives. "
                "You pride yourself on your professionalism, integrity, and commitment to putting your clients' "
                "best interests first."
            ),
            verbose=True,
            llm=self.ollama_model_loader.ollama_mistral,
        )

    def AnalyzeFinancialData(self, request, context):
        try:
            logger.info("Analyzing financial data")
            task = self.create_financial_analysis_task()
            crew = self.create_crew(task)
            result = crew.kickoff()
            logger.info("Financial analysis completed successfully")
            return helper_pb2.FinancialResponse(suggestions=result)
        except Exception as e:
            error_msg = f"Failed to process request: {str(e)}"
            logger.error(f"Error during analysis: {error_msg}")
            context.set_details(error_msg)
            context.set_code(grpc.StatusCode.INTERNAL)
            return helper_pb2.FinancialResponse()

    def create_financial_analysis_task(self):
        logger.info("Creating financial analysis task")
        return Task(
            description=(
                "Analyze the following financial transactions and trades provided in the input data.\n\n"
                "Use the latest stock prices and additional data from the market to assess each transaction. "
                "Generate a comprehensive analysis of the client's stock transactions, spending patterns, "
                "and dividend income. This should include profitability analysis, budgeting assessments, and "
                "strategic investment recommendations based on market trends and the client's risk profile.\n\n"
                "Your suggestions should help the client make informed financial decisions and optimize their "
                "investment portfolio. The output should be in a JSON format, structured as detailed stock "
                "analyses, spending analyses, and dividend insights. Each section should reflect accurate "
                "calculations and tailored advice, without directing the Provide the insights in the following "
                "JSON format with corrected and precise data:\n"
                "[\n"
                "  {\n"
                "    \"type\": \"stock_analysis\",\n"
                "    \"insights\": [\n"
                "      {\n"
                "        \"stock\": \"STOCK_TICKER\",\n"
                "        \"remaining_quantity\": ACTUAL_REMAINING_QUANTITY,\n"
                "        \"buy_price\": ACCURATE_BUY_PRICE,\n"
                "        \"sell_price\": ACCURATE_SELL_PRICE,\n"
                "        \"current_price\": ACCURATE_CURRENT_PRICE,\n"
                "        \"profit_loss\": EXACT_PROFIT_LOSS,\n"
                "        \"suggestion\": \"DETAILED_INVESTMENT_ADVICE\"\n"
                "      },\n"
                "      ...\n"
                "    ]\n"
                "  },\n"
                "  {\n"
                "    \"type\": \"spending_analysis\",\n"
                "    \"insights\": [\n"
                "      {\n"
                "        \"category\": \"SPENDING_CATEGORY\",\n"
                "        \"amount\": EXACT_AMOUNT,\n"
                "        \"suggestion\": \"SPECIFIC_BUDGET_ADVICE\"\n"
                "      },\n"
                "      ...\n"
                "    ]\n"
                "  },\n"
                "  {\n"
                "    \"type\": \"dividend_analysis\",\n"
                "    \"insights\": [\n"
                "      {\n"
                "        \"stock\": \"STOCK_TICKER\",\n"
                "        \"dividend\": ACCURATE_DIVIDEND_AMOUNT,\n"
                "        \"suggestion\": \"STRATEGIC_REINVESTMENT_ADVICE\"\n"
                "      },\n"
                "      ...\n"
                "    ]\n"
                "  }\n"
                "]\n"
            ),
            expected_output='A JSON object with structured and comprehensive insights and suggestions to guide '
                            'the client\'s financial decision-making',
            agent=self.local_expert,
            human_input=False,
        )

    def create_crew(self, task):
        logger.info("Creating crew")
        return Crew(
            agents=[self.local_expert],
            tasks=[task],
            verbose=2
        )

    def GetStockQuote(self, request, context):
        try:
            logger.info(f"Fetching stock quote for {request.symbol}")
            provider = request.provider or "fmp"
            quote = obb.equity.price.quote(symbol=request.symbol, provider=provider)
            stock_quote = self.create_stock_quote_response(quote, request.provider)
            logger.info(f"Stock quote fetched successfully for {request.symbol}")
            return stock_quote
        except Exception as e:
            error_msg = f"Failed to fetch stock data for {request.symbol}: {str(e)}"
            logger.error(error_msg)
            context.set_details(error_msg)
            context.set_code(grpc.StatusCode.INTERNAL)
            return helper_pb2.StockQuoteResponse()

    def create_stock_quote_response(self, quote, provider):
        logger.info("Creating stock quote response")
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

    def GetHistoricalStockData(self, request, context):
        try:
            logger.info(f"Fetching historical stock data for {request.symbol}")
            provider = request.provider if request.provider else "fmp"
            historical_prices = obb.equity.price.historical(
                symbol=request.symbol,
                start_date=request.start_date,
                end_date=request.end_date,
                provider=provider
            )
            historical_stock_data = self.create_historical_stock_data_response(historical_prices)
            logger.info(f"Historical stock data fetched successfully for {request.symbol}")
            return historical_stock_data
        except Exception as e:
            error_msg = f"Failed to fetch historical stock data for {request.symbol}: {str(e)}"
            logger.error(error_msg)
            context.set_details(error_msg)
            context.set_code(grpc.StatusCode.INTERNAL)
            return helper_pb2.HistoricalStockDataResponse()

    def create_historical_stock_data_response(self, historical_prices):
        logger.info("Creating historical stock data response")
        quote_dict = historical_prices.to_dict()
        if not quote_dict:
            logger.warning("No historical stock data found")
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
