FROM golang:1.15.5-alpine3.12 AS builder

RUN apk update && apk add --no-cache git
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/http/
COPY ./src/looping .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/http

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /go/bin/http /go/bin/http

USER appuser

EXPOSE 8000

ENTRYPOINT ["/go/bin/http"]