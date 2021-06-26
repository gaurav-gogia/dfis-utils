FROM golang:latest

WORKDIR /go/src/dfis-utils
COPY . .

RUN apt-get update && apt-get install -y libpcap-dev

CMD go build -o dfis-utils .
