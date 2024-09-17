# golangclient
GoLang client for 6connect ProVision

In case that you require more functionality for 3rd party projects, please contact us.

## Example Usage Golang

```golang
client, err := provisionclient.NewClient("https://url.to.6connect.com", "example.user@example.com", "password", true)
if err != nil {
  fmt.Println("ERROR")
}

netblocks, err := client.IPAM.GetNetblocks(&map[string]string{
			"cidr": "192.168.192.176/28",
		})
    
record, err := client.DNS.AddZoneRecord(provisionclient.DNSRecord{
		ParentID:    "428964",
		Name:        "Golang TXT Record",
		RecordType:  "TXT",
		RecordHost:  "golang.example.com.",
		RecordValue: "Created from GOLANGCLIENT",
		RecordTTL:   3600,
	})
```
