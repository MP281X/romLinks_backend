# run the auth service
userService:
	go run ./services/userService/userService.go

# run the device service
deviceService:
	go run ./services/deviceService/deviceService.go

# run the file storage service
fileStorageService:
	go run ./services/fileStorageService/fileStorageService.go

# run the rom service
romService:
	go run ./services/romService/romService.go

# build all the docker-image, push them to the registry and restart the service
docker:
	make docker-close
	make docker-build
	make docker-run
	make docker-push
	docker image prune 

# run all the services
docker-run:
	docker compose -f docker-compose.yaml up -d

# close all the services
docker-close:
	docker compose -f docker-compose.yaml down

# build all the image
docker-build:
	sudo docker build -f ./services/userService/Dockerfile -t ghcr.io/mp281x/user-service:latest .
	sudo docker build -f ./services/romService/Dockerfile -t ghcr.io/mp281x/rom-service:latest .
	sudo docker build -f ./services/fileStorageService/Dockerfile -t ghcr.io/mp281x/file_storage-service:latest .
	sudo docker build -f ./services/deviceService/Dockerfile -t ghcr.io/mp281x/device-service:latest .

# push all the image to the github registry
docker-push:
	docker push ghcr.io/mp281x/user-service:latest
	docker push ghcr.io/mp281x/rom-service:latest
	docker push ghcr.io/mp281x/file_storage-service:latest
	docker push ghcr.io/mp281x/device-service:latest

# clear data
docker-clear:
	sudo rm -r ./data