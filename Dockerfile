# Build Stage
FROM golang:alpine3.20 AS builder
RUN apk add alpine-sdk

ENV CGO_ENABLED=1  \
    GOOS=linux

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy and build
COPY . ./
RUN go build --ldflags "-extldflags -static" -o application -tags musl


# Runner Stage
FROM alpine:3.20 AS runner
WORKDIR /app
EXPOSE 8080

# Copy binary from the builder stage
COPY --from=builder /app/application ./

CMD ["/app/application"]