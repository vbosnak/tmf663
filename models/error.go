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
	"encoding/json"
	"io"
	"os"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrTypeAssertionError is thrown when type an interface does not match the asserted type
	ErrTypeAssertionError = errors.New("unable to assert type")
)

// ParsingError indicates that an error has occurred when parsing request parameters
type ParsingError struct {
	Param string
	Err   error
}

func (e *ParsingError) Unwrap() error {
	return e.Err
}

func (e *ParsingError) Error() string {
	if e.Param == "" {
		return e.Err.Error()
	}

	return e.Param + ": " + e.Err.Error()
}

// RequiredError indicates that an error has occurred when parsing request parameters
type RequiredError struct {
	Field string
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("required field '%s' is zero value.", e.Field)
}

// ErrorHandler defines the required method for handling error. You may implement it and inject this into a controller if
// you would like errors to be handled differently from the DefaultErrorHandler
type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error, result *ImplResponse)

// DefaultErrorHandler defines the default logic on how to handle errors from the controller. Any errors from parsing
// request params will return a StatusBadRequest. Otherwise, the error code originating from the servicer will be used.
func DefaultErrorHandler(w http.ResponseWriter, _ *http.Request, err error, result *ImplResponse) {
	var parsingErr *ParsingError
	if ok := errors.As(err, &parsingErr); ok {
		// Handle parsing errors
		_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusBadRequest), w)
		return
	}

	var requiredErr *RequiredError
	if ok := errors.As(err, &requiredErr); ok {
		// Handle missing required errors
		_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusUnprocessableEntity), w)
		return
	} 

	// Handle all other errors
	_ = EncodeJSONResponse(err.Error(), &result.Code, w)
}

// EncodeJSONResponse uses the json encoder to write an interface to the http response with an optional status code
func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	wHeader := w.Header()

	f, ok := i.(*os.File)
	if ok {
		data, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		wHeader.Set("Content-Type", http.DetectContentType(data))
		wHeader.Set("Content-Disposition", "attachment; filename="+f.Name())
		if status != nil {
			w.WriteHeader(*status)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		_, err = w.Write(data)
		return err
	}
	wHeader.Set("Content-Type", "application/json; charset=UTF-8")

	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if i != nil {
		return json.NewEncoder(w).Encode(i)
	}

	return nil
}
