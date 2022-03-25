package leasewebgo

import (
	"path"
)

const floatingIpsBasePath = "/floatingIps/v2/ranges"

// DeviceServiceOp implements DeviceService
type FloatingIpsServiceOp struct {
	client *Client
}

// DedicatedServerService interface defines available device methods
type FloatingIpsService interface {
	CreateRange(rangeId string, rangeCreateRequest *RangeCreateRequest, opts *ListOptions) (*RangeDefinition, *Response, []byte, error)
	UpdateRange(rangeId string, floatingIpDefinitionId string, rangeUpdateRequest *RangeUpdateRequest, opts *ListOptions) (*RangeDefinition, *Response, []byte, error)
	ListRanges(opts *ListOptions) (*RangesRoot, *Response, []byte, error)
	ListRangeDefinitions(rangeId string, opts *ListOptions) (*rangeDefinitionsRoot, *Response, []byte, error)
}

type RangesRoot struct {
	Ranges []Range `json:"ranges"`
	Meta   meta    `json:"_metadata"`
}

type rangeDefinitionsRoot struct {
	RangeDefinitions []RangeDefinition `json:"floatingIpDefinitions,omitempty"`
	Meta             meta              `json:"_metadata,omitempty"`
}

type Range struct {
	Id          string `json:"id,omitempty"`
	Range       string `json:"range,omitempty"`
	CustomerId  string `json:"customerId,omitempty"`
	SalesOrgId  string `json:"salesOrgId,omitempty"`
	Pop         string `json:"pop,omitempty"`
	Location    string `json:"location,omitempty"`
	Type        string `json:"type,omitempty"`
	EquipmentId string `json:"equipmentId,omitempty"`
}

type RangeDefinition struct {
	Id         string `json:"id,omitempty"`
	RangeId    string `json:"rangeId,omitempty"`
	Location   string `json:"location,omitempty"`
	Type       string `json:"type,omitempty"`
	CustomerId string `json:"customerId,omitempty"`
	SalesOrgId string `json:"salesOrgId,omitempty"`
	FloatingIp string `json:"floatingIp,omitempty"`
	AnchorIp   string `json:"anchorIp,omitempty"`
	Status     string `json:"status,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	UpdatedAt  string `json:"updatedAt,omitempty"`
}

type RangeCreateRequest struct {
	FloatingIp string `json:"floatingIp,omitempty"`
	AnchorIp   string `json:"anchorIp,omitempty"`
}

type RangeUpdateRequest struct {
	AnchorIp string `json:"anchorIp,omitempty"`
}

func (d *FloatingIpsServiceOp) CreateRange(rangeId string, rangeCreateRequest *RangeCreateRequest, opts *ListOptions) (*RangeDefinition, *Response, []byte, error) {
	endpointPath := path.Join(floatingIpsBasePath, "/", rangeId, "/", "floatingIpDefinitions")
	apiPathQuery := opts.WithQuery(endpointPath)
	rangeDefinition := new(RangeDefinition)
	resp, body, err := d.client.DoRequest("POST", apiPathQuery, rangeCreateRequest, rangeDefinition)

	if err != nil {
		return nil, resp, body, err
	}
	return rangeDefinition, resp, body, nil
}

func (d *FloatingIpsServiceOp) UpdateRange(rangeId string, floatingIpDefinitionId string, rangeUpdateRequest *RangeUpdateRequest, opts *ListOptions) (*RangeDefinition, *Response, []byte, error) {
	endpointPath := path.Join(floatingIpsBasePath, "/", rangeId, "/", "floatingIpDefinitions", "/", floatingIpDefinitionId)
	apiPathQuery := opts.WithQuery(endpointPath)
	rangeDefinition := new(RangeDefinition)
	resp, body, err := d.client.DoRequest("PUT", apiPathQuery, rangeUpdateRequest, rangeDefinition)

	if err != nil {
		return nil, resp, body, err
	}
	return rangeDefinition, resp, body, nil
}

func (f *FloatingIpsServiceOp) ListRanges(opts *ListOptions) (*RangesRoot, *Response, []byte, error) {
	endpointPath := floatingIpsBasePath
	apiPathQuery := opts.WithQuery(endpointPath)
	root := new(RangesRoot)
	resp, body, err := f.client.DoRequest("GET", apiPathQuery, nil, root)

	if err != nil {
		return nil, resp, body, err
	}
	return root, resp, body, nil
}

func (f *FloatingIpsServiceOp) ListRangeDefinitions(rangeId string, opts *ListOptions) (*rangeDefinitionsRoot, *Response, []byte, error) {
	endpointPath := path.Join(floatingIpsBasePath, "/", rangeId, "/", "floatingIpDefinitions")
	apiPathQuery := opts.WithQuery(endpointPath)
	root := new(rangeDefinitionsRoot)
	resp, body, err := f.client.DoRequest("GET", apiPathQuery, nil, root)

	if err != nil {
		return nil, resp, body, err
	}
	return root, resp, body, nil
}
