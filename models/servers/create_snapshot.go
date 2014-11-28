package servers

type CreateSnapshotReq struct {
	SnapshotExpirationDays int
	ServerIds              []string
}
