FROM golang:1.16-alpine

WORKDIR /home/sprout/Documents/Go/buddy_test

COPY go.mod /home/go
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /docker/buddy_test

EXPOSE 8080

CMD ["/docker/buddy_test"]