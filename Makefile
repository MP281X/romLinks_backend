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
	ssh -t mp281x@mp281x.xyz "cd romLinks; make update; exit;"
	clear 
	docker-compose ps

# generate the ssl certificate 
ssl-certs:
	sudo certbot certonly --standalone -d romlinks.xyz  --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos
	sudo certbot certonly --standalone -d device.romlinks.xyz --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos
	sudo certbot certonly --standalone -d user.romlinks.xyz --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos
	sudo certbot certonly --standalone -d rom.romlinks.xyz --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos
	sudo certbot certonly --standalone -d filestorage.romlinks.xyz --staple-ocsp -m paludgnachmatteo.dev@gmail.com --agree-tos

	rm ./certs/*
	
	sudo cp /etc/letsencrypt/live/romlinks.xyz/fullchain.pem ./certs/website.pem
	sudo cp /etc/letsencrypt/live/romlinks.xyz/privkey.pem ./certs/website.key
	sudo chown mp281x ./certs/website.pem
	sudo chown mp281x ./certs/website.key

	sudo cp /etc/letsencrypt/live/device.romlinks.xyz/fullchain.pem ./certs/deviceService.pem
	sudo cp /etc/letsencrypt/live/device.romlinks.xyz/privkey.pem ./certs/deviceService.key
	sudo chown mp281x ./certs/deviceService.pem
	sudo chown mp281x ./certs/deviceService.key


	sudo cp /etc/letsencrypt/live/user.romlinks.xyz/fullchain.pem ./certs/userService.pem
	sudo cp /etc/letsencrypt/live/user.romlinks.xyz/privkey.pem ./certs/userService.key
	sudo chown mp281x ./certs/userService.pem
	sudo chown mp281x ./certs/userService.key

	sudo cp /etc/letsencrypt/live/rom.romlinks.xyz/fullchain.pem ./certs/romService.pem
	sudo cp /etc/letsencrypt/live/rom.romlinks.xyz/privkey.pem ./certs/romService.key
	sudo chown mp281x ./certs/romService.pem
	sudo chown mp281x ./certs/romService.key

	sudo cp /etc/letsencrypt/live/filestorage.romlinks.xyz/fullchain.pem ./certs/fileStorageService.pem
	sudo cp /etc/letsencrypt/live/filestorage.romlinks.xyz/privkey.pem ./certs/fileStorageService.key
	sudo chown mp281x ./certs/fileStorageService.pem
	sudo chown mp281x ./certs/fileStorageService.key

# backup the server data
DATE := $(shell date +%d-%m-%Y)
backup:
	rm -f ./romLinks_backup/backup_$(DATE).tar.gz
	ssh -t mp281x@mp281x.xyz "cd romLinks; sudo rm -f backup.tar.gz; sudo tar -zcvf backup.tar.gz data; sudo chmod +r backup.tar.gz; exit;"
	scp mp281x@mp281x.xyz:/home/mp281x/romLinks/backup.tar.gz ./romLinks_backup/backup_$(DATE).tar.gz
