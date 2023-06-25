// Package classification User Service API.
//
//	    Schemes: http
//	    Version: 1.0.0
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Security:
//		- basic
//
//
//		SecurityDefinitions:
//		  Bearer:
//		    type: apiKey
//		    name: Authorization
//		    in: header
//
// swagger:meta
package public

//go:generate swagger generate spec -o ./static/swagger.json --scan-models
