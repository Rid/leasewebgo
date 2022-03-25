package leasewebgo

const dedicatedServerBasePath = "/bareMetals/v2/servers"

// DeviceServiceOp implements DeviceService
type DedicatedServerServiceOp struct {
	client *Client
}

// DedicatedServerService interface defines available device methods
type DedicatedServerService interface {
	List(opts *ListOptions) (*dedicatedServersRoot, *Response, []byte, error)
	// Get(DeviceID string, opts *GetOptions) (*Device, *Response, error)
	// Create(*DeviceCreateRequest) (*Device, *Response, error)
	// Update(string, *DeviceUpdateRequest) (*Device, *Response, error)
	// Reboot(string) (*Response, error)
	// HardwareInformation(string) (*HardwareInformation, *Response, error)

	// // IPs
	// ListIps(string) ([]IPAddress, *Response, error)
	// ShowIp(string, string) (*IPAddress, *Response, error)
	// UpdateIp(string, string, *IPAddressUpdateRequest) (*IPAddress, *Response, error)
	// NullRouteIp(string, string) (*Response, error)
	// RemoteNullRouteIp(string, string) (*Response, error)
	// ShowNullRouteIp(string, string) (*NullRoute, *Response, error)

	// // Network Interfaces
	// ListNetworkInterfaces(string) ([]NetworkInterface, *Response, error)
	// CloseAllNetworkInterfaces(string) (*Response, error)
	// OpenAllNetworkInterfaces(string) (*Response, error)
	// ShowNetworkInterface(string, string) (*NetworkInterface, *Response, error)
	// CloseNetworkInterface(string, string) (*Response, error)
	// OpenNetworkInterface(string, string) (*Response, error)

	// // Private Networks
	// DeleteServerFromPrivateNetwork(string, string) (*Response, error)
	// AddServerToPrivateNetwork(string, string, string) (*Response, error)

	// // DHCP Leases
	// DeleteDhcpReservation(string, string) (*Response, error)
	// ListDhcpReservations(string) ([]DhcpReservation, *Response, error)
	// CreateDhcpReservation(string, *DhcpReservationCreateRequest) (*DhcpReservation, *Response, error)

	// // Jobs
	// CancelActiveJob(string) (*Response, error)
	// ExpireActiveJob(string) (*Response, error)
	// LaunchHardwareScan(string) (*Response, error)
	// LaunchInstallation(string) (*Response, error)
	// LaunchIpmiReset(string) (*Response, error)
	// ListJobs(string) ([]Job, *Response, error)
	// ShowJob(string, string) (*Job, *Response, error)
	// LaunchRescueMode(string) (*Response, error)

	// // Credentials
	// ListCredentials(string) ([]Credential, *Response, error)
	// CreateNewCredentials(string, *CredentialCreateRequest) (*Credential, *Response, error)
	// ListCredentialsByType(string, string) ([]Credential, *Response, error)
	// DeleteUserCredentials(string, string) (*Response, error)
	// ShowUserCredentials(string, string) (*Credential, *Response, error)
	// UpdateUserCredentials(string, string, *CredentialUpdateRequest) (*Credential, *Response, error)

	// // Metrics
	// ShowBandwidthMetrics(string, *BandwidthOpts) (*Bandwidth, *Response, error)
	// ShowDatatrafficMetrics(string, *BandwidthOpts) (*Datatraffic, *Response, error)

	// // Notification Settings
	// ListBandwidthNotificationSettings(string) ([]NotificationSetting, *Response, error)
	// CreateBandwidthNotificationSetting(string, *NotificationSettingCreateRequest) (*NotificationSetting, *Response, error)
	// DeleteBandwidthNotificationSetting(string, string) (*Response, error)
	// ShowBandwidthNotificationSetting(string, string) (*NotificationSetting, *Response, error)
	// UpdateBandwidthNotificationSetting(string, string, *NotificationSettingUpdateRequest) (*NotificationSetting, *Response, error)
	// ListDatatrafficNotificationSettings(string) ([]NotificationSetting, *Response, error)
	// CreateDatatrafficNotificationSetting(string, *NotificationSettingCreateRequest) (*NotificationSetting, *Response, error)
	// ShowDatatrafficNotificationSetting(string, string) (*NotificationSetting, *Response, error)
	// UpdateDatatrafficNotificationSetting(string, string, *NotificationSettingUpdateRequest) (*NotificationSetting, *Response, error)
	// InspectDDoSNotificationSettings(string) (*DDoSNotificationSettings, *Response, error)
	// UpdateDDoSNotificationSettings(string, *DDoSNotificationSettingsUpdateRequest) (*DDoSNotificationSettings, *Response, error)

	// // Power
	// PowerCycleServer(string) (*Response, error)
	// ShowPowerStatus(string) (*PowerStatus, *Response, error)
	// PowerOffServer(string) (*Response, error)
	// PowerOnServer(string) (*Response, error)

	// // Operating System
	// ListOperatingSystems(string) ([]OperatingSystem, *Response, error)
	// ShowOperatingSystem(string, string) (*OperatingSystem, *Response, error)
	// ListControlPanels(string) ([]ControlPanel, *Response, error)

	// // Control Panels
	// ListControlPanels(string) ([]ControlPanel, *Response, error)

	// // Rescue Images
	// RescueImages() ([]RescueImage, *Response, error)
}

