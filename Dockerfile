
FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY . .

RUN go build -o main main.go


FROM alpine:latest

RUN apk add libc6-compat

COPY --from=build /app/main /usr/local/bin/main
COPY --from=build /app/templates /templates 
# template directory could be more linux friendly

ENTRYPOINT ["/usr/local/bin/main"]