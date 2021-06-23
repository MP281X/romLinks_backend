package devicesdb

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var DeviceCollection *mongo.Collection

func AddDevice(device *DeviceModel) error {
	// validate the input data
	err := device.ValidateDeviceData()
	if err != nil {
		return err
	}
	// add the device to the db
	_, err = DeviceCollection.InsertOne(context.TODO(), device)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			return errors.New("codename already used")
		}
		fmt.Println(err.Error())
		return errors.New("unable to add the device to the db")
	}
	logger.DbWrite("inserted " + device.Codename + " in the db")
	return nil
}

func GetDevice(codename string) (*DeviceModel, error) {
	var device DeviceModel
	// get the device info from the db
	logger.DbRead("searched " + codename + " in the db")
	err := DeviceCollection.FindOne(context.TODO(), bson.M{"codename": codename}).Decode(&device)
	if err != nil {
		return nil, errors.New("invalid device codename")
	}
	return &device, nil
}

func EditDevice(device *DeviceModel) error {
	// validate the input data
	err := device.ValidateDeviceData()
	if err != nil {
		return err
	}
	// replace the old info with the new one
	logger.DbWrite("edit the info of " + device.Codename)
	_, err = DeviceCollection.ReplaceOne(context.TODO(), bson.M{"codename": device.Codename}, device)
	if err != nil {
		return errors.New("unable to edit the device info")
	}

	return nil
}
