run:
	go run main.go
docker-image-build:
	docker image build -t bank_test .
docker-image-run:
	docker image run -p 8080:8080 bank_test
docker-compose:
	docker-compose up -d