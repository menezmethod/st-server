# Builder stage
FROM golang:1.22 as builder

WORKDIR /go/src/st-auth-svc

# Copy go.mod and go.sum files and download dependencies
COPY src/st-auth-svc/go.mod .
COPY src/st-auth-svc/go.sum .
RUN go mod download

# Copy the source code
COPY src/st-auth-svc/cmd cmd/
COPY src/st-auth-svc/pkg pkg/

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/st-auth-svc cmd/main.go

# Final stage
FROM alpine:latest

# Add ca-certificates in case your service makes outbound HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /go/bin/st-auth-svc .
COPY --from=builder /go/src/st-auth-svc/pkg/db/user.yml /app/pkg/db/user.yml

# Define the arguments with optional default values
ARG PORT
ARG DB_URL

# Set the arguments as environment variables
ENV PORT=:${PORT} \
    DB_URL=${DB_URL}

EXPOSE 50051

CMD ["./st-auth-svc"]
