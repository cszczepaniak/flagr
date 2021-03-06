// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewDeleteTagParams creates a new DeleteTagParams object
// no default values defined in spec.
func NewDeleteTagParams() DeleteTagParams {

	return DeleteTagParams{}
}

// DeleteTagParams contains all the bound params for the delete tag operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteTag
type DeleteTagParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*numeric ID of the flag
	  Required: true
	  Minimum: 1
	  In: path
	*/
	FlagID int64
	/*numeric ID of the tag
	  Required: true
	  Minimum: 1
	  In: path
	*/
	TagID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteTagParams() beforehand.
func (o *DeleteTagParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rFlagID, rhkFlagID, _ := route.Params.GetOK("flagID")
	if err := o.bindFlagID(rFlagID, rhkFlagID, route.Formats); err != nil {
		res = append(res, err)
	}

	rTagID, rhkTagID, _ := route.Params.GetOK("tagID")
	if err := o.bindTagID(rTagID, rhkTagID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFlagID binds and validates parameter FlagID from path.
func (o *DeleteTagParams) bindFlagID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("flagID", "path", "int64", raw)
	}
	o.FlagID = value

	if err := o.validateFlagID(formats); err != nil {
		return err
	}

	return nil
}

// validateFlagID carries on validations for parameter FlagID
func (o *DeleteTagParams) validateFlagID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("flagID", "path", int64(o.FlagID), 1, false); err != nil {
		return err
	}

	return nil
}

// bindTagID binds and validates parameter TagID from path.
func (o *DeleteTagParams) bindTagID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("tagID", "path", "int64", raw)
	}
	o.TagID = value

	if err := o.validateTagID(formats); err != nil {
		return err
	}

	return nil
}

// validateTagID carries on validations for parameter TagID
func (o *DeleteTagParams) validateTagID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("tagID", "path", int64(o.TagID), 1, false); err != nil {
		return err
	}

	return nil
}
