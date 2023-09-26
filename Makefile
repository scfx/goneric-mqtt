build:
	#Build Go App
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o ./bin/main 
	#Build Docker Image
	docker-buildx build --platform linux/amd64 -t gonericmqtt_ms -f docker/Dockerfile .
	#Build Microservice
	docker save gonericmqtt_ms > bin/image.tar
	zip -j bin/gonericmqtt bin/image.tar docker/cumulocity.json
	#Clean up
	rm bin/image.tar
	rm bin/main

run:
	docker run -p 8080:80 --env-file .env gonericmqtt_ms 

deploy:
	c8y microservices create --file ./bin/gonericmqtt.zip