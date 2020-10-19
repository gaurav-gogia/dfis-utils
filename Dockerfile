FROM golang:1.15

WORKDIR /go/src/dfis-utils
COPY . .

RUN apt-get update && apt-get install -y libpcap-dev

CMD go build -o dfis-utils .
