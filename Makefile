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

# clear the log 
clear_log:
	sudo rm -r ./log/

# clear the build
clear_build:
	sudo rm -r ./build/

# clear the asset
clear_asset:
	sudo rm -r ./asset/

# build all the service
buildAll:
	go mod download
	mkdir ./build
	go build -o ./build ./services/authService/authService.go
	go build -o ./build ./services/deviceService/deviceService.go
	go build -o ./build ./services/fileStorageService/fileStorageService.go
	go build -o ./build ./services/romService/romService.go

# download the requried go mod
downloadGoMod:
	go mod download