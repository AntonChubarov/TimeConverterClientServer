
#
# NOTE: THIS DOCKERFILE IS GENERATED VIA "apply-templates.sh"
#
# PLEASE DO NOT EDIT IT DIRECTLY.
#

FROM golang:1.16.6-alpine3.14 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o . cmd/main.go
## Our start command which kicks off
## our newly created binary executable


FROM alpine:3.14

WORKDIR app

COPY --from=builder app/main app/main

EXPOSE 8080

CMD ["/main"]