import json
from typing import List, Optional

from crewai import Task
from pydantic import BaseModel, ValidationError


class Insight(BaseModel):
    asset: Optional[str] = None
    category: Optional[str] = None
    suggestion: str


class AnalysisType(BaseModel):
    type: str
    insights: List[Insight]


class FinancialAnalysis(BaseModel):
    analyses: List[AnalysisType]


def create_financial_analysis_task(local_expert):
    return Task(
        description=(
            "Analyze the following financial transactions and trades provided in the input data.\n\n"
            "Use the latest stock prices and additional data from the market to assess each transaction. "
            "Generate a comprehensive analysis of the client's stock transactions, spending patterns, "
            "and dividend income. This should include profitability analysis, budgeting assessments, and "
            "strategic investment recommendations based on market trends and the client's risk profile.\n\n"
            "Your suggestions should help the client make informed financial decisions and optimize their "
            "investment portfolio. The output should be in a strict JSON format, adhering to the following "
            "Pydantic model schema:\n\n"
            f"{FinancialAnalysis.schema_json(indent=2)}\n\n"
            "Please provide only the JSON output, without any additional text or explanations."
        ),
        expected_output="A JSON object strictly adhering to the provided Pydantic model schema",
        agent=local_expert,
        human_input=False,
    )
