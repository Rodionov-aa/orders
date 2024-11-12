FROM golang:1.23 AS build_orders
COPY . /orders
WORKDIR /orders
RUN go build 







FROM alpine:3.20


RUN addgroup -g 1000 -S orders && \
adduser -u 1000 -h /orders -G orders -S orders
COPY --from=build_orders --chown=orders:orders /orders/cmd/orders /orders/orders
WORKDIR /orders
USER orders
CMD ["./orders"]