# build angular app
FROM node:12.7-alpine AS APP_BUILD
WORKDIR /usr/src/app
COPY ./dashboard/package.json ./
RUN npm install
COPY ./dashboard/ .
RUN npm update
RUN npm run-script build -- --base-href /dashboard/

# build go server
FROM golang:1.13.7-buster AS GO_SERVER
WORKDIR /
COPY ./server .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# copy angular app and go server to final container, execute server
FROM alpine:latest
COPY --from=GO_SERVER /server .
COPY --from=GO_SERVER /assetlinks.json .
COPY --from=APP_BUILD /usr/src/app/dist/dashboard ./dashboard
ENTRYPOINT './server'
