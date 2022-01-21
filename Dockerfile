FROM golang:1.17-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app/app
RUN go mod tidy -compat=1.17
RUN go build -o main

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app.env .
COPY --from=builder /app/app/main .
EXPOSE 8080
CMD ["./main"]