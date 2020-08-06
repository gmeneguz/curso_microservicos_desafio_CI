FROM golang:1.14.3-alpine3.11 AS builder

WORKDIR /go/src/app
COPY src ./src
RUN go build -o main ./src/*.go   

FROM scratch

COPY --from=builder /go/src/app/main /

CMD ["/main"]