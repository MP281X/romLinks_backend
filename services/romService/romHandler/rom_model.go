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
	Official       bool               `bson:"official" json:"official"`
	Codename       []string           `bson:"codename" json:"codename"`
	Review         *ReviewModel       `bson:"review" json:"review"`
	UploadedBy     string             `bson:"uploadedby" json:"uploadedby"`
}

type ReviewModel struct {
	Battery             float32 `bson:"battery" json:"battery"`
	BatteryRevNum       int     `bson:"batteryrevnum" json:"batteryrevnum"`
	Performance         float32 `bson:"performance" json:"performance"`
	PerformanceRevNum   int     `bson:"performancerevnum" json:"performancerevnum"`
	Stability           float32 `bson:"stability" json:"stability"`
	StabilityRevNum     int     `bson:"stabilityrevnum" json:"stabilityrevnum"`
	Customization       float32 `bson:"customization" json:"customization"`
	CustomizationRevNum int     `bson:"customizationrevnum" json:"customizationrevnum"`
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

	if len(rom.Codename) == 0 {
		rom.Codename = []string{}
	}

	for i, codename := range rom.Codename {
		rom.Codename[i] = strings.ToLower(codename)
	}

	rom.Review = &ReviewModel{
		Battery:             0.0,
		BatteryRevNum:       0,
		Performance:         0.0,
		PerformanceRevNum:   0,
		Stability:           0.0,
		StabilityRevNum:     0,
		Customization:       0.0,
		CustomizationRevNum: 0,
	}
	return nil
}

type VersionModel struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RomId          string             `bson:"romid" json:"romid"`
	Codename       string             `bson:"codename" json:"codename"`
	Date           time.Time          `bson:"date" json:"date"`
	ChangeLog      []string           `bson:"changelog" json:"changelog"`
	Error          []string           `bson:"error" json:"error"`
	GappsLink      string             `bson:"gappslink" json:"gappslink"`
	VanillaLink    string             `bson:"vanillalink" json:"vanillalink"`
	UploadedBy     string             `bson:"uploadedby" json:"-"`
	DownloadNumber int                `bson:"downloadnumber" json:"downloadnumber"`
	RelaseType     string             `bson:"relasetype" json:"relasetype"`
}

func (v *VersionModel) Validate() error {
	if v.RomId == "" {
		return errors.New("enter the rom id")
	}

	if v.Codename == "" {
		return errors.New("enter the device codename")
	}

	v.Codename = strings.ToLower(v.Codename)

	v.Date = time.Now()

	if len(v.ChangeLog) == 0 {
		v.ChangeLog = []string{}

	}

	if len(v.Error) == 0 {
		v.Error = []string{}
	}

	if v.GappsLink == "" && v.VanillaLink == "" {
		return errors.New("enter a download link")
	}

	v.DownloadNumber = 0

	if v.RelaseType == "" {
		return errors.New("enter the relase type")
	}

	return nil
}

type RomVersionModel struct {
	Version []*VersionModel `bson:"version" json:"version"`
	Rom     []*RomModel     `bson:"rom" json:"rom"`
}
