# Builder stage
FROM golang:1.22 as builder

WORKDIR /go/src/st-gateway

# Copy go.mod and go.sum files specific to st-gateway service
COPY src/st-gateway/go.mod .
COPY src/st-gateway/go.sum .

# Download dependencies
RUN go mod download

# Copy the source code of the st-gateway service
COPY src/st-gateway/cmd cmd/
COPY src/st-gateway/pkg pkg/

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/st-gateway cmd/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /go/bin/st-gateway .

# Define the arguments
ARG API_VERSION
ARG PORT
ARG AUTH_SVC_URL
ARG JOURNAL_SVC_URL
ARG JWT_SECRET_KEY

# Set the arguments as environment variables
ENV API_VERSION=${API_VERSION}
ENV PORT=:${PORT}
ENV AUTH_SVC_URL=${AUTH_SVC_URL}
ENV JOURNAL_SVC_URL=${JOURNAL_SVC_URL}
ENV JWT_SECRET_KEY=${JWT_SECRET_KEY}

EXPOSE 8080

CMD ["./st-gateway"]
