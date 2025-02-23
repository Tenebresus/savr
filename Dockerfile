FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY ./entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh
RUN go mod download

EXPOSE 8080

COPY . . 

ARG CMD
ENV CMD=$CMD
RUN go build -o /bin/$CMD ./cmd/$CMD

ENTRYPOINT ["/entrypoint.sh"]
