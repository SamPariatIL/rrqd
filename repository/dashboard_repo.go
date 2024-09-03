package repository

type DashboardRepository interface {
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

type dashboardRepository struct{}

func NewDashboardRepository() DashboardRepository {
	return &dashboardRepository{}
}

func (dr *dashboardRepository) GetAdminDashboard() {}

func (dr *dashboardRepository) GetNodeDashboard() {}

func (dr *dashboardRepository) GetVendorDashboard() {}

func (dr *dashboardRepository) GetAdminPackages() {}

func (dr *dashboardRepository) GetAdminBags() {}

func (dr *dashboardRepository) GetAdminReturns() {}

func (dr *dashboardRepository) GetPackageDetails() {}

func (dr *dashboardRepository) GetBagDetails() {}

func (dr *dashboardRepository) GetReturnDetails() {}
