# Builder phase, includes golang toolchain
FROM golang:1.20.11-alpine3.18 as builder
COPY . /src
WORKDIR /src
RUN go build -o mq2prom ./cmd

# Runtime phase, contains bare alpine plus the built binary and the config file
# IMPORTANT: keep the alpine version on the builder and the runtime base images aligned
FROM alpine:3.18
RUN mkdir /mq2prom
WORKDIR /mq2prom
COPY --from=builder /src/mq2prom /src/config.yaml .
CMD ["./mq2prom"]
