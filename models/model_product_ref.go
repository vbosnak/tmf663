// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Shopping Cart -  Customization
 *
 *  The Shopping Cart API provides a standardized mechanism for the management of shopping carts. Including creation, update, retrieval, deletion and notification of event.  Shopping Cart entity is used for the temporarily selection and reservation of product offerings in e-commerce and retail purchase.  Shopping cart supports purchase of both tangible and intangible goods and service (e.g. handset, telecom network service). The charge includes the one-off fee such as the fee for handset and the recurring fee such as the fee of a network service.  Shopping Cart contains list of cart items, a reference to party or party role (e.g. customer) or contact medium in case of unknown customer, In addition the calculated total items price including promotions.   Copyright  TM Forum 2019. All Rights Reserved   
 *
 * API version: 4.0.0
 */

package openapi




type ProductRef struct {

	// Unique identifier of a related entity.
	Id string `json:"id"`

	// Reference of the related entity.
	Href string `json:"href,omitempty"`

	// Name of the related entity.
	Name string `json:"name,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type string `json:"@type,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType,omitempty"`
}

// AssertProductRefRequired checks if the required fields are not zero-ed
func AssertProductRefRequired(obj ProductRef) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertProductRefConstraints checks if the values respects the defined constraints
func AssertProductRefConstraints(obj ProductRef) error {
	return nil
}
