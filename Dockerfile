FROM golang:alpine3.21 AS Builder

RUN apk add git bash gcc musl-dev upx

WORKDIR /app

ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN export BUILD_DATE=`date +%FT%T%z` && \
  go build -ldflags "-w -s" -o electerm-sync-server main.go
RUN upx -9 -o electerm-sync-server.minify electerm-sync-server && mv electerm-sync-server.minify electerm-sync-server

FROM alpine:3.21
COPY --from=Builder /app/electerm-sync-server /bin/electerm-sync-server

EXPOSE 7837
WORKDIR /app
ENTRYPOINT ["electerm-sync-server"]
