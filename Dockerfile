FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/dov-id/cert-integrator-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/cert-integrator-svc /go/src/github.com/dov-id/cert-integrator-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/cert-integrator-svc /usr/local/bin/cert-integrator-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["cert-integrator-svc"]
