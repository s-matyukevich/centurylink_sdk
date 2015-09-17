package groups

import (
	"time"
	"github.com/s-matyukevich/centurylink_sdk/base"
)

type GetGroupMonitoringStatisticsRes struct {
	// Injected field; to traverse links
	Connection	base.Connection

	Name  string
	Stats []StatsDef
}

type StatsDef struct {
	Timestamp                time.Time
	Cpu                      float64
	CpuPercent               float64
	MemoryMB                 float64
	MemoryPercent            float64
	NetworkReceivedKbps      float64
	NetworkTransmittedKbps   float64
	DiskUsageTotalCapacityMB float64
	DiskUsage                []DiskUsageDef
	GuestDiskUsage           []GuestUsageDef
}

type DiskUsageDef struct {
	Id         string
	capacityMB int
}

type GuestUsageDef struct {
	Path       string
	CapacityMB int
	ConsumedMB int
}
