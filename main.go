package main

import (
	"fmt"

	"github.com/google/uuid"
)

var (
	self      IS041_3Node
	sources   []IS041_3Source
	senders   []IS041_3Sender
	receivers []IS041_3Receiver
	flows     []IS041_3Flow
)

func main() {
	sources = make([]IS041_3Source, 0)
	senders = make([]IS041_3Sender, 0)
	receivers = make([]IS041_3Receiver, 0)
	flows = make([]IS041_3Flow, 0)

	self = IS041_3Node{
		Description: "Implemented for Learning Purposes",
		API: IS041_3NodeAPI{
			Endpoints: []IS041_3NodeAPIEndpoint{
				{
					Host:     "localhost",
					Port:     8080,
					Protocol: "http",
				},
			},
			Versions: []string{"v1.3"},
		},
		Caps: IS041_3NodeCaps{},
		Clocks: []IS041_3NodeClocks{
			{
				Name:    "clk0",
				RefType: "internal",
			},
		},
		Hostname: "nmos-implementation-workshop.local",
		Href:     "http://localhost:8080",
		ID:       "d70a4137-97f9-48f1-8d19-831f422b1aa1",
		Interfaces: []IS041_3NodeInterface{
			{
				Name:      "eth0",
				PortID:    "74-26-96-db-87-31",
				ChassisID: "74-26-96-db-87-31",
			},
		},
		Label: "NMOS Implementation Workshop Device",
		Services: []IS041_3NodeServices{
			{
				Authorization: false,
				Href:          "http://localhost:8080/x-nmos/query/v1.3/",
				Type:          "urn:x-nmos:service:query",
			},
		},
		Tags:    Tag{},
		Version: "1441700172:318426300",
	}

	sources = []IS041_3Source{
		{
			ID:          "d70a4137-97f9-48f1-8d19-831f422b1aa4",
			DeviceID:    "d70a4137-97f9-48f1-8d19-831f422b1aa1",
			Description: "Example source",
			Format:      "urn:x-nmos:format:video",
			Label:       "Source 1",
			Version:     "1441700172:318426300",
			ClockName:   "clk0",
			Caps:        IS041_3SourceCap{},
			Parents: []interface{}{
				"d70a4137-97f9-48f1-8d19-831f422b1aa1",
			},
		},
	}
	senders = []IS041_3Sender{
		{
			ID:           "d70a4137-97f9-48f1-8d19-831f422b1aa1",
			DeviceID:     "d70a4137-97f9-48f1-8d19-831f422b1aa1",
			Description:  "Example sender",
			Tags:         Tag{},
			Label:        "Sender 1",
			Version:      "1539094460:303320729",
			ManifestHref: "http://localhost:8081/x-nmos/senders/sender1/transportfile.sdp",
			FlowID:       "d70a4137-97f9-48f1-8d19-831f422b1aa2",
			Transport:    "urn:x-nmos:transport:rtp.mcast",
			InterfaceBindings: []string{
				"eth0",
			},
			Subscription: Subscription{
				Active:     true,
				ReceiverID: "d70a4137-97f9-48f1-8d19-831f422b1aa3",
			},
		},
	}

	getRouteServer().Run(":8080")
}

func AddSender(index int) {
	id := uuid.NewString()
	senders = append(senders, IS041_3Sender{
		Description:  fmt.Sprintf("Video output %d on NMOS Test Implementation", index),
		Tags:         Tag{},
		Label:        fmt.Sprintf("VIDEO-OUT-%d", index),
		Version:      "1539094460:303320729",
		ManifestHref: fmt.Sprintf("http://localhost:8081/x-nmos/senders/%s/transportfile.sdp", id),
		FlowID:       id,
		Subscription: Subscription{
			Active:     true,
			ReceiverID: nil,
			SenderID:   "",
		},
	})
}
