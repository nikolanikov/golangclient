package golangsdk

type Netblock struct {
	ID                  PVID     `json:"id,omitempty"`
	Type                string   `json:"type,omitempty"`
	TopAggregate        PVID     `json:"top_aggregate,omitempty"`
	CIDR                string   `json:"cidr,omitempty"`
	Address             string   `json:"address,omitempty"`
	EndAddress          string   `json:"end_address,omitempty"`
	IsAggregate         bool     `json:"is_aggregate,omitempty"`
	Assigned            bool     `json:"assigned,omitempty"`
	SparseAllocationId  PVID     `json:"sparse_allocation_id,omitempty"`
	IsImportant         bool     `json:"is_important,omitempty"`
	Swipped             bool     `json:"swipped,omitempty"`
	LastUpdateTime      string   `json:"last_update_time,omitempty"`
	LIRID               PVID     `json:"lir_id,omitempty"`
	Mask                int      `json:"mask,omitempty"`
	NetMask             string   `json:"netmask,omitempty"`
	ASN                 PVID     `json:"asn,omitempty"`
	AllowSubAssignments bool     `json:"allow_sub_assignments,omitempty"`
	Child1              PVID     `json:"child1,omitempty"`
	Child2              PVID     `json:"child2,omitempty"`
	ResourceID          PVID     `json:"resource_id,omitempty"`
	ResourceName        string   `json:"resource_name,omitempty"`
	Description         string   `json:"description,omitempty"`
	Parent              PVID     `json:"parent,omitempty"`
	RIR                 string   `json:"rir,omitempty"`
	Notes               string   `json:"notes,omitempty"`
	GenericCode         string   `json:"generic_code,omitempty"`
	AssignTime          string   `json:"assign_time,omitempty"`
	SWIPTime            string   `json:"swip_time,omitempty"`
	NetHandle           string   `json:"net_handle,omitempty"`
	CustomerHandle      string   `json:"customer_handle,omitempty"`
	VLANID              PVID     `json:"vlan_id,omitempty"`
	ORGID               PVID     `json:"org_id,omitempty"`
	Permissions         []string `json:"permissions,omitempty"`
	Region              string   `json:"region,omitempty"`
	RegionID            PVID     `json:"region_id,omitempty"`
	RuleID              PVID     `json:"rule_id,omitempty"`
	ReservedTime        string   `json:"reserved_time,omitempty"`
	ReservedBy          PVID     `json:"reserved_by,omitempty"`
	DHCPResourceID      PVID     `json:"dhcp_resource_id,omitempty"`
	CMNETBLOCKID        PVID     `json:"cmnetblock_resource_id,omitempty"`
	UMBRELLAID          PVID     `json:"umbrella_resource_id,omitempty"`
	Meta1               string   `json:"meta1,omitempty"`
	Meta2               string   `json:"meta2,omitempty"`
	Meta3               string   `json:"meta3,omitempty"`
	Meta4               string   `json:"meta4,omitempty"`
	Meta5               string   `json:"meta5,omitempty"`
	Meta6               string   `json:"meta6,omitempty"`
	Meta7               string   `json:"meta7,omitempty"`
	Meta8               string   `json:"meta8,omitempty"`
	Meta9               string   `json:"meta9,omitempty"`
	Meta10              string   `json:"meta10,omitempty"`
	NAT                 string   `json:"nat,omitempty"`
	HostCount           string   `json:"host_count,omitempty"`
	RegionName          string   `json:"region_name,omitempty"`
	Range               []string `json:"range,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	UtilizationStatus   string   `json:"utilization_status,omitempty"`
}
