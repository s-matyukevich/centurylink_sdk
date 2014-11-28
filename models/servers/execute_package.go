package servers

type ExecutePackageReq struct {
	Servers []string
	Package PackageDef
}
