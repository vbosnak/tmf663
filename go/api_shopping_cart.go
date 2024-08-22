// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Shopping Cart -  Customization
 *
 *  The Shopping Cart API provides a standardized mechanism for the management of shopping carts. Including creation, update, retrieval, deletion and notification of event.  Shopping Cart entity is used for the temporarily selection and reservation of product offerings in e-commerce and retail purchase.  Shopping cart supports purchase of both tangible and intangible goods and service (e.g. handset, telecom network service). The charge includes the one-off fee such as the fee for handset and the recurring fee such as the fee of a network service.  Shopping Cart contains list of cart items, a reference to party or party role (e.g. customer) or contact medium in case of unknown customer, In addition the calculated total items price including promotions.   Copyright  TM Forum 2019. All Rights Reserved   
 *
 * API version: 4.0.0
 */

package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	models "tmf663/models"
)

// ShoppingCartAPIController binds http requests to an api service and writes the service results to the http response
type ShoppingCartAPIController struct {
	service ShoppingCartAPIServicer
	errorHandler models.ErrorHandler
}

// ShoppingCartAPIOption for how the controller is set up.
type ShoppingCartAPIOption func(*ShoppingCartAPIController)

// WithShoppingCartAPIErrorHandler inject ErrorHandler into controller
func WithShoppingCartAPIErrorHandler(h models.ErrorHandler) ShoppingCartAPIOption {
	return func(c *ShoppingCartAPIController) {
		c.errorHandler = h
	}
}

// NewShoppingCartAPIController creates a default api controller
func NewShoppingCartAPIController(s ShoppingCartAPIServicer, opts ...ShoppingCartAPIOption) *ShoppingCartAPIController {
	controller := &ShoppingCartAPIController{
		service:      s,
		errorHandler: models.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ShoppingCartAPIController
func (c *ShoppingCartAPIController) Routes() Routes {
	return Routes{
		"CreateShoppingCart": Route{
			strings.ToUpper("Post"),
			"/tmf-api/shoppingCart/v4/shoppingCart",
			c.CreateShoppingCart,
		},
		"DeleteShoppingCart": Route{
			strings.ToUpper("Delete"),
			"/tmf-api/shoppingCart/v4/shoppingCart/{id}",
			c.DeleteShoppingCart,
		},
		"ListShoppingCart": Route{
			strings.ToUpper("Get"),
			"/tmf-api/shoppingCart/v4/shoppingCart",
			c.ListShoppingCart,
		},
		"PatchShoppingCart": Route{
			strings.ToUpper("Patch"),
			"/tmf-api/shoppingCart/v4/shoppingCart/{id}",
			c.PatchShoppingCart,
		},
		"RetrieveShoppingCart": Route{
			strings.ToUpper("Get"),
			"/tmf-api/shoppingCart/v4/shoppingCart/{id}",
			c.RetrieveShoppingCart,
		},
	}
}

// CreateShoppingCart - Creates a ShoppingCart
func (c *ShoppingCartAPIController) CreateShoppingCart(w http.ResponseWriter, r *http.Request) {
	shoppingCartParam := models.ShoppingCartCreate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&shoppingCartParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertShoppingCartCreateRequired(shoppingCartParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertShoppingCartCreateConstraints(shoppingCartParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateShoppingCart(r.Context(), shoppingCartParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteShoppingCart - Deletes a ShoppingCart
func (c *ShoppingCartAPIController) DeleteShoppingCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &models.RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.DeleteShoppingCart(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListShoppingCart - List or find ShoppingCart objects
func (c *ShoppingCartAPIController) ListShoppingCart(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	var fieldsParam string
	if query.Has("fields") {
		param := query.Get("fields")

		fieldsParam = param
	} else {
	}
	var offsetParam int32
	if query.Has("offset") {
		param, err := parseNumericParameter[int32](
			query.Get("offset"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &models.ParsingError{Param: "offset", Err: err}, nil)
			return
		}

		offsetParam = param
	} else {
	}
	var limitParam int32
	if query.Has("limit") {
		param, err := parseNumericParameter[int32](
			query.Get("limit"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &models.ParsingError{Param: "limit", Err: err}, nil)
			return
		}

		limitParam = param
	} else {
	}
	result, err := c.service.ListShoppingCart(r.Context(), fieldsParam, offsetParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// PatchShoppingCart - Updates partially a ShoppingCart
func (c *ShoppingCartAPIController) PatchShoppingCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &models.RequiredError{"id"}, nil)
		return
	}
	shoppingCartParam := models.ShoppingCartUpdate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&shoppingCartParam); err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertShoppingCartUpdateRequired(shoppingCartParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertShoppingCartUpdateConstraints(shoppingCartParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PatchShoppingCart(r.Context(), idParam, shoppingCartParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// RetrieveShoppingCart - Retrieves a ShoppingCart by ID
func (c *ShoppingCartAPIController) RetrieveShoppingCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &models.RequiredError{"id"}, nil)
		return
	}
	var fieldsParam string
	if query.Has("fields") {
		param := query.Get("fields")

		fieldsParam = param
	} else {
	}
	result, err := c.service.RetrieveShoppingCart(r.Context(), idParam, fieldsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
