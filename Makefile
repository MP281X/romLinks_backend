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