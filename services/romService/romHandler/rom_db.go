package romhandler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// add a new rom to the db
func (r *DbLog) addRomDB(rom *RomModel, token string) error {

	// validate the input data
	err := rom.Validate()
	if err != nil {
		return err
	}

	// get the data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return logger.ErrTokenRead
	}

	// set the uploader name and the current time
	for _, romVersion := range rom.Version {
		romVersion.UploadedBy = tokenData.Username
		if romVersion.Date.IsZero() {
			romVersion.Date = time.Now()
		}
	}

	// if the user is a mod or is verified the rom is pubblic without additional validation
	if tokenData.Moderator || tokenData.Verified {
		rom.Verified = true
	}

	// insert the rom in the db
	_, err = r.Db.InsertOne(context.TODO(), rom)
	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.DbWrite("added a new rom")

	return nil
}

// get the data of a rom
func (r *DbLog) getRomDB(codename string, androidVersion float32, romName string) (*RomModel, error) {

	codename = strings.ToLower(codename)
	romName = strings.ToLower(romName)

	// decode the rom data there
	var rom RomModel

	// search the rom in the db and decode the rom data
	err := r.Db.FindOne(context.TODO(), bson.M{
		"$and": []bson.M{
			{"codename": codename},
			{"generalromdata.androidversion": androidVersion},
			{"generalromdata.romname": romName},
		},
	}).Decode(&rom)
	if err != nil {
		return nil, logger.ErrDbWrite
	}

	r.L.DbRead("readed the data of a rom")

	// return the rom data
	return &rom, nil
}

// get a list of unverified rom
func (r *DbLog) getUnverifiedRomDB(token string) ([]*RomModel, error) {

	// decode the rom list there
	var romsList []*RomModel

	// search the roms in the db
	roms, err := r.Db.Find(context.TODO(), bson.M{"verified": false})
	if err != nil {
		return nil, logger.ErrDbRead
	}

	defer roms.Close(context.TODO())

	// interate every result and add them to the romList slice
	for roms.Next(context.TODO()) {

		var rom RomModel

		// decode the rom data
		if err = roms.Decode(&rom); err != nil {
			log.Fatal(err)
		}

		// add the rom to the rom list
		romsList = append(romsList, &rom)
	}

	// return a list of rom unverified
	return romsList, nil
}

func (r *DbLog) approveRomDB(romId string, token string) error {

	romId = strings.ToLower(romId)

	// convert the rom id in a object id
	id, _ := primitive.ObjectIDFromHex(romId)

	// set true the verified filed
	_, err := r.Db.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		bson.E{"$set", bson.M{"verified": true}},
	})
	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.DbWrite("approved a rom")

	return nil
}

//TODO: aggiungere altri filtri
func (r *DbLog) getRomListDB(codename string, androidVersion float32) ([]*RomModel, error) {

	codename = strings.ToLower(codename)

	// decode the rom list there
	var romsList []*RomModel

	// search the rom in the db
	roms, err := r.Db.Find(context.TODO(), bson.M{
		"$and": []bson.M{
			{"verified": true},
			{"generalromdata.androidversion": androidVersion},
			{"codename": codename},
		},
	})
	if err != nil {
		return nil, logger.ErrDbRead
	}

	r.L.DbRead("readed a list of rom")

	defer roms.Close(context.TODO())

	// interate every result and add them to the romList slice
	for roms.Next(context.TODO()) {

		var rom RomModel

		// decode the rom data
		if err = roms.Decode(&rom); err != nil {
			log.Fatal(err)
		}

		// add the rom data to the rom list
		romsList = append(romsList, &rom)
	}

	// return the list of rom
	return romsList, nil
}

func (r *DbLog) editRomDB(rom *RomModel, token string) error {

	// decode the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// unverify the rom if the user isn't a mod or is verified
	if !tokenData.Moderator && !tokenData.Verified {
		rom.Verified = false
	}

	fmt.Println(tokenData)
	// check if it has uploaded al least one relase of the rom
	var validUser bool = false
	for _, romV := range rom.Version {
		fmt.Println(romV)
		if romV.UploadedBy == tokenData.Username {
			fmt.Println("ciao")
			validUser = true
		}
	}

	fmt.Println(validUser)
	if !validUser {
		return logger.ErrUnauthorized
	}

	// validate the rom data
	err = rom.Validate()
	if err != nil {
		return err
	}

	// replace the rom data
	_, err = r.Db.ReplaceOne(context.TODO(), bson.M{"_id": rom.ID}, rom)
	if err != nil {
		return errors.New("unable to edit the rom info")
	}

	r.L.DbWrite("edited the data of a rom")

	return nil
}
