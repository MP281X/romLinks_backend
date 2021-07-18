package devicehandler

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (r *DbLog) editDeviceDB(codename string, device *EditDeviceModel, token string) (string, error) {

	// get the token data
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return "", err
	}

	// check if the user who modify is the same who uploaded the rom data
	var x DeviceModel
	err = r.Db.FindOne(context.TODO(), bson.M{"codename": codename}).Decode(&x)
	if err != nil {
		return "", logger.ErrDbRead
	}

	// check if the user is authorized
	if !tokenData.Moderator && tokenData.Username != x.CreatedBy {
		return "", logger.ErrUnauthorized
	}

	// replace the old info with the new one
	_, err = r.Db.UpdateOne(context.TODO(), bson.M{"codename": codename}, bson.M{
		"$set": device,
	})
	if err != nil {
		return "", errors.New("unable to edit the device info")
	}

	r.L.DbWrite("edited the info of " + x.Codename)

	return codename, nil
}

// get a list of device name
func (r *DbLog) searchDeviceNameDB(codename string) ([]string, error) {

	deviceCodenameList := []string{}

	// search the rom name in the db
	cursor, err := r.Db.Find(context.TODO(), bson.M{"$text": bson.M{"$search": codename}}, options.Find().SetSort(bson.D{}).SetLimit(3))
	if err != nil {
		fmt.Println(err.Error())
		return nil, logger.ErrDbRead
	}

	// add the device name to the rom name list
	for cursor.Next(context.TODO()) {

		var name bson.M
		if err = cursor.Decode(&name); err != nil {
			return nil, logger.ErrDbRead
		}
		deviceCodenameList = append(deviceCodenameList, name["codename"].(string))
	}

	// return the device name list
	return deviceCodenameList, nil
}

// get a list of devices uploaded by the user
func (r *DbLog) getUploadedDB(token string) ([]*DeviceModel, error) {

	// decode the device list there
	var deviceList []*DeviceModel

	// get the data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return nil, logger.ErrTokenRead
	}

	// search the roms in the db
	devices, err := r.Db.Find(context.TODO(), bson.M{"createdby": tokenData.Username}, options.Find().SetSort(bson.D{}))
	if err != nil {
		return nil, logger.ErrDbRead
	}

	defer devices.Close(context.TODO())
	if err = devices.All(context.TODO(), &deviceList); err != nil {
		return nil, logger.ErrDbRead
	}

	return deviceList, nil
}
