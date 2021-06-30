FROM golang:alpine

RUN mkdir /apiservice
ADD . /apiservice
WORKDIR /apiservice

RUN go mod download
RUN go build -o service .

CMD ["./service"]