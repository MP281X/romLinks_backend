package romhandler

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// add a new rom to the db
func (r *DbLog) addRomDB(rom *RomModel, token string) (string, error) {

	// validate the input data
	err := rom.Validate()
	if err != nil {
		return "", err
	}

	// get the data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return "", logger.ErrTokenRead
	}

	// set the uploader name
	rom.UploadedBy = tokenData.Username

	// if the user is a mod or is verified the rom is pubblic without additional validation
	if tokenData.Moderator || tokenData.Verified {
		rom.Verified = true
	}

	// insert the rom in the db
	id, err := r.DbR.InsertOne(context.TODO(), rom)
	if err != nil {
		return "", logger.ErrDbWrite
	}

	// get the rom id
	userId := fmt.Sprintf("%v", id.InsertedID)
	userId = userId[10 : len(userId)-2]

	r.L.DbWrite("added a new rom")

	return userId, nil
}

// add a new version to the db
func (r *DbLog) addVersionDB(version *VersionModel, token string) (string, error) {

	// validate the input data
	err := version.Validate()
	if err != nil {
		return "", err
	}

	// get the data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return "", logger.ErrTokenRead
	}

	// set the uploader name
	version.UploadedBy = tokenData.Username

	// insert the version in the db
	id, err := r.DbV.InsertOne(context.TODO(), version)
	if err != nil {
		return "", logger.ErrDbWrite
	}

	// get the rom id
	userId := fmt.Sprintf("%v", id.InsertedID)
	userId = userId[10 : len(userId)-2]

	r.L.DbWrite("added a new version")

	return userId, nil
}

// get the data of a rom
func (r *DbLog) getRomDB(codename string, androidVersion float32, romName string) (*RomModel, error) {

	codename = strings.ToLower(codename)
	romName = strings.ToLower(romName)
	// decode the rom data there
	var rom RomModel

	// search the rom in the db and decode the rom data
	err := r.DbR.FindOne(context.TODO(), bson.M{
		"$and": []bson.M{
			{"codename": codename},
			{"androidversion": androidVersion},
			{"romname": romName},
		},
	}).Decode(&rom)
	if err != nil {
		return nil, logger.ErrDbRead
	}

	r.L.DbRead("readed the data of a rom")

	// return the rom data
	return &rom, nil
}

// get the data of a rom
func (r *DbLog) getRomByIdDB(id string) (*RomModel, error) {

	// decode the rom data there
	var rom RomModel

	romId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	// search the rom in the db and decode the rom data
	err = r.DbR.FindOne(context.TODO(), bson.M{"_id": romId}).Decode(&rom)
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
	roms, err := r.DbR.Find(context.TODO(), bson.M{"verified": false})
	if err != nil {
		return nil, logger.ErrDbRead
	}

	defer roms.Close(context.TODO())

	// interate every result and add them to the romList slice
	for roms.Next(context.TODO()) {

		var rom RomModel

		// decode the rom data
		if err = roms.Decode(&rom); err != nil {
			return nil, logger.ErrDbRead
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
	_, err := r.DbR.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		bson.E{"$set", bson.M{"verified": true}},
	})
	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.DbWrite("approved a rom")

	return nil
}

// //TODO: aggiungere altri filtri
func (r *DbLog) getRomListDB(codename string, androidVersion float32) ([]*RomModel, error) {

	codename = strings.ToLower(codename)

	// decode the rom list there
	var romsList []*RomModel

	// search the rom in the db
	roms, err := r.DbR.Find(context.TODO(), bson.M{
		"$and": []bson.M{
			{"verified": true},
			{"androidversion": androidVersion},
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
			return nil, logger.ErrDbRead
		}

		// add the rom data to the rom list
		romsList = append(romsList, &rom)
	}

	// return the list of rom
	return romsList, nil
}

func (r *DbLog) getVersionListDB(codename string, romId string) ([]*VersionModel, error) {

	codename = strings.ToLower(codename)
	// decode the version list there
	var versionList []*VersionModel

	// search the version in the db
	versions, err := r.DbV.Find(context.TODO(), bson.M{
		"$and": []bson.M{
			{"romid": romId},
			{"codename": codename},
		},
	})
	if err != nil {
		return nil, logger.ErrDbRead
	}

	r.L.DbRead("readed a list of version")

	defer versions.Close(context.TODO())

	// interate every result and add them to the versionList slice
	for versions.Next(context.TODO()) {

		var version VersionModel

		// decode the version data
		if err = versions.Decode(&version); err != nil {
			return nil, logger.ErrDbRead
		}

		// add the version data to the version list
		versionList = append(versionList, &version)
	}

	// return the list of version
	return versionList, nil
}
