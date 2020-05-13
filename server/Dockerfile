FROM golang:1.13.7-buster AS GO_SERVER

WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest

COPY --from=GO_SERVER /server .
ENTRYPOINT './server'
