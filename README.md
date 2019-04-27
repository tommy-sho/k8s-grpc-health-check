# Health checking of gRPC server on kubernetes cluster
Demo repository of  health-checking for gRPC server working on kubernetes.
Use [grpc-health-probe](https://github.com/grpc-ecosystem/grpc-health-probe) in this demo.
Kubernetes does not support gRPC server health checking natively.
The Liveness and readiness probes is what keep checking your application is available or not, They can detect unhealthy pods, avoid flow traffic to them. 

## Usage


## LICENSE
MIT



