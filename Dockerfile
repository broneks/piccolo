FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY .env .
COPY . .

RUN go mod download

EXPOSE 8000

CMD [ "make", "run" ]
