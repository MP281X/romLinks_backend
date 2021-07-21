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
	version.Codename = strings.TrimSpace(version.Codename)
	version.Codename = strings.ToLower(version.Codename)

	// insert the version in the db
	id, err := r.DbV.InsertOne(context.TODO(), version)
	if err != nil {
		return "", logger.ErrDbWrite
	}

	// convert the rom id in a object id
	romId, _ := primitive.ObjectIDFromHex(version.RomId)

	// set true the verified filed
	_, err = r.DbR.UpdateOne(context.TODO(), bson.M{"_id": romId}, bson.D{
		{Key: "$addToSet", Value: bson.M{"codename": version.Codename}},
	})
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

	// return a list of rom unverified
	return romsList, nil
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
	if !tokenData.Moderator && !tokenData.Verified {
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

	r.L.DbWrite("approved a rom")

	return nil
}

// get a list of rom
func (r *DbLog) getRomListDB(codename string, androidVersion float32, orderby string) ([]*RomModel, error) {

	codename = strings.ToLower(codename)

	// decode the rom list there
	var romsList []*RomModel

	findOptions := options.Find()
	fmt.Println(orderby)
	if orderby != "/" {
		findOptions.SetSort(bson.M{orderby[1:]: -1})
	}

	// search the rom in the db
	roms, err := r.DbR.Find(context.TODO(), bson.M{
		"$and": []bson.M{
			{"verified": true},
			{"androidversion": androidVersion},
			{"codename": codename},
		},
	}, findOptions)
	if err != nil {
		return nil, logger.ErrDbRead
	}

	r.L.DbRead("readed a list of rom")

	defer roms.Close(context.TODO())

	// interate every result and add them to the romList slice
	if err = roms.All(context.TODO(), &romsList); err != nil {
		return nil, logger.ErrDbRead
	}

	// return the list of rom
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
			{"romid": romId},
			{"codename": codename},
		},
	}, options.Find().SetSort(bson.D{}).SetLimit(20))
	if err != nil {
		return nil, logger.ErrDbRead
	}

	r.L.DbRead("readed a list of version")

	defer versions.Close(context.TODO())

	// interate every result and add them to the versionList slice
	if err = versions.All(context.TODO(), &versionList); err != nil {
		return nil, logger.ErrDbRead
	}

	// return the list of version
	return versionList, nil
}

// get a list of rom name
func (r *DbLog) searchRomNameDB(romName string) ([]string, error) {

	romNameList := []string{}

	// search the rom name in the db
	cursor, err := r.DbR.Find(context.TODO(), bson.M{"$text": bson.M{"$search": romName}}, options.Find().SetSort(bson.D{}).SetLimit(3))
	if err != nil {
		fmt.Println(err.Error())
		return nil, logger.ErrDbRead
	}

	defer cursor.Close(context.TODO())

	// interate every result and add them to the versionList slice
	if err = cursor.All(context.TODO(), &romName); err != nil {
		return nil, logger.ErrDbRead
	}

	// return the rom name list
	return romNameList, nil
}

// add one to the download counter
func (r *DbLog) incrementDownloadDB(romId string, token string) error {

	romId = strings.ToLower(romId)

	// convert the rom id in a object id
	id, _ := primitive.ObjectIDFromHex(romId)

	// check if the user is logged
	_, err := encryption.GetTokenData(token)
	if err != nil {
		return logger.ErrTokenRead
	}

	// increment the download counter
	_, err = r.DbV.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		{Key: "$inc", Value: bson.M{"downloadnumber": 1}},
	})
	if err != nil {
		return logger.ErrDbWrite
	}

	r.L.DbWrite("incremented the download counter")

	return nil
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
	// return a list of rom unverified
	return uploaded, nil
}

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
		{Key: "$push", Value: bson.M{"comment": comment}},
	})
	if err != nil {
		return logger.ErrDbWrite
	}

	// insert the comment in the db
	_, err = r.DbC.InsertOne(context.TODO(), comment)
	if err != nil {
		return logger.ErrDbWrite
	}

	// return a list of rom unverified
	return nil
}

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
