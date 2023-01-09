package provisionclient

type DHCPPushStatusMessage struct {
	MSGid       string `json:"msgid,omitempty"`
	Message     string `json:"message,omitempty"`
	State       string `json:"state,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
}
