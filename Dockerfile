#build stage
FROM golang:1.16-alpine3.13 as builder
WORKDIR /app
COPY . .
RUN go build -o main

#Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8081
CMD ["/app/main"]