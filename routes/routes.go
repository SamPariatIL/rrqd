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

	locations := router.Group("/locations")
	{
		locations.GET("/:locationId")
		locations.GET("/")

		locations.POST("/")

		locations.PATCH("/:locationId")
		locations.PATCH("/:locationId:/status")
	}
}
