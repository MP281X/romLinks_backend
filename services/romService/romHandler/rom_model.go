package romhandler

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GeneralRomModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	RomName        string             `bson:"romname" json:"romname"`
	AndroidVersion float32            `bson:"androidversion" json:"androidversion"`
	Screenshot     []string           `bson:"screenshot" json:"screenshot"`
	Logo           string             `bson:"logo" json:"logo"`
	Description    string             `bson:"description" json:"description"`
	Link           []string           `bson:"link" json:"link"`
}

func (rom *GeneralRomModel) Validate() error {
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
	return nil
}

type RomModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	GeneralRomData *GeneralRomModel   `bson:"generalromdata" json:"generalromdata"`
	Verified       bool               `bson:"verified" json:"verified"`
	Official       bool               `bson:"official" json:"official"`
	Codename       string             `bson:"codename" json:"codename"`
	Review         *ReviewModel       `bson:"review" json:"review"`
	Version        []*VersionModel    `bson:"version" json:"version"`
}

func (rom *RomModel) Validate() error {
	if rom.GeneralRomData == nil {
		return errors.New("enter the general rom data")
	}
	err := rom.GeneralRomData.Validate()
	if err != nil {
		return err
	}

	if rom.Codename == "" {
		return errors.New("enter the device codename")
	}

	rom.Codename = strings.ToLower(rom.Codename)
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

	if len(rom.Version) == 0 {
		return errors.New("enter a rom relase")
	}

	for _, version := range rom.Version {
		if len(version.Error) == 0 {
			version.Error = []string{}
		}
		if version.GappsLink == "" && version.VanillaLink == "" {
			return errors.New("enter a download link")
		}

	}
	return nil
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

type VersionModel struct {
	Date        time.Time `bson:"date" json:"date"`
	ChangeLog   string    `bson:"changelog" json:"changelog"`
	Error       []string  `bson:"error" json:"error"`
	GappsLink   string    `bson:"gappslink" json:"gappslink"`
	VanillaLink string    `bson:"vanillalink" json:"vanillalink"`
	UploadedBy  string    `bson:"uploadedby" json:"uploadedby"`
}
