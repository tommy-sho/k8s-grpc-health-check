

PROJECT  := $(shell gcloud projects list| sed -n 2P |cut -d' ' -f1)
ZONE     := asia-northeast1-c
CLUSTER  := healthcheck-cluster
SERVICES := gateway backend
name     := hoge

.PHONY: build push apply delete proto

proto:
	protoc  \
	--go_out=plugins=grpc:. \
    ./proto/*.proto

build:
	for service in ${SERVICES}; do \
		docker image build -t grpc-health/$$service:latest $$service/; \
	done


apply:
	kubectl apply -f manifest/

delete:
	kubectl delete -f manifest/


run:
	@cd pkg/client && \
	go run main.go -name=${name}





gke-init:
	gcloud container clusters create ${CLUSTER} --zone=${ZONE} --num-nodes=3 --preemptible --machine-type=f1-micro --disk-size=10
	gcloud container clusters get-credentials ${CLUSTER} --zone=${ZONE}

gke-yaml:
	for service in ${SERVICES}; do \
		   sed -e 's/grpc-health/gcr.io\/${PROJECT}/g' ./manifest/$$service.yaml > ./manifest/$$service-gke.yaml; \
	done

gke-tag:
	for service in ${SERVICES}; do \
		docker tag grpc-health/$$service:latest gcr.io/${PROJECT}/$$service:latest; \
		docker push gcr.io/${PROJECT}/$$service:latest;\
	done