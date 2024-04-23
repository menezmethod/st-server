import logging
from langchain_groq import ChatGroq
from config.config import GROQ_API_KEY

logger = logging.getLogger(__name__)


class GroqModelLoader:
    def __init__(self):
        try:
            self.groq_model = ChatGroq(
                api_key=GROQ_API_KEY,
                model="mixtral-8x7b-32768"
            )
            logger.info("Groq model loaded successfully")
        except Exception as e:
            logger.error(f"Failed to load Groq model: {str(e)}")
            raise

    def get_model(self):
        return self.groq_model
