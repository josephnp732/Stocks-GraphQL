FROM golang:1.12-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/stocks-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/stocks-app .

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/stocks-app/out/stocks-app /app/stocks-app

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/stocks-app"]