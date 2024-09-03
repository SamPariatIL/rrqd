package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	auth := router.Group("/api/v1/auth")
	{
		auth.GET("/refresh-token")
		auth.GET("/verify-otp")
		auth.GET("/warehouse/verify-otp")
		auth.GET("/delivery/verify-otp")

		auth.POST("/login")
		auth.POST("/change-password")
		auth.POST("/send-otp")
		auth.POST("/reset-password")
		auth.POST("/warehouse/send-otp")
		auth.POST("/delivery/send-otp")
		auth.POST("/principals/login")
		auth.POST("/principals/reset-password")
		auth.POST("/principals/change-password")
	}

	dashboard := router.Group("/api/v1/dashboard")
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

	deliveries := router.Group("/api/v1/deliveries")
	{
		deliveries.GET("/order")
		deliveries.GET("/payment-details")
		deliveries.GET("/history/rider")
		deliveries.GET("/mid-mile")
		deliveries.GET("/mid-mile-history")
		deliveries.GET("/send-return-otp")

		deliveries.POST("/scan-delivery")
		deliveries.POST("/failed-delivery")
		deliveries.POST("/initiate-return")
		deliveries.POST("/create-return-shipment")

		deliveries.PATCH("/update-status")
		deliveries.PATCH("/update-multiple-status")
		deliveries.PATCH("/verify-delivery")
		deliveries.PATCH("/collect-payment")
		deliveries.PATCH("/pick-up-bag")
		deliveries.PATCH("/drop-bag")
		deliveries.PATCH("/freight-details")
		deliveries.PATCH("/assign-transport")
		deliveries.PATCH("/verify-return-otp")
		deliveries.PATCH("/drop-return-package")
		deliveries.PATCH("/pick-principal-return")
		deliveries.PATCH("/drop-principal-return")
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

	packages := router.Group("/api/v1/packages")
	{
		packages.GET("/")
		packages.GET("/:packageId")
		packages.GET("/previous-issue")

		packages.POST("/")
		packages.POST("/scan")
		packages.POST("/create-bag")
		packages.POST("/scan-bag")
		packages.POST("/assign-transport")
		packages.POST("/create-issue")
		packages.POST("/create-shipments")
		packages.POST("/print")
		packages.POST("/incharge/assign-rider")
		packages.POST("/all")
		packages.POST("/bags/all")
		packages.POST("/bulk")
		packages.POST("/returns")
		packages.POST("/initiate-return")
		packages.POST("/bulk-returns")

		packages.PATCH("/receive")
		packages.PATCH("/move")
		packages.PATCH("/receive/bag")
		packages.PATCH("/move/bag")
	}

	payments := router.Group("/api/v1/payments")
	{
		payments.GET("/")
		payments.GET("/orders")
		payments.GET("/riders")
		payments.GET("/vendor-payment")
		payments.GET("/rider-payment")
		payments.GET("/shipment-payment")

		payments.POST("/generate-link")
	}

	principals := router.Group("/api/v1/principals")
	{
		principals.GET("/")
		principals.GET("/:principalId")
		principals.GET("/locations")

		principals.POST("/")

		principals.PATCH("/:principalId")
		principals.PATCH("/:principalId/:status")
	}

	riders := router.Group("/api/v1/riders")
	{
		riders.GET("/")
		riders.GET("/:riderId")

		riders.POST("/")

		riders.PATCH("/:riderId")
		riders.PATCH("/:riderId/:status")
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

	users := router.Group("/api/v1/users")
	{
		users.GET("/")
		users.GET("/:userId")

		users.POST("/")

		users.PATCH("/:userId")
		users.PATCH("/:userId/:status")
	}

	vendors := router.Group("/api/v1/vendors")
	{
		vendors.GET("/")
		vendors.GET("/:vendorId")
		vendors.GET("/users")
		vendors.GET("/shipment-details")
		vendors.GET("/bags")
		vendors.GET("/bag-details")
		vendors.GET("/rider-shipments")

		vendors.POST("/")
		vendors.POST("/bulk-serviceable-areas")
		vendors.POST("/packages")
		vendors.POST("/riders")
		vendors.POST("/payments")
		vendors.POST("/issues")

		vendors.PATCH("/:vendorId")
		vendors.PATCH("/:vendorId/:status")
	}
}
