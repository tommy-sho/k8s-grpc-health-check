
SERVICES := gateway backend

.PHONY: image-build push apply delete

image-build:
	for service in ${SERVICES}; do \
		docker image build -t gcr.io/my-first-project-236315/grpc-test/$$service:latest $$service/; \
	done

push:
	for service in ${SERVICES}; do \
		docker push gcr.io/my-first-project-236315/grpc-test/$$service:latest; \
	done

apply:
	kubectl apply -f kubernetes-manifest/deployment-and-service

delete:
	kubectl delete -f kubernetes-manifest/deployment-and-service
