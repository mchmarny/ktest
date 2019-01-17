GCP_PROJECT_NAME=knative-samples
BINARY_NAME=tellmeall

all: test

run:
	go run main.go

deps:
	go mod tidy

image:
	gcloud builds submit \
		--project=$(GCP_PROJECT_NAME) \
		--tag gcr.io/$(GCP_PROJECT_NAME)/$(BINARY_NAME):latest .

image-test:
	gcloud builds submit \
		--project=$(GCP_PROJECT_NAME) \
		--tag gcr.io/$(GCP_PROJECT_NAME)/$(BINARY_NAME):test .

deploy:
	kubectl apply -f https://raw.githubusercontent.com/mchmarny/tellmeall/master/deployments/tellmeall.yaml

deploy-test:
	kubectl apply -f deployments/tellmeall-test.yaml

install-gpu-drivers:
	kubectl apply -f https://raw.githubusercontent.com/GoogleCloudPlatform/container-engine-accelerators/master/nvidia-driver-installer/ubuntu/daemonset-preloaded.yaml
