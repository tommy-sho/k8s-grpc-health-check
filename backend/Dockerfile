FROM golang:1.12-alpine as builder
RUN apk add --no-cache ca-certificates git
ENV PROJECT /github.com/tommy-sho/grpc-loadbalncing/backend
WORKDIR /go/src/$PROJECT

ENV GO111MODULE on
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/backend

FROM alpine as release
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/backend /backend
ENV PORT_NUMBER=50001
ENTRYPOINT ["/backend"]
