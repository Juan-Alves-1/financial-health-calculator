
FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY . .

RUN go build -o main main.go


FROM alpine:latest

RUN apk add libc6-compat

COPY --from=build /app/main /usr/local/bin/main

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/main"]