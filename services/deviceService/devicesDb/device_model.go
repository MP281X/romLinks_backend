package devicesdb

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceModel struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Codename        string             `bson:"codename" json:"codename"`
	Name            string             `bson:"name" json:"name"`
	Photo           []string           `bson:"photo" json:"photo"`
	Brand           string             `bson:"brand" json:"brand"`
	Specs           *SpecsModel        `bson:"specs" json:"specs"`
	BootloaderLinks []string           `bson:"bootloaderlinks" json:"bootloaderlinks"`
	RecoveryLinks   []string           `bson:"recoverylinks" json:"recoverylinks"`
	GcamLinks       []string           `bson:"gcamlinks" json:"gcamlinks"`
}

type SpecsModel struct {
	Camera    string `bson:"camera" json:"camera"`
	Battery   int    `bson:"battery" json:"battery"`
	Processor string `bson:"processor" json:"processor"`
}

func (device *DeviceModel) ValidateDeviceData() error {
	if device.Codename == "" {
		return errors.New("enter a codename")
	}
	device.Codename = strings.ToLower(device.Codename)
	if device.Name == "" {
		return errors.New("enter the device name")
	}
	device.Name = strings.ToLower(device.Name)
	if len(device.Photo) == 0 {
		return errors.New("upload an image")
	}
	device.Brand = strings.ToLower(device.Brand)
	if device.Brand == "" {
		return errors.New("enter the device brand name")
	}
	if len(device.BootloaderLinks) == 0 {
		device.BootloaderLinks = []string{}
	}
	if len(device.RecoveryLinks) == 0 {
		device.RecoveryLinks = []string{}
	}
	if len(device.GcamLinks) == 0 {
		device.GcamLinks = []string{}
	}
	if device.Specs == nil {
		device.Specs = &SpecsModel{}
	}

	return nil
}
