FROM golang:1.22 as builder

WORKDIR /go/src/st-journal-svc

COPY src/st-journal-svc/go.mod .
COPY src/st-journal-svc/go.sum .

RUN go mod download

COPY src/st-journal-svc/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/st-journal-svc cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/bin/st-journal-svc .
COPY --from=builder /go/src/st-journal-svc/pkg/db/journal.yml /app/pkg/db/journal.yml


ARG PORT
ARG DB_URL

ENV PORT=:${PORT} \
    DB_URL=${DB_URL}

EXPOSE 50052

CMD ["./st-journal-svc"]
