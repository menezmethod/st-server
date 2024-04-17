FROM python:3.11-slim as builder

WORKDIR /app

RUN apt-get update && apt-get install -y \
    build-essential \
    libssl-dev \
    libffi-dev \
    python3-dev

COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

FROM python:3.11-slim

WORKDIR /app

RUN apt-get update && apt-get install -y \
    libssl-dev \
    libffi-dev

RUN pip install grpcio
RUN pip install grpcio-reflection
RUN pip install openbb
RUN pip install crewai

COPY --from=builder /app /app

ARG PORT=50053
ENV PORT=${PORT}

EXPOSE ${PORT}

ENV PYTHONUNBUFFERED=1

CMD ["python", "server.py"]