//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ObjectsClassDeleteHandlerFunc turns a function with the right signature into a objects class delete handler
type ObjectsClassDeleteHandlerFunc func(ObjectsClassDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ObjectsClassDeleteHandlerFunc) Handle(params ObjectsClassDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ObjectsClassDeleteHandler interface for that can handle valid objects class delete params
type ObjectsClassDeleteHandler interface {
	Handle(ObjectsClassDeleteParams, *models.Principal) middleware.Responder
}

// NewObjectsClassDelete creates a new http.Handler for the objects class delete operation
func NewObjectsClassDelete(ctx *middleware.Context, handler ObjectsClassDeleteHandler) *ObjectsClassDelete {
	return &ObjectsClassDelete{Context: ctx, Handler: handler}
}

/*ObjectsClassDelete swagger:route DELETE /objects/{className}/{id} objects objectsClassDelete

Delete object based on its class and UUID.

Delete a single data object.

*/
type ObjectsClassDelete struct {
	Context *middleware.Context
	Handler ObjectsClassDeleteHandler
}

func (o *ObjectsClassDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewObjectsClassDeleteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}