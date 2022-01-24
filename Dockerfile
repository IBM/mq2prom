FROM golang:1.16-alpine
RUN mkdir /mq2p
COPY . /mq2p
WORKDIR /mq2p
RUN go build -o mq2p ./cmd
EXPOSE 9641
CMD ./mq2p
