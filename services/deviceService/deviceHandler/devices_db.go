package devicehandler

import (
	"context"
	"errors"
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
			return logger.ErrDuplicateKey
		}
		return logger.ErrDbWrite
	}

	r.L.DbWrite("added " + device.Codename + " to the db")

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

	r.L.DbRead("searched " + codename + " in the db")

	return &device, nil
}

//TODO: da fixare
func (r *DbLog) editDeviceDB(device *DeviceModel, token string) error {

	// validate the input data
	err := device.ValidateDeviceData()
	if err != nil {
		return err
	}

	// get the token data
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// check if the user is authorized
	if !tokenData.Moderator || tokenData.Username != device.CreatedBy {
		return logger.ErrUnauthorized
	}

	// replace the old info with the new one
	_, err = r.Db.ReplaceOne(context.TODO(), bson.M{"codename": device.Codename}, device)
	if err != nil {
		return errors.New("unable to edit the device info")
	}

	r.L.DbWrite("edited the info of " + device.Codename)

	return nil
}
