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
	docker compose -f docker-compose.yaml down

	docker build -f ./services/userService/Dockerfile -t ghcr.io/mp281x/user-service:latest .
	docker build -f ./services/romService/Dockerfile -t ghcr.io/mp281x/rom-service:latest .
	docker build -f ./services/fileStorageService/Dockerfile -t ghcr.io/mp281x/file_storage-service:latest .
	docker build -f ./services/deviceService/Dockerfile -t ghcr.io/mp281x/device-service:latest .
	
	docker compose -f docker-compose.yaml up -d

	docker push ghcr.io/mp281x/user-service:latest
	docker push ghcr.io/mp281x/rom-service:latest
	docker push ghcr.io/mp281x/file_storage-service:latest
	docker push ghcr.io/mp281x/device-service:latest

	docker image prune 
