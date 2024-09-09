package routes

import (
	"github.com/SamPariatIL/rrqd/handlers"
	"github.com/SamPariatIL/rrqd/repository"
	"github.com/SamPariatIL/rrqd/services"
	"github.com/SamPariatIL/rrqd/vendors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	db := vendors.Database

	// transactionManager := utils.NewTransactionManager(db)

	dashboardRepository := repository.NewDashboardRepository(db)
	dashboardService := services.NewDashboardService(dashboardRepository)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	dropdownRepository := repository.NewDropdownRepository(db)
	dropdownService := services.NewDropdownService(dropdownRepository)
	dropdownHandler := handlers.NewDropdownHandler(dropdownService)

	packageRepository := repository.NewPackageRepository(db)
	packageService := services.NewPackageService(packageRepository)
	packageHandler := handlers.NewPackageHandler(packageService)

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

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
		dashboard.GET("/admin", dashboardHandler.GetAdminBags)
		dashboard.GET("/node", dashboardHandler.GetNodeDashboard)
		dashboard.GET("/vendor", dashboardHandler.GetVendorDashboard)
		dashboard.GET("/package-details", dashboardHandler.GetPackageDetails)
		dashboard.GET("/bag-details", dashboardHandler.GetBagDetails)
		dashboard.GET("/return-details", dashboardHandler.GetReturnDetails)

		dashboard.POST("/admin-packages", dashboardHandler.GetAdminPackages)
		dashboard.POST("/admin-bags", dashboardHandler.GetAdminBags)
		dashboard.POST("/admin-returns", dashboardHandler.GetAdminReturns)
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
		dropdowns.GET("/states", dropdownHandler.GetStates)
		dropdowns.GET("/freights", dropdownHandler.GetFreights)
		dropdowns.GET("/cities", dropdownHandler.GetCities)
		dropdowns.GET("/pincodes", dropdownHandler.GetPincodes)
		dropdowns.GET("/roles", dropdownHandler.GetRoles)
		dropdowns.GET("/nodes", dropdownHandler.GetNodes)
		dropdowns.GET("/vendors", dropdownHandler.GetVendors)
		dropdowns.GET("/nodes/:nodeId", dropdownHandler.GetNodeSections)
		dropdowns.GET("/node-incharges", dropdownHandler.GetNodeIncharges)
		dropdowns.GET("/node/riders", dropdownHandler.GetNodeRiders)
		dropdowns.GET("/city/:cityId/nodes", dropdownHandler.GetCityNodes)
		dropdowns.GET("/issue-types", dropdownHandler.GetIssueTypes)
		dropdowns.GET("/pincode-details", dropdownHandler.GetPincodeDetails)
		dropdowns.GET("/vendors/:nodeId", dropdownHandler.GetVendorsByNodeId)
		dropdowns.GET("/node-hub", dropdownHandler.GetNodeHubs)
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
		locations.GET("/")
		locations.GET("/:locationId")

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
		packages.GET("/", packageHandler.GetPackages)
		packages.GET("/:packageId", packageHandler.GetPackageById)
		packages.GET("/previous-issue", packageHandler.GetPreviousIssue)

		packages.POST("/", packageHandler.CreatePackage)
		packages.POST("/scan", packageHandler.ScanPackage)
		packages.POST("/create-bag")
		packages.POST("/scan-bag", packageHandler.ScanBag)
		packages.POST("/assign-transport", packageHandler.AssignTransport)
		packages.POST("/create-issue")
		packages.POST("/create-shipments")
		packages.POST("/print", packageHandler.UpdatePrintStatus)
		packages.POST("/incharge/assign-rider", packageHandler.AssignRider)
		packages.POST("/all")
		packages.POST("/bags/all", packageHandler.GetSelectedBags)
		packages.POST("/bulk", packageHandler.CreateBulkReturns)
		packages.POST("/returns", packageHandler.GetReturns)
		packages.POST("/initiate-return", packageHandler.InitiateReturn)
		packages.POST("/bulk-returns", packageHandler.CreateBulkReturns)

		packages.PATCH("/receive", packageHandler.ReceivePackage)
		packages.PATCH("/move", packageHandler.MovePackage)
		packages.PATCH("/receive/bag", packageHandler.ReceiveBag)
		packages.PATCH("/move/bag", packageHandler.MoveBag)
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
		users.GET("/", userHandler.GetUsers)
		users.GET("/:userId", userHandler.GetUserById)

		users.POST("/", userHandler.CreateUser)

		users.PATCH("/:userId", userHandler.UpdateUser)
		users.PATCH("/:userId/:status", userHandler.UpdateUserStatus)
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
