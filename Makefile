
# run all the services
run:
	docker-compose up -d
	clear
	docker-compose ps

# close all the services
stop:
	docker-compose down --remove-orphans 

# restart all the services
restart:
	make stop
	make run
	clear

# update all the services
update:
	docker pull ghcr.io/mp281x/user-service:latest
	docker pull ghcr.io/mp281x/rom-service:latest
	docker pull ghcr.io/mp281x/file_storage-service:latest
	docker pull ghcr.io/mp281x/device-service:latest
	docker pull mongo
	make stop
	make run
	docker image prune -f

# build all the docker-image, push them to the registry and restart the service
docker-build:
	docker build -f ./services/userService/Dockerfile -t ghcr.io/mp281x/user-service:latest .
	docker build -f ./services/romService/Dockerfile -t ghcr.io/mp281x/rom-service:latest .
	docker build -f ./services/fileStorageService/Dockerfile -t ghcr.io/mp281x/file_storage-service:latest .
	docker build -f ./services/deviceService/Dockerfile -t ghcr.io/mp281x/device-service:latest .
	make stop
	make run
	docker push ghcr.io/mp281x/user-service:latest
	docker push ghcr.io/mp281x/rom-service:latest
	docker push ghcr.io/mp281x/file_storage-service:latest
	docker push ghcr.io/mp281x/device-service:latest
	docker image prune -f
	clear 
	docker-compose ps

#! server specific command
# generate the ssl certificate 
ssl-certs:
	sudo certbot certonly --standalone -d mp281x.xyz --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos
	sudo certbot certonly --standalone -d romlinks.device.mp281x.xyz --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos
	sudo certbot certonly --standalone -d romlinks.user.mp281x.xyz --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos
	sudo certbot certonly --standalone -d romlinks.rom.mp281x.xyz --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos
	sudo certbot certonly --standalone -d romlinks.filestorage.mp281x.xyz --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos

	rm ./certs/*
	
	sudo cp /etc/letsencrypt/live/mp281x.xyz/fullchain.pem ./certs/website.pem
	sudo cp /etc/letsencrypt/live/mp281x.xyz/privkey.pem ./certs/website.key
	sudo chown mp281x ./certs/website.pem
	sudo chown mp281x ./certs/website.key

	sudo cp /etc/letsencrypt/live/romlinks.device.mp281x.xyz/fullchain.pem ./certs/deviceService.pem
	sudo cp /etc/letsencrypt/live/romlinks.device.mp281x.xyz/privkey.pem ./certs/deviceService.key
	sudo chown mp281x ./certs/deviceService.pem
	sudo chown mp281x ./certs/deviceService.key


	sudo cp /etc/letsencrypt/live/romlinks.user.mp281x.xyz/fullchain.pem ./certs/userService.pem
	sudo cp /etc/letsencrypt/live/romlinks.user.mp281x.xyz/privkey.pem ./certs/userService.key
	sudo chown mp281x ./certs/userService.pem
	sudo chown mp281x ./certs/userService.key

	sudo cp /etc/letsencrypt/live/romlinks.rom.mp281x.xyz/fullchain.pem ./certs/romService.pem
	sudo cp /etc/letsencrypt/live/romlinks.rom.mp281x.xyz/privkey.pem ./certs/romService.key
	sudo chown mp281x ./certs/romService.pem
	sudo chown mp281x ./certs/romService.key

	sudo cp /etc/letsencrypt/live/romlinks.filestorage.mp281x.xyz/fullchain.pem ./certs/fileStorageService.pem
	sudo cp /etc/letsencrypt/live/romlinks.filestorage.mp281x.xyz/privkey.pem ./certs/fileStorageService.key
	sudo chown mp281x ./certs/fileStorageService.pem
	sudo chown mp281x ./certs/fileStorageService.key

# backup the container data
DATE := $(shell date +%d-%m-%Y)
backup: 
	sudo tar -zcvf ./romLinks_backup/"backup_$(DATE).tar.gz" data
	clear
