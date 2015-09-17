package groups

import (
	"time"
	"github.com/s-matyukevich/centurylink_sdk/base"
)

type GetGroupBillingDetailsRes struct {
	// Injected field; to traverse links
	Connection	base.Connection

	Date   time.Time
	Groups map[string]GroupDef
}

type GroupDef struct {
	Name    string
	Servers map[string]ServerDef
}

type ServerDef struct {
	TemplateCost    float64
	ArchiveCost     float64
	MonthlyEstimate float64
	MonthToDate     float64
	CurrentHour     float64
}
