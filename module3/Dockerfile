FROM golang:1.11 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM ubuntu:20.04
COPY --from=build /httpserver/httpserver /httpserver/httpserver
#COPY --from=build /httpserver/etc /httpserver/etc
EXPOSE 8360ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]

FROM golang:1.11 AS build
WORKDIR /httpserver/
COPY . .ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM alpine:latest
COPY --from=build /httpserver/httpserver /httpserver/httpserver
#COPY --from=build /httpserver/etc /httpserver/etc
EXPOSE 8360
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]