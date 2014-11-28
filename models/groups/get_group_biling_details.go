package groups

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
	"time"
)

type GroupBilingDetailsRes struct {
	models.ResModelBase
	Date   time.Time
	Groups []GroupDef
}

type GroupDef struct {
	Name    string
	Servers []ServerDef
}

type ServerDef struct {
	TemplateCost    float64
	ArchiveCost     float64
	MonthlyEstimate float64
	MonthToDate     float64
	CurrentHour     float64
}
