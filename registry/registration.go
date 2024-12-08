package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName // 当前服务依赖的服务列表
	ServiceUpdateURL string        // 服务更新地址
	HeartbeatURL     string        // 服务心跳地址
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	ProtalService  = ServiceName("Protald")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
