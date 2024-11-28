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

RUN if [ "$ENV" = "local" ]; then \
      go mod download; \
      go install github.com/air-verse/air@latest; \
    else \
      go build --o bin/piccolo ./cmd/piccolo; \
    fi

EXPOSE 8001

CMD if [ "$ENV" = "local" ]; then \
      air; \
    else \
      ./piccolo; \
    fi
