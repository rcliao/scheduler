FROM golang:alpine as builder

RUN mkdir /build
WORKDIR /build
COPY . /build
RUN go build -o main cmd/main.go

FROM alpine

COPY --from=builder /build/main /app/main
WORKDIR /app

CMD ["./main"]
