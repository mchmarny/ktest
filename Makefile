GCP_PROJECT_NAME=s9-demo
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
	kubectl apply -f app.yaml