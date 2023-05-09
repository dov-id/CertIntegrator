FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/dov-id/CertIntegrator
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/CertIntegrator /go/src/github.com/dov-id/CertIntegrator


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/CertIntegrator /usr/local/bin/CertIntegrator
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["CertIntegrator"]
