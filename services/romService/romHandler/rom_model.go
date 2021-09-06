package romhandler

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RomModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RomName        string             `bson:"romname" json:"romname"`
	AndroidVersion float32            `bson:"androidversion" json:"androidversion"`
	Screenshot     []string           `bson:"screenshot" json:"screenshot"`
	Logo           string             `bson:"logo" json:"logo"`
	Description    string             `bson:"description" json:"description"`
	Link           []string           `bson:"link" json:"link"`
	Verified       bool               `bson:"verified" json:"verified"`
	Codename       []string           `bson:"codename" json:"codename"`
	Review         *ReviewModel       `bson:"review" json:"review"`
	Comment        []*CommentModel    `bson:"comment" json:"comment"`
	UploadedBy     string             `bson:"uploadedby" json:"uploadedby"`
}

type FilterRomModel struct {
	RomName        string  `bson:"romname,omitempty" json:"romname,omitempty"`
	AndroidVersion float32 `bson:"androidversion,omitempty" json:"androidversion,omitempty"`
	Verified       bool    `bson:"verified,omitempty" json:"verified,omitempty"`
	Codename       string  `bson:"codename,omitempty" json:"codename,omitempty"`
	OrderBy        string  `bson:"-" json:"orderby,omitempty"`
	Uploadedby     string  `bson:"uploadedby,omitempty" json:"uploadedby,omitempty"`
}
type FilterVersionModel struct {
	Codename   string `bson:"codename,omitempty" json:"codename,omitempty"`
	RomId      string `bson:"romid,omitempty" json:"romid,omitempty"`
	UploadedBy string `bson:"uploadedby,omitempty" json:"uploadedby,omitempty"`
	Verified   bool   `bson:"verified,omitempty" json:"verified,omitempty"`
}

func (rom *RomModel) Validate() error {
	if rom.RomName == "" {
		return errors.New("enter the rom name")
	}
	rom.RomName = strings.ToLower(rom.RomName)
	if rom.AndroidVersion == 0 {
		return errors.New("enter the android version")
	}
	if len(rom.Screenshot) == 0 {
		return errors.New("upload one screenshot")
	}
	if rom.Logo == "" {
		return errors.New("upload the rom logo")
	}
	if rom.Description == "" {
		return errors.New("enter a description for the rom")
	}
	if len(rom.Link) == 0 {
		rom.Link = []string{}
	}

	rom.Codename = []string{}

	rom.Review = &ReviewModel{
		Battery:       0.0,
		Performance:   0.0,
		Stability:     0.0,
		Customization: 0.0,
		ReviewNum:     0,
	}

	rom.Comment = []*CommentModel{}

	return nil
}

type EditRomModel struct {
	Logo        string   `bson:"logo,omitempty" json:"logo,omitempty"`
	Screenshot  []string `bson:"screenshot,omitempty" json:"screenshot,omitempty"`
	Description string   `bson:"description,omitempty" json:"description,omitempty"`
	Link        []string `bson:"link,omitempty" json:"link,omitempty"`
}
type VersionModel struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RomId       string             `bson:"romid" json:"romid"`
	Codename    string             `bson:"codename" json:"codename"`
	Date        time.Time          `bson:"date" json:"date"`
	Official    bool               `bson:"official" json:"official"`
	ChangeLog   []string           `bson:"changelog" json:"changelog"`
	Error       []string           `bson:"error" json:"error"`
	GappsLink   string             `bson:"gappslink" json:"gappslink"`
	VanillaLink string             `bson:"vanillalink" json:"vanillalink"`
	UploadedBy  string             `bson:"uploadedby" json:"uploadedby"`
	RelaseType  string             `bson:"relasetype" json:"relasetype"`
	Verified    bool               `bson:"verified" json:"verified"`
	Version     string             `bson:"version" json:"version"`
}

func (v *VersionModel) Validate() error {
	if v.RomId == "" {
		return errors.New("enter the rom id")
	}

	if v.Codename == "" {
		return errors.New("enter the device codename")
	}

	if v.Version == "" {
		return errors.New("enter the rom version")
	}

	v.Codename = strings.ToLower(v.Codename)

	if len(v.ChangeLog) == 0 {
		v.ChangeLog = []string{}

	}

	if len(v.Error) == 0 {
		v.Error = []string{}
	}

	if v.GappsLink == "" && v.VanillaLink == "" {
		return errors.New("enter a download link")
	}

	if v.RelaseType == "" {
		return errors.New("enter the relase type")
	}

	v.Verified = false

	return nil
}

type RomVersionModel struct {
	Version []*VersionModel `bson:"version" json:"version"`
	Rom     []*RomModel     `bson:"rom" json:"rom"`
}

type CommentModel struct {
	RomId         string  `bson:"romid" json:"romid"`
	Codename      string  `bson:"codename" json:"codename"`
	Username      string  `bson:"username" json:"username"`
	Msg           string  `bson:"msg" json:"msg"`
	Battery       float32 `bson:"battery" json:"battery"`
	Performance   float32 `bson:"performance" json:"performance"`
	Stability     float32 `bson:"stability" json:"stability"`
	Customization float32 `bson:"customization" json:"customization"`
}

func (c *CommentModel) Validate() error {

	if c.Codename == "" {
		return errors.New("enter the device codename")
	}

	if c.RomId == "" {
		return errors.New("invalid rom id")
	}

	if c.Battery < 1 || c.Battery > 5 {
		return errors.New("invalid star range")
	}

	if c.Performance < 1 || c.Performance > 5 {
		return errors.New("invalid star range")
	}

	if c.Stability < 1 || c.Stability > 5 {
		return errors.New("invalid star range")
	}

	if c.Customization < 1 || c.Customization > 5 {
		return errors.New("invalid star range")
	}

	c.Username = ""

	return nil
}

type ReviewModel struct {
	Battery       float32 `bson:"battery" json:"battery"`
	Performance   float32 `bson:"performance" json:"performance"`
	Stability     float32 `bson:"stability" json:"stability"`
	Customization float32 `bson:"customization" json:"customization"`
	ReviewNum     int     `bson:"reviewnum" json:"reviewnum"`
}

type RequestModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Codename       string             `bson:"codename,omitempty" json:"codename,omitempty"`
	AndroidVersion float64            `bson:"androidversion,omitempty" json:"androidversion,omitempty"`
	RomName        string             `bson:"romname,omitempty" json:"romname,omitempty"`
}
