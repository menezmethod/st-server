FROM golang:1.22 as builder

WORKDIR /go/src/st-gateway

COPY src/st-gateway/go.mod .
COPY src/st-gateway/go.sum .

RUN go mod download

COPY src/st-gateway/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/st-gateway cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/bin/st-gateway .

ARG API_VERSION
ARG PORT
ARG AUTH_SVC_URL
ARG JOURNAL_SVC_URL
ARG JWT_SECRET_KEY

ENV API_VERSION=${API_VERSION}
ENV PORT=:${PORT}
ENV AUTH_SVC_URL=${AUTH_SVC_URL}
ENV JOURNAL_SVC_URL=${JOURNAL_SVC_URL}
ENV JWT_SECRET_KEY=${JWT_SECRET_KEY}

EXPOSE 8080

CMD ["./st-gateway"]
