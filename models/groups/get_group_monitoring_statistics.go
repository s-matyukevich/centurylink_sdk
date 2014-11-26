package groups

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
	"time"
)

type GroupMonitoringStatisticsRes struct {
	models.ResModelBase
	Name  string
	Stats []StatsDefinition
}

type StatsDefinition struct {
	Timestamp                time.Time
	Cpu                      float64
	CpuPercent               float64
	MemoryMB                 float64
	MemoryPercent            float64
	NetworkReceivedKbps      float64
	NetworkTransmittedKbps   float64
	DiskUsageTotalCapacityMB float64
	DiskUsage                []DiskUsageDefinition
	GuestDiskUsage           []GuestUsageDefinition
}

type DiskUsageDefinition struct {
	Id         string
	capacityMB int
}

type GuestUsageDefinition struct {
	Path       string
	CapacityMB int
	ConsumedMB int
}
