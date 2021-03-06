// Code generated by go-swagger; DO NOT EDIT.

package feature_flags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPrivateDomainCreationFeatureFlagHandlerFunc turns a function with the right signature into a get private domain creation feature flag handler
type GetPrivateDomainCreationFeatureFlagHandlerFunc func(GetPrivateDomainCreationFeatureFlagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPrivateDomainCreationFeatureFlagHandlerFunc) Handle(params GetPrivateDomainCreationFeatureFlagParams) middleware.Responder {
	return fn(params)
}

// GetPrivateDomainCreationFeatureFlagHandler interface for that can handle valid get private domain creation feature flag params
type GetPrivateDomainCreationFeatureFlagHandler interface {
	Handle(GetPrivateDomainCreationFeatureFlagParams) middleware.Responder
}

// NewGetPrivateDomainCreationFeatureFlag creates a new http.Handler for the get private domain creation feature flag operation
func NewGetPrivateDomainCreationFeatureFlag(ctx *middleware.Context, handler GetPrivateDomainCreationFeatureFlagHandler) *GetPrivateDomainCreationFeatureFlag {
	return &GetPrivateDomainCreationFeatureFlag{Context: ctx, Handler: handler}
}

/*GetPrivateDomainCreationFeatureFlag swagger:route GET /config/feature_flags/private_domain_creation featureFlags getPrivateDomainCreationFeatureFlag

Get the Private Domain Creation feature flag

curl --insecure -i %s/v2/config/feature_flags/private_domain_creation -X GET -H 'Authorization: %s'

*/
type GetPrivateDomainCreationFeatureFlag struct {
	Context *middleware.Context
	Handler GetPrivateDomainCreationFeatureFlagHandler
}

func (o *GetPrivateDomainCreationFeatureFlag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPrivateDomainCreationFeatureFlagParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
