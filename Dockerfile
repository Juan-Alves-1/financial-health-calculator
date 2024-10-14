# Build stage
FROM golang:1.23.2 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -tags musl -ldflags '-w -extldflags "-static"' -a -installsuffix cgo -o main main.go

# Final stage
FROM alpine:latest
COPY --from=build /app/main /usr/local/bin/main
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/main"]
