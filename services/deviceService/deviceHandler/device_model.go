package devicehandler

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	CreatedBy string             `bson:"createdby" json:"-"`
	Codename  string             `bson:"codename" json:"codename"`
	Name      string             `bson:"name" json:"name"`
}

func (device *DeviceModel) ValidateDeviceData() error {
	if device.Codename == "" {
		return errors.New("enter a codename")
	}
	device.Codename = strings.ToLower(device.Codename)
	if device.Name == "" {
		return errors.New("enter the device name")
	}

	return nil
}
