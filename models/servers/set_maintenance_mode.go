package servers

type SetMaintenanceModeReq struct {
	Servers []ServerDef
}

type ServerDef struct {
	Id                string
	InMaintenanceMode bool
}
