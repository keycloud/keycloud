FROM golang:1.13.7-buster AS GO_SERVER

WORKDIR /

RUN apt install git && git clone https://github.com/keycloud/keycloud.git && cd keycloud/server && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest

COPY --from=GO_SERVER /keycloud/server/config.ini .
COPY --from=GO_SERVER /keycloud/server/server .

ENTRYPOINT './server'
