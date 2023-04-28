package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getRouteServer() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	// router.Use(static.Serve("/", static.LocalFile("./static", true)))

	// IMPLEMENT THE FOLLOING AS PER AMWA NMOS IS-04 v1.1 SPECIFICIATION
	// https://specs.amwa.tv/is-04/releases/v1.1.3/APIs/NodeAPI.html
	nmos := router.Group("/x-nmos")
	nmos.GET("/", HandleRoot)
	nmosNode := router.Group("/x-nmos/node")
	nmosNode.GET("/", HandleNodeRoot)
	nmosIS04Nodev1_3 := router.Group("/x-nmos/node/v1.3")
	nmosIS04Nodev1_3.GET("/", HandleNodev1_3Root)
	nmosIS04Nodev1_3.GET("", HandleNodev1_3Root)
	nmosIS04Nodev1_3.GET("/self", HandleNodev1_3Self)
	nmosIS04Nodev1_3.GET("/sources", HandleNodev1_3Sources)
	nmosIS04Nodev1_3.GET("/sources/:id", HandleNodev1_3SourcesbyID)
	nmosIS04Nodev1_3.GET("/flows", HandleNodev1_3Flows)
	nmosIS04Nodev1_3.GET("/flows/:id", HandleNodev1_3FlowsbyID) // TODO IMPLEMENT FLOW BY ID
	nmosIS04Nodev1_3.GET("/devices", HandleNodev1_3Devices)
	nmosIS04Nodev1_3.GET("/devices/:id", HandleNodev1_3DevicesbyID)
	nmosIS04Nodev1_3.GET("/senders", HandleNodev1_3Senders)
	nmosIS04Nodev1_3.GET("/senders/:id", HandleNodev1_3SendersbyID)
	nmosIS04Nodev1_3.GET("/receivers", HandleNodev1_3Receivers)
	nmosIS04Nodev1_3.GET("/receivers/:id", HandleNodev1_3ReceiversbyID)              // TODO IMPLEMENT RECEIVER BY ID
	nmosIS04Nodev1_3.PUT("/receivers/:id/target", HandleNodev1_3ReceiversbyIDTarget) // TODO IMPLEMENT RECEIVER BY ID TARGET

	return router
}
