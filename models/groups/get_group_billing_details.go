package groups

import (
	"time"
)

type GetGroupBillingDetailsRes struct {
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
