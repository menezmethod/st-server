FROM python:3.11-slim as builder

WORKDIR /app

RUN apt-get update && apt-get install -y \
    build-essential \
    libssl-dev \
    libffi-dev \
    python3-dev

COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

FROM python:3.11-slim

WORKDIR /app

RUN apt-get update && apt-get install -y \
    libssl-dev \
    libffi-dev

COPY --from=builder /usr/local/lib/python3.11/site-packages /usr/local/lib/python3.11/site-packages

COPY . .

RUN pip install --no-cache-dir grpcio grpcio-tools protoc-gen-openapiv2 grpcio-reflection openbb crewai langchain-groq

ARG PORT=50053
ENV PORT=${PORT}

EXPOSE ${PORT}

ENV PYTHONUNBUFFERED=1

CMD ["python", "main.py"]