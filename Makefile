GCP_PROJECT_NAME=knative-samples
BINARY_NAME=ktest
RELEASE_VERSION=0.6

.PHONY: deps image

all: test

run:
	go run main.go

deps:
	go mod tidy tag submit

image:
	gcloud builds submit \
		--project=$(GCP_PROJECT_NAME) \
		--tag gcr.io/$(GCP_PROJECT_NAME)/$(BINARY_NAME):latest .

deploy:
	kubectl apply -f https://raw.githubusercontent.com/mchmarny/ktest/master/deployments/ktest.yaml

gpu:
	kubectl apply -f https://raw.githubusercontent.com/GoogleCloudPlatform/container-engine-accelerators/master/nvidia-driver-installer/ubuntu/daemonset-preloaded.yaml

submit:
	gcloud builds submit \
		--config deployments/cloudbuild.yaml

tag:
	git tag "release-v${RELEASE_VERSION}"
	git push origin "release-v${RELEASE_VERSION}"
	git log --oneline