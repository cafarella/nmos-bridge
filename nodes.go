package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoot(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, []string{"node/"})
}

func HandleNodeRoot(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, []string{"v1.3/"})
}

func HandleNodev1_3Root(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, []string{"self/", "sources/", "flows/", "devices/", "senders/", "receivers/"})
}

func HandleNodev1_3Self(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, self)
}

func HandleNodev1_3Sources(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, sources)
}

func HandleNodev1_3SourcesbyID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for _, source := range sources {
		if source.ID == id {
			c.JSON(http.StatusOK, source)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Source not found",
	})
}

func HandleNodev1_3Flows(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, flows)
}

func HandleNodev1_3FlowsbyID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for _, flow := range flows {
		if flow.ID == id {
			c.JSON(http.StatusOK, flow)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Flow not found",
	})
}

func HandleNodev1_3Devices(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	devices := make([]IS041_3Device, 0)
	for _, sender := range senders {
		device := IS041_3Device{
			Description: sender.Description,
			Tags:        sender.Tags,
			Receivers:   []string{},
			Controls:    []IS041_3DeviceControls{},
			Label:       sender.Label,
			Version:     sender.Version,
			Senders:     []string{sender.ID},
			Type:        "urn:x-nmos:device:generic",
			ID:          sender.DeviceID,
			NodeID:      self.ID,
		}
		devices = append(devices, device)
	}
	for _, receiver := range receivers {
		device := IS041_3Device{
			Description: receiver.Description,
			Tags:        receiver.Tags,
			Receivers:   []string{receiver.ID},
			Controls:    []IS041_3DeviceControls{},
			Label:       receiver.Label,
			Version:     receiver.Version,
			Senders:     []string{},
			Type:        "urn:x-nmos:device:generic",
			ID:          receiver.DeviceID,
			NodeID:      self.ID,
		}
		devices = append(devices, device)
	}
	c.JSON(http.StatusOK, devices)
}

func HandleNodev1_3DevicesbyID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for _, sender := range senders {
		if sender.DeviceID == id {
			device := IS041_3Device{
				Description: sender.Description,
				Tags:        sender.Tags,
				Receivers:   []string{},
				Controls:    []IS041_3DeviceControls{},
				Label:       sender.Label,
				Version:     sender.Version,
				Senders:     []string{sender.ID},
				Type:        "urn:x-nmos:device:generic",
				ID:          sender.DeviceID,
				NodeID:      self.ID,
			}
			c.JSON(http.StatusOK, device)
			return
		}
	}
	for _, receiver := range receivers {
		if receiver.DeviceID == id {
			device := IS041_3Device{
				Description: receiver.Description,
				Tags:        receiver.Tags,
				Receivers:   []string{receiver.ID},
				Controls:    []IS041_3DeviceControls{},
				Label:       receiver.Label,
				Version:     receiver.Version,
				Senders:     []string{},
				Type:        "urn:x-nmos:device:generic",
				ID:          receiver.DeviceID,
				NodeID:      self.ID,
			}
			c.JSON(http.StatusOK, device)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Device not found",
	})
}

func HandleNodev1_3Senders(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, senders)
}

func HandleNodev1_3SendersbyID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for _, sender := range senders {
		if sender.ID == id {
			c.JSON(http.StatusOK, sender)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Sender not found",
	})
}

func HandleNodev1_3Receivers(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, receivers)
}

func HandleNodev1_3ReceiversbyID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for _, receiver := range receivers {
		if receiver.ID == id {
			c.JSON(http.StatusOK, receiver)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Receiver not found",
	})
}

func HandleNodev1_3ReceiversbyIDTarget(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	var target Subscription
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, receiver := range receivers {
		if receiver.ID == id {
			if target.Active && target.ReceiverID == nil && target.SenderID == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Receiver must have a sender or receiver specified when activating subscription"})
				return
			}

			receivers[i].Subscription = target
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Receiver not found",
	})
}
