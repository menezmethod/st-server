import json
import logging

import grpc
from crewai import Crew, Agent
from google.protobuf.internal.well_known_types import Any
from google.protobuf.json_format import ParseDict
from openbb import obb

from model_loader import GroqModelLoader
from financial_analysis_task import create_financial_analysis_task, FinancialAnalysis
from stock_data_utils import create_stock_quote_response, create_historical_stock_data_response
import helper_pb2
import helper_pb2_grpc

logger = logging.getLogger(__name__)


class FinancialAnalysisService(helper_pb2_grpc.STHelperServicer):
    def __init__(self):
        self.groq_model_loader = GroqModelLoader()
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
            llm=self.groq_model_loader.get_model(),
        )

    def AnalyzeFinancialData(self, request, context):
        try:
            logger.info("Analyzing financial data")
            # Create the financial analysis task
            task = create_financial_analysis_task(self.local_expert)
            # Get the model for analysis
            groq_model_loader = GroqModelLoader()
            model = groq_model_loader.get_model()

            # Use structured output with the LLM and get the result
            structured_llm = model.with_structured_output(FinancialAnalysis, method="json_mode")
            result = structured_llm.invoke(task.description)

            # Check if it's already a JSON string
            if isinstance(result, str):
                suggestions_json = result
            else:
                # Convert to JSON string if it's not
                suggestions_json = json.dumps(result)

            logger.info("Financial analysis completed successfully")
            # Return the raw JSON string in the gRPC response
            return helper_pb2.FinancialResponse(suggestions=suggestions_json)
        except Exception as e:
            # Handle exceptions and return an error response
            error_msg = f"Failed to process request: {str(e)}"
            logger.error(f"Error during analysis: {error_msg}")
            context.set_details(error_msg)
            context.set_code(grpc.StatusCode.INTERNAL)
            return helper_pb2.FinancialResponse(suggestions="{}")

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
            stock_quote = create_stock_quote_response(quote, request.provider)
            logger.info(f"Stock quote fetched successfully for {request.symbol}")
            return stock_quote
        except Exception as e:
            error_msg = f"Failed to fetch stock data for {request.symbol}: {str(e)}"
            logger.error(error_msg)
            context.set_details(error_msg)
            context.set_code(grpc.StatusCode.INTERNAL)
            return helper_pb2.StockQuoteResponse()

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
            historical_stock_data = create_historical_stock_data_response(historical_prices)
            logger.info(f"Historical stock data fetched successfully for {request.symbol}")
            return historical_stock_data
        except Exception as e:
            error_msg = f"Failed to fetch historical stock data for {request.symbol}: {str(e)}"
            logger.error(error_msg)
            context.set_details(error_msg)
            context.set_code(grpc.StatusCode.INTERNAL)
            return helper_pb2.HistoricalStockDataResponse()
