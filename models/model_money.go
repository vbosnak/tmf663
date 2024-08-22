// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Shopping Cart -  Customization
 *
 *  The Shopping Cart API provides a standardized mechanism for the management of shopping carts. Including creation, update, retrieval, deletion and notification of event.  Shopping Cart entity is used for the temporarily selection and reservation of product offerings in e-commerce and retail purchase.  Shopping cart supports purchase of both tangible and intangible goods and service (e.g. handset, telecom network service). The charge includes the one-off fee such as the fee for handset and the recurring fee such as the fee of a network service.  Shopping Cart contains list of cart items, a reference to party or party role (e.g. customer) or contact medium in case of unknown customer, In addition the calculated total items price including promotions.   Copyright  TM Forum 2019. All Rights Reserved   
 *
 * API version: 4.0.0
 */

package openapi




// Money - A base / value business entity used to represent money
type Money struct {

	// Currency (ISO4217 norm uses 3 letters to define the currency)
	Unit string `json:"unit,omitempty"`

	// A positive floating point number
	Value float32 `json:"value,omitempty"`
}

// AssertMoneyRequired checks if the required fields are not zero-ed
func AssertMoneyRequired(obj Money) error {
	return nil
}

// AssertMoneyConstraints checks if the values respects the defined constraints
func AssertMoneyConstraints(obj Money) error {
	return nil
}
