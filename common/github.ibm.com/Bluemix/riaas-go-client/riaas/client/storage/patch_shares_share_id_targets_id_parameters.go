// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchSharesShareIDTargetsIDParams creates a new PatchSharesShareIDTargetsIDParams object
// with the default values initialized.
func NewPatchSharesShareIDTargetsIDParams() *PatchSharesShareIDTargetsIDParams {
	var ()
	return &PatchSharesShareIDTargetsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSharesShareIDTargetsIDParamsWithTimeout creates a new PatchSharesShareIDTargetsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSharesShareIDTargetsIDParamsWithTimeout(timeout time.Duration) *PatchSharesShareIDTargetsIDParams {
	var ()
	return &PatchSharesShareIDTargetsIDParams{

		timeout: timeout,
	}
}

// NewPatchSharesShareIDTargetsIDParamsWithContext creates a new PatchSharesShareIDTargetsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSharesShareIDTargetsIDParamsWithContext(ctx context.Context) *PatchSharesShareIDTargetsIDParams {
	var ()
	return &PatchSharesShareIDTargetsIDParams{

		Context: ctx,
	}
}

// NewPatchSharesShareIDTargetsIDParamsWithHTTPClient creates a new PatchSharesShareIDTargetsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSharesShareIDTargetsIDParamsWithHTTPClient(client *http.Client) *PatchSharesShareIDTargetsIDParams {
	var ()
	return &PatchSharesShareIDTargetsIDParams{
		HTTPClient: client,
	}
}

/*PatchSharesShareIDTargetsIDParams contains all the parameters to send to the API endpoint
for the patch shares share ID targets ID operation typically these are written to a http.Request
*/
type PatchSharesShareIDTargetsIDParams struct {

	/*Body*/
	Body PatchSharesShareIDTargetsIDBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The mount target identifier

	*/
	ID string
	/*ShareID
	  The file share identifier

	*/
	ShareID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) WithTimeout(timeout time.Duration) *PatchSharesShareIDTargetsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) WithContext(ctx context.Context) *PatchSharesShareIDTargetsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) WithHTTPClient(client *http.Client) *PatchSharesShareIDTargetsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) WithBody(body PatchSharesShareIDTargetsIDBody) *PatchSharesShareIDTargetsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) SetBody(body PatchSharesShareIDTargetsIDBody) {
	o.Body = body
}

// WithGeneration adds the generation to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) WithGeneration(generation int64) *PatchSharesShareIDTargetsIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) WithID(id string) *PatchSharesShareIDTargetsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) SetID(id string) {
	o.ID = id
}

// WithShareID adds the shareID to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) WithShareID(shareID string) *PatchSharesShareIDTargetsIDParams {
	o.SetShareID(shareID)
	return o
}

// SetShareID adds the shareId to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) SetShareID(shareID string) {
	o.ShareID = shareID
}

// WithVersion adds the version to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) WithVersion(version string) *PatchSharesShareIDTargetsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch shares share ID targets ID params
func (o *PatchSharesShareIDTargetsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSharesShareIDTargetsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param share_id
	if err := r.SetPathParam("share_id", o.ShareID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
