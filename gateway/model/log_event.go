package model

type LogAuditTrailEvent struct {
	ID        string
	AppName   string `json:"app_name"`
	IpAddress string `json:"ip_address"`
	Path      string `json:"path"`
	Action    string `json:"action"`
	Original  any    `json:"original"`
	Changes   any    `json:"changes"`
	Token     string `json:"token"`
}

func (l *LogAuditTrailEvent) GetID() string {
	return l.ID
}
