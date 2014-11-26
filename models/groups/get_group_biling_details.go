package groups

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
	"time"
)

type GroupBilingDetailsRes struct {
	models.ResModelBase
	Date   time.Time
	Groups []GroupDefinition
}

type GroupDefinition struct {
	Name    string
	Servers []ServerDefinition
}

type ServerDefinition struct {
	TemplateCost    float64
	ArchiveCost     float64
	MonthlyEstimate float64
	MonthToDate     float64
	CurrentHour     float64
}
