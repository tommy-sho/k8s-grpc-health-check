FROM golang:1.12-alpine as builder
RUN apk add --no-cache ca-certificates git

ENV PROJECT /github.com/tommy-sho/grpc-loadbalncing/gateway
WORKDIR /go/src/$PROJECT

ENV GO111MODULE on
COPY . .
RUN go install .

FROM alpine as release
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/gateway /gateway
ENV PORT_NUMBER=50000
ENTRYPOINT ["/gateway"]
