FROM golang:1.22 as builder

WORKDIR /go/src/st-auth-svc

COPY src/st-auth-svc/go.mod .
COPY src/st-auth-svc/go.sum .
RUN go mod download

COPY src/st-auth-svc/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/st-auth-svc cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/bin/st-auth-svc .
COPY --from=builder /go/src/st-auth-svc/pkg/db/user.yml /app/pkg/db/user.yml

ARG PORT
ARG DB_URL

ENV PORT=:${PORT} \
    DB_URL=${DB_URL}

EXPOSE 50051

CMD ["./st-auth-svc"]
