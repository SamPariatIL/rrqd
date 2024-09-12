package cmd

import "github.com/SamPariatIL/rrqd/vendors"

func Execute() {
	initVendors()
}

func initVendors() {
	vendors.VendorSetup()
}
