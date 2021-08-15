package romhandler

import (
	"context"
	"fmt"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// add the rom name to the rom name slice
	r.RN.AddValue(rom.RomName)

	// get the rom id
	userId := fmt.Sprintf("%v", id.InsertedID)
	userId = userId[10 : len(userId)-2]

	r.L.Info(tokenData.Username + " added a new rom: " + rom.RomName)

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
	version.Codename = strings.TrimSpace(version.Codename)
	version.Codename = strings.ToLower(version.Codename)

	if tokenData.Verified || tokenData.Moderator {
		version.Verified = true
	}

	// insert the version in the db
	id, err := r.DbV.InsertOne(context.TODO(), version)
	if err != nil {
		return "", logger.ErrDbWrite
	}

	// convert the rom id in a object id
	romId, _ := primitive.ObjectIDFromHex(version.RomId)

	// add the codename of the device to the codename list in the rom
	_, err = r.DbR.UpdateOne(context.TODO(), bson.M{"_id": romId}, bson.D{
		{Key: "$addToSet", Value: bson.M{"codename": version.Codename}},
	})
	if err != nil {
		return "", logger.ErrDbWrite
	}

	// get the rom id
	userId := fmt.Sprintf("%v", id.InsertedID)
	userId = userId[10 : len(userId)-2]

	r.L.Info(tokenData.Username + " added a new version for " + version.RomId)

	return userId, nil
}

// get a list of unverified rom
func (r *DbLog) getUnverifiedRomDB(token string) ([]*RomModel, error) {

	// decode the rom list there
	var romsList []*RomModel

	// get the data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return nil, logger.ErrTokenRead
	}

	// check if the user is a moderator
	if !tokenData.Moderator {
		return nil, logger.ErrUnauthorized
	}

	// search the roms in the db
	roms, err := r.DbR.Find(context.TODO(), bson.M{"verified": false}, options.Find().SetSort(bson.D{}).SetLimit(20))
	if err != nil {
		return nil, logger.ErrDbRead
	}

	defer roms.Close(context.TODO())
	if err = roms.All(context.TODO(), &romsList); err != nil {
		return nil, logger.ErrDbRead
	}

	r.L.Info(tokenData.Username + " searched for unverifyed rom")

	// return a list of unverified rom
	return romsList, nil
}

// get a list of unverified version
func (r *DbLog) getUnverifiedVersionDB(token string) ([]*VersionModel, error) {

	// decode the version list there
	var versionList []*VersionModel

	// get the data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return nil, logger.ErrTokenRead
	}

	// check if the user is a moderator
	if !tokenData.Moderator {
		return nil, logger.ErrUnauthorized
	}

	// search the versions in the db
	versions, err := r.DbV.Find(context.TODO(), bson.M{"verified": false}, options.Find().SetSort(bson.D{}).SetLimit(20))
	if err != nil {
		return nil, logger.ErrDbRead
	}

	defer versions.Close(context.TODO())
	if err = versions.All(context.TODO(), &versionList); err != nil {
		return nil, logger.ErrDbRead
	}

	r.L.Info(tokenData.Username + " searched for unverifyed version")

	// return a list of unverified version
	return versionList, nil
}

// approve a rom
func (r *DbLog) approveRomDB(romId string, token string) error {

	romId = strings.ToLower(romId)

	// check if the token is valid
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// check if the user has the permission
	if !tokenData.Moderator {
		return logger.ErrUnauthorized
	}

	// convert the rom id in a object id
	id, _ := primitive.ObjectIDFromHex(romId)

	// set true the verified filed
	_, err = r.DbR.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		{Key: "$set", Value: bson.M{"verified": true}},
	})
	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.Info(tokenData.Username + " approved a rom: " + romId)

	return nil
}

// approve a rom
func (r *DbLog) approveVersionDB(versionId string, token string) error {

	versionId = strings.ToLower(versionId)

	// check if the token is valid
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// check if the user has the permission
	if !tokenData.Moderator {
		return logger.ErrUnauthorized
	}

	// convert the rom id in a object id
	id, _ := primitive.ObjectIDFromHex(versionId)

	// set true the verified filed
	_, err = r.DbV.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		{Key: "$set", Value: bson.M{"verified": true}},
	})
	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.Info(tokenData.Username + " approved a version: " + versionId)

	return nil
}

