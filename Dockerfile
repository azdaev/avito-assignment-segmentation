FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./cmd/app/main.go

EXPOSE 8080
CMD [ "/app/main" ] 