package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	dashboard := router.Group("/api/v1/dashboards")
	{
		dashboard.GET("/admin")
		dashboard.GET("/node")
		dashboard.GET("/vendor")
		dashboard.GET("/package-details")
		dashboard.GET("/bag-details")
		dashboard.GET("/return-details")

		dashboard.POST("/admin-packages")
		dashboard.POST("/admin-bags")
		dashboard.POST("/admin-returns")
	}

	dropdowns := router.Group("/api/v1/dropdowns")
	{
		dropdowns.GET("/states")
		dropdowns.GET("/freights")
		dropdowns.GET("/cities")
		dropdowns.GET("/pincodes")
		dropdowns.GET("/roles")
		dropdowns.GET("/nodes")
		dropdowns.GET("/vendors")
		dropdowns.GET("/nodes/:nodeId")
		dropdowns.GET("/node-in-charges")
		dropdowns.GET("/node/riders")
		dropdowns.GET("/city/:cityId/nodes")
		dropdowns.GET("/issue-types")
		dropdowns.GET("/pincode-details")
		dropdowns.GET("/vendors-by-node-id")
		dropdowns.GET("/node-hub")
	}

	issues := router.Group("/api/v1/issues")
	{
		issues.GET("/:issueId")
		issues.GET("/issues")

		issues.POST("/")
		issues.POST("/new")

		issues.PATCH("/:issueId/:status")
	}

	locations := router.Group("/api/v1/locations")
	{
		locations.GET("/:locationId")
		locations.GET("/")

		locations.POST("/")

		locations.PATCH("/:locationId")
		locations.PATCH("/:locationId/:status")
	}

	routes := router.Group("/api/v1/routes")
	{
		routes.GET("/:routeId")
		routes.GET("/total-data")
		routes.GET("/city-node")
		routes.GET("/dark-node")
		routes.GET("/node-list")

		routes.POST("/")
		routes.POST("/get-paginated-data")

		routes.PATCH("/:routeId")
		routes.PATCH("/:routeId/:status")
	}

	nodes := router.Group("/api/v1/nodes")
	{
		nodes.GET("/:nodeId")
		nodes.GET("/total-data")
		nodes.GET("/incharge/bags")
		nodes.GET("/incharge/packages")
		nodes.GET("/incharge/issued-shipments")
		nodes.GET("/users")
		nodes.GET("/incharge/bag/:bagId")
		nodes.GET("/incharge/package/:packageId")
		nodes.GET("/:nodeId/users")
		nodes.GET("/incharge/node")
		nodes.GET("/payments/riders/:riderId")
		nodes.GET("/payments/vendors/:vendorId")
		nodes.GET("/payments")
		nodes.GET("/return-details")

		nodes.POST("/")
		nodes.POST("/get-paginated-data")
		nodes.POST("/complaints")
		nodes.POST("/returns")
		nodes.POST("/get-multiple-returns")
		nodes.POST("/raise-complaint")

		nodes.PATCH("/:nodeId")
		nodes.PATCH("/:nodeId/:status")
	}
}
