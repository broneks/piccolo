FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download
RUN go install github.com/air-verse/air

EXPOSE 8000

CMD [ "make", "run" ]
