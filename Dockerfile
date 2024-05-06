FROM node:22.0.0-slim as frontend

WORKDIR /app

COPY frontend/package.json ./

RUN npm install

COPY frontend/ .

RUN npm run build

FROM golang:1.22.2 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./server ./server
COPY ./frontend/frontend.go ./frontend/frontend.go
COPY --from=frontend /app/dist ./frontend/dist

RUN CGO_ENABLED=0 GOOS=linux go build -o out ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/out .

CMD ["./out"]
