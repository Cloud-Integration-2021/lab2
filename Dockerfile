FROM golang:alpine3.13 as builder

RUN apk update && apk add git
COPY . $GOPATH/src/github.com/lab2/
WORKDIR $GOPATH/src/github.com/lab2/
RUN go build -o /go/bin/lab2

ENV ENVIRONMENT prod

FROM alpine:3.13.1
EXPOSE 8080
COPY --from=builder /go/bin/lab2 /bin/lab2

ENTRYPOINT ["/bin/lab2"]