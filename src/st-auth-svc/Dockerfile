FROM golang:1.22 as builder

WORKDIR /go/src/st-server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/st-auth-svc ./src/st-auth-svc/cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/bin/st-auth-svc .
COPY --from=builder /go/src/st-server/src/st-auth-svc/pkg/db/user.yml /app/pkg/db/user.yml

ARG PORT
ARG DB_URL

ENV PORT=:${PORT} \
    DB_URL=${DB_URL}

EXPOSE 50051

CMD ["./st-auth-svc"]
