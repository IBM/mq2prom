FROM golang:1.16-alpine
RUN mkdir /mq2prom
COPY . /mq2prom
WORKDIR /mq2prom
RUN go build -o mq2prom ./cmd
CMD ./mq2prom
