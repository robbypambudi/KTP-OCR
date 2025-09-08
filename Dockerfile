# base container
from golang:1.24 AS builder

# builder container
FROM golang:1.24 AS builder

LABEL stage=builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update && apt upgrade -y && \
    apt install -y gcc leptonica-progs tesseract-ocr tesseract-ocr-ind libopencv-dev && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /app/main .

# runner container
FROM debian:latest AS production

RUN apt update && apt upgrade -y && \
    apt install -y gcc leptonica-progs tesseract-ocr tesseract-ocr-ind libopencv-dev && \
    rm -rf /var/lib/apt/lists/*

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/main .

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8090

CMD ["./main"]
