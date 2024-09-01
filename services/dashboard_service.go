package services

type DashboardService interface {
	GetAdminDashboard()
	GetNodeDashboard()
	GetVendorDashboard()
	GetAdminPackages()
	GetAdminBags()
	GetAdminReturns()
	GetPackageDetails()
	GetBagDetails()
	GetReturnDetails()
}

type dashboardService struct {
}

func NewDashboardService() DashboardService {
	return &dashboardService{}
}

func (ds *dashboardService) GetAdminDashboard() {}

func (ds *dashboardService) GetNodeDashboard() {}

func (ds *dashboardService) GetVendorDashboard() {}

func (ds *dashboardService) GetAdminPackages() {}

func (ds *dashboardService) GetAdminBags() {}

func (ds *dashboardService) GetAdminReturns() {}

func (ds *dashboardService) GetPackageDetails() {}

func (ds *dashboardService) GetBagDetails() {}

func (ds *dashboardService) GetReturnDetails() {}
