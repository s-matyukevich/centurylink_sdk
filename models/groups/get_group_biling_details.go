package groups

import (
	"time"
)

type GetGroupBilingDetailsRes struct {
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
