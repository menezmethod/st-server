# Builder stage
FROM golang:1.22 as builder

WORKDIR /go/src/st-journal-svc

# Copying go.mod and go.sum and downloading dependencies
COPY src/st-journal-svc/go.mod .
COPY src/st-journal-svc/go.sum .
RUN go mod download

# Copying the source code
COPY src/st-journal-svc/cmd cmd/
COPY src/st-journal-svc/pkg pkg/

# Building the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/st-journal-svc cmd/main.go

# Final stage
FROM alpine:latest

# Adding ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copying the binary from the builder stage
COPY --from=builder /go/bin/st-journal-svc .
COPY --from=builder /go/src/st-journal-svc/pkg/db/journal.yml /app/pkg/db/journal.yml


# Define the arguments with optional default values
ARG PORT
ARG DB_URL

# Set the arguments as environment variables
ENV PORT=:${PORT} \
    DB_URL=${DB_URL}

EXPOSE 50052

CMD ["./st-journal-svc"]
