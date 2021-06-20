# run the auth service
authService:
	go run ./services/authService/authService.go

# run the device service
deviceService:
	go run ./services/deviceService/deviceService.go

# run the file storage service
fileStorageService:
	go run ./services/fileStorageService/fileStorageService.go

# run the rom service
romService:
	go run ./services/romService/romService.go

# clear all the log file
clear_log:
	rm ./log/*

# clear all the build
clear_build:
	rm ./build/*

# build all the service
buildAll:
	go build -o ./build ./services/authService/authService.go
	go build -o ./build ./services/deviceService/deviceService.go
	go build -o ./build ./services/fileStorageService/fileStorageService.go
	go build -o ./build ./services/romService/romService.go

# remove the config.yaml file from the git changes
untrackConfigFile:
	git update-index --assume-unchanged config/authService_config.yaml
	git update-index --assume-unchanged config/deviceService_config.yaml
	git update-index --assume-unchanged config/fileStorageService_config.yaml
	git update-index --assume-unchanged config/romService_config.yaml

# add the config.yaml file to the git changes
trackConfigFile:
	git update-index --no-assume-unchanged config/authService_config.yaml
	git update-index --no-assume-unchanged config/deviceService_config.yaml
	git update-index --no-assume-unchanged config/fileStorageService_config.yaml
	git update-index --no-assume-unchanged config/romService_config.yaml