type dedicatedServersRoot struct {
	DedicatedServer []DedicatedServer `json:"servers"`
	Meta            meta              `json:"_metadata"`
}

// DedicatedServer represents a Dedicated Server from the Leaseweb API
type DedicatedServer struct {
	AssetId             string               `json:"assetId"`
	Contract            *Contract            `json:"contract,omitempty"`
	FeatureAvailability *FeatureAvailability `json:"featureAvailability,omitempty"`
	Id                  string               `json:"id,omitempty"`
	Location            *Location            `json:"location,omitempty"`
	NetworkInterfaces   *NetworkInterfaces   `json:"networkInterfaces,omitempty"`
	PowerPorts          []PowerPorts         `json:"powerPorts,omitempty"`
	PrivateNetworks     []PrivateNetworks    `json:"privateNetworks,omitempty"`
	Rack                *Rack                `json:"rack,omitempty"`
}

type Contract struct { //nolint:golint
	CustomerId     string `json:"customerId"`
	DeliveryStatus string `json:"deliveryStatus"`
	Id             string `json:"id"`
	Reference      string `json:"reference"`
	SalesOrgId     string `json:"salesOrgId"`
}

type FeatureAvailability struct { //nolint:golint
	Automation       bool `json:"automation"`
	IpmiReboot       bool `json:"ipmiReboot"`
	PowerCycle       bool `json:"powerCycle"`
	PrivateNetwork   bool `json:"privateNetwork"`
	RemoteManagement bool `json:"remoteManagement"`
}

type Location struct { //nolint:golint
	Rack  string `json:"rack"`
	Site  string `json:"site"`
	Suite string `json:"suite"`
	Unit  string `json:"unit"`
}

type NetworkInterfaces struct { //nolint:golint
	Internal         *Internal         `json:"internal,omitempty"`
	Public           *Public           `json:"public,omitempty"`
	RemoteManagement *RemoteManagement `json:"remoteManagement,omitempty"`
}

type Internal struct { //nolint:golint
	Gateway string  `json:"gateway"`
	Ip      string  `json:"ip"`
	Mac     string  `json:"mac"`
	Ports   []Ports `json:"ports"`
}

type Ports struct { //nolint:golint
	Name string `json:"name"`
	Port string `json:"port"`
}

type Public struct { //nolint:golint
	Gateway string  `json:"gateway"`
	Ip      string  `json:"ip"`
	Mac     string  `json:"mac"`
	Ports   []Ports `json:"ports"`
}

type RemoteManagement struct { //nolint:golint
	Gateway string  `json:"gateway"`
	Ip      string  `json:"ip"`
	Mac     string  `json:"mac"`
	Ports   []Ports `json:"ports"`
}

type PowerPorts struct { //nolint:golint
	Name string `json:"name"`
	Port string `json:"port"`
}

type PrivateNetworks struct { //nolint:golint
	Id        string `json:"id"`
	LinkSpeed int    `json:"linkSpeed"`
	Status    string `json:"status"`
	VlanId    string `json:"vlanId"`
}

type Rack struct { //nolint:golint
	Type string `json:"type"`
}

func (d *DedicatedServerServiceOp) List(opts *ListOptions) (*dedicatedServersRoot, *Response, []byte, error) {
	endpointPath := dedicatedServerBasePath
	apiPathQuery := opts.WithQuery(endpointPath)
	root := new(dedicatedServersRoot)
	resp, body, err := d.client.DoRequest("GET", apiPathQuery, nil, root)

	if err != nil {
		return nil, resp, body, err
	}
	return root, resp, body, nil
}
