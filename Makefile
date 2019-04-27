
SERVICES := gateway backend

.PHONY: image-build apply delete

image-build:
	for service in ${SERVICES}; do \
		docker image build -t local/$$service:latest $$service/; \
	done

apply:
	kubectl apply -f kubernetes-manifest

delete:
	kubectl delete -f kubernetes-manifest
