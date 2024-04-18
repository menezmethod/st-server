FROM golang:1.22 as builder

WORKDIR /go/src/st-server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/st-gateway ./src/st-gateway/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/bin/st-gateway .

ARG API_VERSION
ARG PORT
ARG AUTH_SVC_URL
ARG JOURNAL_SVC_URL
ARG HELPER_SVC_URL
ARG JWT_SECRET_KEY

ENV API_VERSION=${API_VERSION} \
    PORT=:${PORT} \
    AUTH_SVC_URL=${AUTH_SVC_URL} \
    JOURNAL_SVC_URL=${JOURNAL_SVC_URL} \
    HELPER_SVC_URL=${HELPER_SVC_URL} \
    JWT_SECRET_KEY=${JWT_SECRET_KEY}

EXPOSE 8080

CMD ["./st-gateway"]
