# builder container
FROM golang:1.24-alpine AS builder

LABEL stage=builder

# installing dependencies
RUN apk add --no-cache \
    build-base leptonica-dev tesseract-ocr-dev tesseract-ocr-data-ind opencv-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /app/main .

# runner container
FROM alpine:latest AS production

RUN apk add --no-cache \
    leptonica-dev tesseract-ocr-dev tesseract-ocr-data-ind opencv-dev

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/main .

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8090

CMD ["./main"]
