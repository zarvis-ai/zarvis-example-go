FROM golang:1.10.1-alpine3.7 as builder
COPY main.go .
RUN go build -o /main main.go

FROM alpine:3.7  
CMD ["./main"]
COPY --from=builder /main .
