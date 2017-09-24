FROM golang:1.8.3-jessie

RUN mkdir -p /opt/lfi
WORKDIR /opt/lfi
COPY logstash-fileimporter.go .
RUN go build

CMD ./lfi