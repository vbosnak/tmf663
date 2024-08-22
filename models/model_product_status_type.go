// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Shopping Cart -  Customization
 *
 *  The Shopping Cart API provides a standardized mechanism for the management of shopping carts. Including creation, update, retrieval, deletion and notification of event.  Shopping Cart entity is used for the temporarily selection and reservation of product offerings in e-commerce and retail purchase.  Shopping cart supports purchase of both tangible and intangible goods and service (e.g. handset, telecom network service). The charge includes the one-off fee such as the fee for handset and the recurring fee such as the fee of a network service.  Shopping Cart contains list of cart items, a reference to party or party role (e.g. customer) or contact medium in case of unknown customer, In addition the calculated total items price including promotions.   Copyright  TM Forum 2019. All Rights Reserved   
 *
 * API version: 4.0.0
 */

package openapi


import (
	"fmt"
)


// ProductStatusType : Possible values for the status of the product
type ProductStatusType string

// List of ProductStatusType
const (
	CREATED ProductStatusType = "created"
	PENDING_ACTIVE ProductStatusType = "pendingActive"
	CANCELLED ProductStatusType = "cancelled"
	ACTIVE ProductStatusType = "active"
	PENDING_TERMINATE ProductStatusType = "pendingTerminate"
	TERMINATED ProductStatusType = "terminated"
	SUSPENDED ProductStatusType = "suspended"
	ABORTED ProductStatusType = "aborted "
)

// AllowedProductStatusTypeEnumValues is all the allowed values of ProductStatusType enum
var AllowedProductStatusTypeEnumValues = []ProductStatusType{
	"created",
	"pendingActive",
	"cancelled",
	"active",
	"pendingTerminate",
	"terminated",
	"suspended",
	"aborted ",
}

// validProductStatusTypeEnumValue provides a map of ProductStatusTypes for fast verification of use input
var validProductStatusTypeEnumValues = map[ProductStatusType]struct{}{
	"created": {},
	"pendingActive": {},
	"cancelled": {},
	"active": {},
	"pendingTerminate": {},
	"terminated": {},
	"suspended": {},
	"aborted ": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductStatusType) IsValid() bool {
	_, ok := validProductStatusTypeEnumValues[v]
	return ok
}

// NewProductStatusTypeFromValue returns a pointer to a valid ProductStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductStatusTypeFromValue(v string) (ProductStatusType, error) {
	ev := ProductStatusType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ProductStatusType: valid values are %v", v, AllowedProductStatusTypeEnumValues)
}



// AssertProductStatusTypeRequired checks if the required fields are not zero-ed
func AssertProductStatusTypeRequired(obj ProductStatusType) error {
	return nil
}

// AssertProductStatusTypeConstraints checks if the values respects the defined constraints
func AssertProductStatusTypeConstraints(obj ProductStatusType) error {
	return nil
}
