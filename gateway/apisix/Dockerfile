FROM golang:1.23 AS builder

COPY . /src
WORKDIR /src/gateway/apisix

RUN make -f ./Makefile build


FROM apache/apisix:3.6.0-redhat

COPY --from=builder /src/gateway/apisix/bin /app
