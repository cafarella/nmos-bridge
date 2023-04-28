package main

type IS041_3Node struct {
	Description string                 `json:"description"`
	API         IS041_3NodeAPI         `json:"api"`
	Caps        IS041_3NodeCaps        `json:"caps"`
	Clocks      []IS041_3NodeClocks    `json:"clocks"`
	Hostname    string                 `json:"hostname"`
	Href        string                 `json:"href"`
	ID          string                 `json:"id"`
	Interfaces  []IS041_3NodeInterface `json:"interfaces"`
	Label       string                 `json:"label"`
	Services    []IS041_3NodeServices  `json:"services"`
	Tags        Tag                    `json:"tags"`
	Version     string                 `json:"version"`
}

type IS041_3NodeAPIEndpoint struct {
	Host          string `json:"host"`
	Protocol      string `json:"protocol"`
	Port          int    `json:"port"`
	Authorization bool   `json:"authorization"`
}

type IS041_3NodeAPI struct {
	Endpoints []IS041_3NodeAPIEndpoint `json:"endpoints"`
	Versions  []string                 `json:"versions"`
}

type IS041_3NodeInterfacesAttachedNetworkDevice struct {
	ChassisID string `json:"chassis_id,omitempty"`
	PortID    string `json:"port_id,omitempty"`
}

type IS041_3NodeInterface struct {
	PortID                string                                      `json:"port_id"`
	Name                  string                                      `json:"name"`
	ChassisID             interface{}                                 `json:"chassis_id"`
	AttachedNetworkDevice *IS041_3NodeInterfacesAttachedNetworkDevice `json:"attached_network_device,omitempty"`
}

type IS041_3NodeClocks struct {
	Gmid      string `json:"gmid,omitempty"`
	Locked    bool   `json:"locked,omitempty"`
	Name      string `json:"name"`
	RefType   string `json:"ref_type"`
	Traceable bool   `json:"traceable,omitempty"`
	Version   string `json:"version,omitempty"`
}

type IS041_3NodeServices struct {
	Authorization bool   `json:"authorization"`
	Href          string `json:"href"`
	Type          string `json:"type"`
}

type IS041_3NodeCaps struct {
}

type IS041_3Source struct {
	Description string           `json:"description"`
	Caps        IS041_3SourceCap `json:"caps"`
	Format      string           `json:"format"`
	Tags        Tag              `json:"tags"`
	Label       string           `json:"label"`
	Version     string           `json:"version"`
	Parents     []interface{}    `json:"parents"`
	ClockName   string           `json:"clock_name"`
	ID          string           `json:"id"`
	DeviceID    string           `json:"device_id"`
}
type Tag struct {
	Host []string `json:"host,omitempty"`
}

type IS041_3SourceCap struct {
}

type IS041_3Device struct {
	Description string                  `json:"description"`
	Tags        Tag                     `json:"tags"`
	Receivers   []string                `json:"receivers"`
	Controls    []IS041_3DeviceControls `json:"controls"`
	Label       string                  `json:"label"`
	Version     string                  `json:"version"`
	Senders     []string                `json:"senders"`
	Type        string                  `json:"type"`
	ID          string                  `json:"id"`
	NodeID      string                  `json:"node_id"`
}

type IS041_3DeviceControls struct {
	Href string `json:"href"`
	Type string `json:"type"`
}

type IS041_3Flow struct {
	Colorspace             string                 `json:"colorspace"`
	Description            string                 `json:"description"`
	Format                 string                 `json:"format"`
	FrameHeight            int                    `json:"frame_height"`
	Tags                   Tag                    `json:"tags"`
	TransferCharacteristic string                 `json:"transfer_characteristic"`
	FrameWidth             int                    `json:"frame_width"`
	Label                  string                 `json:"label"`
	GrainRate              IS041_3FlowGrainRate   `json:"grain_rate"`
	Version                string                 `json:"version"`
	Parents                []interface{}          `json:"parents"`
	InterlaceMode          string                 `json:"interlace_mode"`
	Components             []IS041_3FlowComponent `json:"components"`
	SourceID               string                 `json:"source_id"`
	MediaType              string                 `json:"media_type"`
	ID                     string                 `json:"id"`
	DeviceID               string                 `json:"device_id"`
}

type IS041_3FlowGrainRate struct {
	Denominator int `json:"denominator"`
	Numerator   int `json:"numerator"`
}

type IS041_3FlowComponent struct {
	Width    int    `json:"width"`
	BitDepth int    `json:"bit_depth"`
	Name     string `json:"name"`
	Height   int    `json:"height"`
}

type IS041_3Receiver struct {
	Description       string              `json:"description"`
	Format            string              `json:"format"`
	Tags              Tag                 `json:"tags"`
	Label             string              `json:"label"`
	Subscription      Subscription        `json:"subscription"`
	Version           string              `json:"version"`
	Caps              IS041_3ReceiverCaps `json:"caps"`
	InterfaceBindings []string            `json:"interface_bindings"`
	ID                string              `json:"id"`
	Transport         string              `json:"transport"`
	DeviceID          string              `json:"device_id"`
}

type IS041_3ReceiverCaps struct {
	MediaTypes []string `json:"media_types"`
}

type IS041_3Sender struct {
	Description       string       `json:"description"`
	Tags              Tag          `json:"tags"`
	Label             string       `json:"label"`
	Version           string       `json:"version"`
	ManifestHref      string       `json:"manifest_href"`
	FlowID            string       `json:"flow_id"`
	Subscription      Subscription `json:"subscription"`
	InterfaceBindings []string     `json:"interface_bindings"`
	ID                string       `json:"id"`
	Transport         string       `json:"transport"`
	DeviceID          string       `json:"device_id"`
}

type Subscription struct {
	Active     bool        `json:"active"`
	ReceiverID interface{} `json:"receiver_id,omitempty"`
	SenderID   interface{} `json:"sender_id,omitempty"`
}

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Debug string `json:"debug"`
}
