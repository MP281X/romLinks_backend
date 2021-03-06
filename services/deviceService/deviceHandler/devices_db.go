package devicehandler

import (
	"context"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// add a device to the db
func (r *DbLog) addDeviceDB(device *DeviceModel, token string) error {

	// validate the input data
	err := device.ValidateDeviceData()
	if err != nil {
		return err
	}

	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// add the username to the rom data
	device.CreatedBy = tokenData.Username

	// add the device to the db
	_, err = r.Db.InsertOne(context.TODO(), device)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			return logger.ErrDeviceAlreadyExist
		}
		return logger.ErrDbWrite
	}

	r.L.Info(tokenData.Username + " added " + device.Codename)

	return nil
}

// get a device form the db
func (r *DbLog) getDeviceDB(codename string) (*DeviceModel, error) {

	codename = strings.ToLower(codename)
	var device DeviceModel

	// get the device info from the db
	err := r.Db.FindOne(context.TODO(), bson.M{"codename": codename}).Decode(&device)
	if err != nil {
		return nil, logger.ErrDbRead
	}

	return &device, nil
}

// get a list of device codename
func (r *DbLog) GetDeviceName() error {

	// search the device in the db
	cursor, err := r.Db.Find(context.TODO(), bson.M{}, nil)
	if err != nil {
		return logger.ErrDbRead
	}

	defer cursor.Close(context.TODO())

	// add the device codename to the device name slice
	for cursor.Next(context.TODO()) {
		var val bson.M

		if err = cursor.Decode(&val); err != nil {
			return logger.ErrDbRead
		}
		r.DN.AddValue(val["codename"].(string))
	}

	return nil
}
