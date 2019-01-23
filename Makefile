GCP_PROJECT_NAME=knative-samples
BINARY_NAME=tellmeall
RELEASE_VERSION=0.6.2

.PHONY: deps image policy

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
	kubectl apply -f https://raw.githubusercontent.com/mchmarny/tellmeall/master/deployments/tellmeall.yaml

gpu:
	kubectl apply -f https://raw.githubusercontent.com/GoogleCloudPlatform/container-engine-accelerators/master/nvidia-driver-installer/ubuntu/daemonset-preloaded.yaml

policy:
	PROJECT_NUMBER="$(gcloud projects describe ${PROJECT_ID} --format='get(projectNumber)')"
	gcloud projects add-iam-policy-binding ${PROJECT_NUMBER} \
    	--member=serviceAccount:${PROJECT_NUMBER}@cloudbuild.gserviceaccount.com \
    	--role=roles/container.developer


submit:
	gcloud builds submit \
		--config deployments/cloudbuild.yaml

tag:
	git tag "release-v${RELEASE_VERSION}"
	git push origin "release-v${RELEASE_VERSION}"
	git log --oneline