// get a list of rom
func (r *DbLog) getRomListDB(filter *FilterRomModel) ([]*RomModel, error) {

	filter.Codename = strings.ToLower(filter.Codename)

	// decode the rom list there
	var romsList []*RomModel

	filter.Verified = true

	findOptions := options.Find()
	if filter.OrderBy != "" {
		findOptions.SetSort(bson.M{filter.OrderBy: -1})
	}

	var roms *mongo.Cursor
	var err error

	// search the rom in the db
	roms, err = r.DbR.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, logger.ErrDbRead
	}

	defer roms.Close(context.TODO())

	// interate every result and add them to the romList slice
	if err = roms.All(context.TODO(), &romsList); err != nil {
		return nil, logger.ErrDbRead
	}

	// return the list of rom
	return romsList, nil
}

// get a list of rom by a list of id
func (r *DbLog) getRomByIdDB(rom []string) ([]*RomModel, error) {

	// decode the rom list there
	var romsList []*RomModel

	// convert the list of string in a list of object id
	var romId []primitive.ObjectID = []primitive.ObjectID{}
	for _, id := range rom {
		x, _ := primitive.ObjectIDFromHex(id)
		romId = append(romId, x)
	}

	// get the list of rom from the db
	roms, err := r.DbR.Find(context.TODO(), bson.D{{Key: "_id", Value: bson.M{"$in": romId}}}, options.Find().SetSort(bson.D{}).SetLimit(20))
	if err != nil {
		return nil, logger.ErrDbRead
	}

	// interate every result and add them to the romList slice
	if err = roms.All(context.TODO(), &romsList); err != nil {
		return nil, logger.ErrDbRead
	}

	return romsList, nil
}

// get a list of version
func (r *DbLog) getVersionListDB(codename string, romId string) ([]*VersionModel, error) {

	codename = strings.ToLower(codename)
	// decode the version list there
	var versionList []*VersionModel

	// search the version in the db
	versions, err := r.DbV.Find(context.TODO(), bson.M{
		"$and": []bson.M{
			{"verified": true},
			{"romid": romId},
			{"codename": codename},
		},
	}, options.Find().SetSort(bson.D{}).SetLimit(20))
	if err != nil {
		return nil, logger.ErrDbRead
	}

	defer versions.Close(context.TODO())

	// interate every result and add them to the versionList slice
	if err = versions.All(context.TODO(), &versionList); err != nil {
		return nil, logger.ErrDbRead
	}

	// return the list of version
	return versionList, nil
}

// get a list of rom and version uploaded by the user
func (r *DbLog) getUploadedDB(token string) (*RomVersionModel, error) {

	// decode the rom list there
	var romsList []*RomModel
	var versionList []*VersionModel

	// get the data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return nil, logger.ErrTokenRead
	}

	// search the roms in the db
	roms, err := r.DbR.Find(context.TODO(), bson.M{"uploadedby": tokenData.Username}, options.Find().SetSort(bson.D{}))
	if err != nil {
		return nil, logger.ErrDbRead
	}

	defer roms.Close(context.TODO())
	if err = roms.All(context.TODO(), &romsList); err != nil {
		return nil, logger.ErrDbRead
	}

	// search the versions in the db
	versions, err := r.DbV.Find(context.Background(), bson.M{"uploadedby": tokenData.Username}, options.Find().SetSort(bson.D{}))
	if err != nil {
		return nil, logger.ErrDbRead
	}

	defer versions.Close(context.TODO())
	if err = versions.All(context.TODO(), &versionList); err != nil {
		return nil, logger.ErrDbRead
	}

	// combine the version and the rom list
	var uploaded *RomVersionModel = &RomVersionModel{
		Rom:     romsList,
		Version: versionList,
	}

	r.L.Info(tokenData.Username + " searched all the rom and version who has uploaded")
	// return a list of rom unverified
	return uploaded, nil
}

// add a review for a rom
func (r *DbLog) addReviewDB(token string, comment *CommentModel) error {

	// validate the comment data
	err := comment.Validate()
	if err != nil {
		return err
	}

	// get the data from the token
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return logger.ErrTokenRead
	}

	comment.Username = tokenData.Username

	// convert the rom id in a object id
	id, _ := primitive.ObjectIDFromHex(comment.RomId)

	// add the review to the db
	_, err = r.DbR.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		{Key: "$inc", Value: bson.M{"review.battery": comment.Battery}},
		{Key: "$inc", Value: bson.M{"review.performance": comment.Performance}},
		{Key: "$inc", Value: bson.M{"review.stability": comment.Stability}},
		{Key: "$inc", Value: bson.M{"review.customization": comment.Customization}},
		{Key: "$inc", Value: bson.M{"review.reviewnum": 1}},
	})
	if err != nil {
		return logger.ErrDbWrite
	}

	// insert the comment in the db
	_, err = r.DbC.InsertOne(context.TODO(), comment)
	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.Info(tokenData.Username + " added a review for " + comment.RomId)

	// return a list of rom unverified
	return nil
}

