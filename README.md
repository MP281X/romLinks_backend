
# **RomLinks backend**
## **Find custom roms for your devices**
<p align="center"><a href="https://mp281x.xyz" target="_blank"><img src="./readme/logo.png"/></a></p>

## **Why?**
I wanted to create an app to help users to find and download custom rom for their devices in an easy way <br />
Everyone can contribute to the project adding new custom roms to the app or by leaving reviews to those already present <br /> 
This app was also an opportunity to build my first full-stack project and try new tecnology like `Docker` and `InfluxDB`
<br /><br />


## 📘 **Api documentation**

The api documentation is the dart code used in the [RomLinks Frontend](https://github.com/MP281X/romlinks_frontend) <br />
The code for the rest client is in the [logic](https://github.com/MP281X/romlinks_frontend/tree/master/lib/logic) folder<br />

### **Link to the `Dart` rest client for the RomLinks services:**
- [**Device Service**](https://github.com/MP281X/romlinks_frontend/blob/master/lib/logic/services/device_service.dart)

- [**File Storage Service**](https://github.com/MP281X/romlinks_frontend/blob/master/lib/logic/services/fileStorage_service.dart)

- [**Rom Service**](https://github.com/MP281X/romlinks_frontend/blob/master/lib/logic/services/rom_service.dart)

- [**User Service**](https://github.com/MP281X/romlinks_frontend/blob/master/lib/logic/services/user_service.dart)
<br /><br />

## 🔧 **Installation**

The api is free without limitation but if you want to install all the services locally and test them you can download the repository, create the .env and run the command below <br />
You need `docker`, `docker-compose` and `Make` to run the services

#### **Configiration file** .env
```sh
# influx db config
DOCKER_INFLUXDB_INIT_USERNAME=username
DOCKER_INFLUXDB_INIT_PASSWORD=password
influxToken=influxToken
influxOrg=orgName
influxBucket=bucketName
influxUri=http://influxdb:8086
metrics=false

# mongo config
MONGO_INITDB_ROOT_USERNAME=username
MONGO_INITDB_ROOT_PASSWORD=password
mongoUri=mongodb://username:password@mongodb:27017/?authSource=admin&readPreference=primary&ssl=false

# services config
jwtKey=jwtKey
logFile=false
tls=false

# grafana config
GF_SECURITY_ADMIN_PASSWORD=password
```

#### **Shell command**
```sh
# run all the services
make run
# stop all the services
make stop
# update all the container images
make update
```
<br />

## 👷‍♂️ **Project Structure**
The backend is subdivided in 4 independent microservices:
- [**Device Service:**](./services/deviceService) handle the devices
- [**File Storage Service:**](./services/fileStorageService) handle the images upload/download/compression and host the website
- [**Rom Service:**](./services/romService) handle the roms, the versions and the reviews data
- [**User Service:**](./services/userService) handle the users data and generate the jwt token for the users

The services code is in the [services](./services) folder <br />

- The `*Services.go` file contains the main function and another function for setting the db index or for generating the required folder <br />
- The `*Handler` folder contains the code of the api <br />
- The `_handler` file manage the api endpoint defined in the `routes` file <br />
- The `_model` file contains all the struct with the relative validation functions <br />
- The `_db` file contains all the db handler for storing and retriving data <br />

The [packages](./packages) folder contains all the code that is reused in all the services:
- The `api` folder contains all the gin setup and configuration and the `apiRes` helper function. It also contain the custom middleware for the metrics and for the errors log
- The `db` folder contains the mongodb setup code
- The `logger` folder contains the setup and the helper functions for writing logs to a file
- The `textSearch` folder container the code for the search function 


<br />

## 📜 **License**

[GNU General Public License v3.0](./LICENSE)
<br /><br />

## 📪 **Contact**

If you have any suggestion contact me

<img src="https://img.icons8.com/color/15/000000/telegram-app--v1.png"/> [**Telegram**](https://t.me/MP281X)
<br /><br />

## 🚀 **Languages and tools**

<p>
<a href="https://github.com/MP281X/romlinks_frontend" target="_blank"><img src="./readme/appIcon.png" height="44"/></a>
<a href="https://golang.org/" target="_blank"><img src="https://img.icons8.com/color/48/000000/golang.png" height="44"/></a>
<a href="https://github.com/gin-gonic/gin" target="_blank"><img src="./readme/gin.png" height="44"/></a>
<a href="https://www.mongodb.com/" target="_blank"><img src="https://img.icons8.com/color/50/000000/mongodb.png" height="44"/></a>
<a href="https://www.influxdata.com/" target="_blank"><img src="./readme/influxdb.png" height="44"/></a>
<a href="https://grafana.com/" target="_blank"><img src="./readme/grafana.png" height="44"/></a>
<a href="https://www.docker.com/" target="_blank"><img src="./readme/docker.png" height="44"/></a>
<a href="https://github.com/docker/compose" target="_blank"><img src="./readme/compose.png" height="44"/></a>
</p>