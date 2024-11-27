FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/piccolo ./cmd/piccolo

FROM golang:1.23.2-alpine AS runner

WORKDIR /root/

COPY . .
COPY --from=builder /app/bin/piccolo .

ARG ENV=production
ENV ENV=$ENV
ENV FLY_DATABASE_URL=$FLY_DATABASE_URL

RUN if [ "$ENV" = "local" ]; then \
      go mod download; \
      go install github.com/air-verse/air@latest; \
    else \
      go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
    fi

EXPOSE 8001

CMD if [ "$ENV" = "local" ]; then \
      air; \
    else \
      ./piccolo; \
    fi
