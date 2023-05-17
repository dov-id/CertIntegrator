FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/dov-id/CertIntegrator-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/CertIntegrator-svc /go/src/github.com/dov-id/CertIntegrator-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/CertIntegrator-svc /usr/local/bin/CertIntegrator-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["CertIntegrator-svc"]
