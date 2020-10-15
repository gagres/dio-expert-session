build-image:
	docker build -t gagres/finance .

runapp:
	docker-compose -f .devops/app.yml up -d

lint:
	golint ./...
	go fmt -n ./...

unit_test:
	go test -v ./...