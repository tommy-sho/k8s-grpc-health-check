FROM golang:1.12-alpine as builder
RUN apk add --no-cache ca-certificates git
ENV PROJECT /github.com/tommy-sho/k8s-grpc-health-check/backend
WORKDIR /go/src/$PROJECT

ENV GO111MODULE on
COPY . .
RUN go install .

FROM alpine as release
RUN apk add --no-cache ca-certificates git
COPY --from=builder /go/bin/backend /backend
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
ENV PORT_NUMBER=50001
ENTRYPOINT ["/backend"]
