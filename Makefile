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

# clear the docker images
clear_images:
	sudo rm -r ./docker/*

#build all the docker image
buildAll:
	docker build -f ./services/userService/Dockerfile -t user-service:latest .
	docker build -f ./services/romService/Dockerfile -t rom-service:latest .
	docker build -f ./services/fileStorageService/Dockerfile -t file_storage-service:latest .
	docker build -f ./services/deviceService/Dockerfile -t device-service:latest .

#save all the docker image
saveImages:
	docker save -o ./docker/user-service.tar user-service
	docker save -o ./docker/rom-service.tar rom-service
	docker save -o ./docker/file_storage-service.tar file_storage-service
	docker save -o ./docker/device-service.tar device-service

# disable git tracking for docker-compose file
stopGitDockerfile:
	git update-index --assume-unchanged docker-compose.yaml

# disable git tracking for docker-compose file
startGitDockerfile:
	git update-index --no-assume-unchanged docker-compose.yaml

runDockerCompose:
	docker compose -f docker-compose.yaml up -d

closeDockerCompose:
	docker compose -f docker-compose.yaml down
