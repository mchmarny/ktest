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

deploy:
	kubectl apply -f https://raw.githubusercontent.com/mchmarny/tellmeall/master/app.yaml

deploy-local:
	kubectl apply -f app.yaml