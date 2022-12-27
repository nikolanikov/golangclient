package golangclient

type DNSZone struct {
	ID          PVID   `json:"id"`
	ParentID    PVID   `json:"parent_id"`
	GroupID     PVID   `json:"group_id,omitempty"`
	Name        string `json:"name"`
	ZoneType    string `json:"zone_type,omitempty"`
	Status      string `json:"status,omitempty"`
	ZoneExpire  int    `json:"zone_expire,string,omitempty"`
	ZoneHost    string `json:"zone_host,omitempty"`
	ZoneMail    string `json:"zone_mail,omitempty"`
	ZoneMinimum int    `json:"zone_minimum,string,omitempty"`
	ZoneRefresh int    `json:"zone_refresh,string,omitempty"`
	ZoneRetry   int    `json:"zone_retry,string,omitempty"`
	ZoneSerial  int    `json:"zone_serial,string,omitempty"`
	ZoneTTL     int    `json:"zone_ttl,string,omitempty"`
}

type DNSRecord struct {
	ID          PVID   `json:"id"`
	ParentID    PVID   `json:"parent_id"`
	Name        string `json:"name,omitempty"`
	Modified    string `json:"modified,omitempty"`
	Status      string `json:"status,omitempty"`
	RecordType  string `json:"record_type"`
	RecordHost  string `json:"record_host"`
	RecordValue string `json:"record_value"`
	RecordTTL   int    `json:"record_ttl,string"`
}

type DNSGroup struct {
	ID          PVID   `json:"id"`
	ParentID    PVID   `json:"parent_id"`
	Name        string `json:"name,omitempty"`
	Modified    string `json:"modified,omitempty"`
	Status      string `json:"status,omitempty"`
	RecordType  string `json:"record_type"`
	RecordHost  string `json:"record_host"`
	RecordValue string `json:"record_value"`
	RecordTTL   int    `json:"record_ttl,string"`
}
