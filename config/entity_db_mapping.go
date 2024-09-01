package config

import (
	"reflect"

	"github.com/SamPariatIL/rrqd/entities"
)

const (
	COLLECTION_BAGS          = "bags"
	COLLECTION_CITIES        = "cities"
	COLLECTION_COMPLAINTS    = "complaints"
	COLLECTION_FREIGHT_TYPES = "freightTypes"
	COLLECTION_ISSUE_TYPES   = "issueTypes"
	COLLECTION_ISSUES        = "issues"
	COLLECTION_LOCATIONS     = "locations"
	COLLECTION_LOCKS         = "locks"
	COLLECTION_NODES         = "nodes"
	COLLECTION_OTP           = "otp"
	COLLECTION_PACKAGES      = "packages"
	COLLECTION_PAYMENTS      = "payments"
	COLLECTION_PINCODE       = "pincode"
	COLLECTION_PRINCIPALS    = "principals"
	COLLECTION_RETURNS       = "returns"
	COLLECTION_RIDERS        = "riders"
	COLLECTION_ROLES         = "roles"
	COLLECTION_ROUTES        = "routes"
	COLLECTION_SCAN_HISTORY  = "scanHistory"
	COLLECTION_SHIPMENTS     = "shipments"
	COLLECTION_STATES        = "states"
	COLLECTION_USERS         = "users"
	COLLECTION_VENDORS       = "vendors"
)

type EntityConfig struct {
	CollectionNames map[reflect.Type]string
}

func NewEntityConfig() *EntityConfig {
	return &EntityConfig{
		CollectionNames: map[reflect.Type]string{
			reflect.TypeOf(entities.Bag{}):         COLLECTION_BAGS,
			reflect.TypeOf(entities.City{}):        COLLECTION_CITIES,
			reflect.TypeOf(entities.Complaint{}):   COLLECTION_COMPLAINTS,
			reflect.TypeOf(entities.Freight{}):     COLLECTION_FREIGHT_TYPES,
			reflect.TypeOf(entities.Issue{}):       COLLECTION_ISSUE_TYPES,
			reflect.TypeOf(entities.Location{}):    COLLECTION_LOCATIONS,
			reflect.TypeOf(entities.Node{}):        COLLECTION_NODES,
			reflect.TypeOf(entities.OTP{}):         COLLECTION_OTP,
			reflect.TypeOf(entities.Package{}):     COLLECTION_PACKAGES,
			reflect.TypeOf(entities.Payment{}):     COLLECTION_PAYMENTS,
			reflect.TypeOf(entities.Pincode{}):     COLLECTION_PINCODE,
			reflect.TypeOf(entities.Principal{}):   COLLECTION_PRINCIPALS,
			reflect.TypeOf(entities.Return{}):      COLLECTION_RETURNS,
			reflect.TypeOf(entities.Rider{}):       COLLECTION_RIDERS,
			reflect.TypeOf(entities.Role{}):        COLLECTION_ROLES,
			reflect.TypeOf(entities.Route{}):       COLLECTION_ROUTES,
			reflect.TypeOf(entities.ScanHistory{}): COLLECTION_SCAN_HISTORY,
			reflect.TypeOf(entities.Shipment{}):    COLLECTION_SHIPMENTS,
			reflect.TypeOf(entities.State{}):       COLLECTION_STATES,
			reflect.TypeOf(entities.User{}):        COLLECTION_USERS,
			reflect.TypeOf(entities.Vendor{}):      COLLECTION_VENDORS,
		},
	}
}

func (ec *EntityConfig) GetCollectionName(entity any) string {
	entityType := reflect.TypeOf(entity)
	return ec.CollectionNames[entityType]
}
