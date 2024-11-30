FROM golang:1.23.2-alpine

WORKDIR /app

COPY . .

ARG ENV=production
ENV ENV=$ENV

RUN go mod download

RUN if [ "$ENV" = "local" ]; then \
      go install github.com/air-verse/air@latest; \
    else \
      go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
      go build -o ./bin/piccolo ./cmd/piccolo; \
    fi

EXPOSE 8000

CMD if [ "$ENV" = "local" ]; then \
      air; \
    else \
      ./bin/piccolo; \
    fi
