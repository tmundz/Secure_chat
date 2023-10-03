FROM golang:1.21.1-alpine3.18

WORKDIR /usr/src/app

#download modules
RUN go mod download
#Hot Reload
RUN go install github.com/cosmtrek/air@latest

#COPY . .
COPY . .
#expose 8080 for the http server access
EXPOSE 8080

#RUN go mod tidy
RUN go mod tidy