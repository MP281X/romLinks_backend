package romhandler

import (
	"context"
	"testing"

	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
)

func TestValidation(t *testing.T) {

	// null input data
	x := &RomModel{
		GeneralRomData: &GeneralRomModel{
			RomName:        "",
			AndroidVersion: 1,
		},
	}

	err := x.Validate()
	if err == nil {
		t.Error("the input was null")
	}

	// correct input data
	romData := &RomModel{
		GeneralRomData: &GeneralRomModel{
			RomName:        "testRomName",
			AndroidVersion: 11,
			Screenshot:     []string{"test"},
			Logo:           "testLinkLogo",
			Description:    "test description",
		},
		Codename: "testCodename",
		Version: []*VersionModel{
			&VersionModel{
				ChangeLog:   "improved ...",
				VanillaLink: "vanilla download link",
			},
			&VersionModel{
				ChangeLog:   "improved 2 ...",
				VanillaLink: "vanilla download link 2",
			},
		},
	}

	err = romData.Validate()
	if err != nil {
		t.Error("the input data was correct")
	}
}

func TestDBReq(t *testing.T) {

	// initialize the logger and the db
	d, _ := db.InitDB("test")
	l, _ := logger.InitLogger("test")
	c := d.Collection("test_rom")
	r := &DbLog{Db: c, L: l}

	// clear the test collection
	c.Drop(context.TODO())

	// generate a fake jwt
	token, _ := encryption.GenerateJwt("test", &encryption.TokenData{Verified: false, Moderator: false, Username: "mp281x"})

	// rom model
	romData := &RomModel{
		GeneralRomData: &GeneralRomModel{
			RomName:        "testRomName",
			AndroidVersion: 11,
			Screenshot:     []string{"test"},
			Logo:           "testLinkLogo",
			Description:    "test description",
		},
		Codename: "testCodename",
		Version: []*VersionModel{
			&VersionModel{
				ChangeLog:   "improved ...",
				VanillaLink: "vanilla download link",
			},
			&VersionModel{
				ChangeLog:   "improved 2 ...",
				VanillaLink: "vanilla download link 2",
			},
		},
	}

	// add the rom
	err := r.addRomDB(romData, token)
	if err != nil {
		t.Error(err)
	}

	// get the rom
	_, err = r.getRomDB(romData.Codename, romData.GeneralRomData.AndroidVersion, romData.GeneralRomData.RomName)
	if err != nil {
		t.Error(err)
	}

	// generate a fake jwt
	token2, _ := encryption.GenerateJwt("test", &encryption.TokenData{Verified: true, Moderator: true, Username: "mp281x"})

	// get the list of unverified rom
	unvRom, err := r.getUnverifiedRomDB(token2)
	if err != nil {
		t.Error(err)
	}

	if unvRom[0].GeneralRomData.RomName != "testromname" {
		t.Error("the rom data has changed")
	}

	// approve the rom
	err = r.approveRomDB(unvRom[0].ID.Hex(), token2)
	if err != nil {
		t.Error(err)
	}

	// get the list of verified rom
	romList, err := r.getRomListDB(unvRom[0].Codename, unvRom[0].GeneralRomData.AndroidVersion)
	if err != nil {
		t.Error(err)
	}

	if romList[0].GeneralRomData.RomName != "testromname" {
		t.Error("the rom data has changed")
	}

	// clear the test collection
	c.Drop(context.TODO())

}