// get a list of review for a rom
func (r *DbLog) getReviewDB(romId string) ([]*CommentModel, error) {
	var commentList []*CommentModel

	// search the rom name in the db
	cursor, err := r.DbC.Find(context.TODO(), bson.M{"romid": romId}, options.Find().SetSort(bson.D{}).SetLimit(25))
	if err != nil {
		fmt.Println(err.Error())
		return nil, logger.ErrDbRead
	}

	// interate every result and add them to the comment list
	if err = cursor.All(context.TODO(), &commentList); err != nil {
		return nil, logger.ErrDbRead
	}

	// return the comment list
	return commentList, nil
}

// edit the data of a rom
func (r *DbLog) editRomDataDB(romData *EditRomModel, token string, romId string) error {

	romId = strings.ToLower(romId)
	var data RomModel

	// convert the rom id in a object id
	id, _ := primitive.ObjectIDFromHex(romId)

	// check if the token is valid
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// get the data of the rom to edit
	err = r.DbR.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
	if err != nil {
		return logger.ErrDbRead
	}

	// check if the user has the permission
	if !tokenData.Moderator && tokenData.Username != data.UploadedBy {
		return logger.ErrUnauthorized
	}

	// set true the verified filed
	_, err = r.DbR.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		{Key: "$set", Value: romData},
	})

	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.Info(tokenData.Username + " edited the data of a rom: " + romId)

	return nil
}

// edit the data of a version
func (r *DbLog) editVersionDataDB(versionData *EditVersionModel, token string, versionId string) error {

	versionId = strings.ToLower(versionId)
	var data VersionModel

	// convert the rom id in a object id
	id, _ := primitive.ObjectIDFromHex(versionId)

	// check if the token is valid
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// get the data of the rom to edit
	err = r.DbV.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
	if err != nil {
		return logger.ErrDbRead
	}

	// check if the user has the permission
	if !tokenData.Moderator && tokenData.Username != data.UploadedBy {
		return logger.ErrUnauthorized
	}

	// set true the verified filed
	_, err = r.DbV.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		{Key: "$set", Value: versionData},
	})

	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.Info(tokenData.Username + " edited the data of a version: " + versionId)

	return nil
}

// delete a version
func (r *DbLog) removeVersionDB(token string, versionId string) error {

	versionId = strings.ToLower(versionId)
	var data VersionModel

	// convert the version id in a object id
	id, _ := primitive.ObjectIDFromHex(versionId)

	// check if the token is valid
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// get the data of the version to delete
	err = r.DbV.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
	if err != nil {
		return logger.ErrDbRead
	}

	// check if the user has the permission
	if !tokenData.Moderator && tokenData.Username != data.UploadedBy {
		return logger.ErrUnauthorized
	}

	// delete the version
	_, err = r.DbV.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.Info(tokenData.Username + " deleted a version: " + versionId)

	return nil
}

// delete a rom
func (r *DbLog) removeRomDB(token string, romId string) error {

	romId = strings.ToLower(romId)
	var data RomModel

	// convert the rom id in a object id
	id, _ := primitive.ObjectIDFromHex(romId)

	// check if the token is valid
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	// get the data of the rom to delete
	err = r.DbR.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
	if err != nil {
		return logger.ErrDbRead
	}

	// check if the user has the permission
	if !tokenData.Moderator && tokenData.Username != data.UploadedBy {
		return logger.ErrUnauthorized
	}

	// delete the rom
	_, err = r.DbR.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		return logger.ErrDbWrite
	}

	// remove the rom name from the rom name slice
	r.RN.RemoveValue(data.RomName)

	r.L.Info(tokenData.Username + " deleted a rom: " + romId)

	return nil
}

// get a list of romName
func (r *DbLog) GetRomName() error {

	// search the rom in the db
	cursor, err := r.DbR.Find(context.TODO(), bson.M{}, nil)
	if err != nil {
		return logger.ErrDbRead
	}

	defer cursor.Close(context.TODO())

	// add the rom name to the rom name slice
	for cursor.Next(context.TODO()) {
		var val bson.M

		if err = cursor.Decode(&val); err != nil {
			return logger.ErrDbRead
		}
		r.RN.AddValue(val["romname"].(string))
	}

	return nil
}
