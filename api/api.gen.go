// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Track defines model for track.
type Track struct {
	Album       string  `json:"album"`
	Artist      string  `json:"artist"`
	Duration    *int    `json:"duration,omitempty"`
	Genre       *string `json:"genre,omitempty"`
	Id          *string `json:"id,omitempty"`
	ReleaseYear *int    `json:"release_year,omitempty"`
	Title       string  `json:"title"`
}

// TrackInput defines model for track_input.
type TrackInput struct {
	Album       *string `json:"album,omitempty"`
	Artist      *string `json:"artist,omitempty"`
	Duration    *int    `json:"duration,omitempty"`
	Genre       *string `json:"genre,omitempty"`
	ReleaseYear *int    `json:"release_year,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// PostTracksJSONRequestBody defines body for PostTracks for application/json ContentType.
type PostTracksJSONRequestBody = TrackInput

// PutTracksIdJSONRequestBody defines body for PutTracksId for application/json ContentType.
type PutTracksIdJSONRequestBody = TrackInput

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all music tracks
	// (GET /tracks)
	GetTracks(ctx echo.Context) error
	// Create a new music track
	// (POST /tracks)
	PostTracks(ctx echo.Context) error
	// Delete a music track
	// (DELETE /tracks/{id})
	DeleteTracksId(ctx echo.Context, id string) error
	// Get a music track by ID
	// (GET /tracks/{id})
	GetTracksId(ctx echo.Context, id string) error
	// Update a music track
	// (PUT /tracks/{id})
	PutTracksId(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTracks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTracks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTracks(ctx)
	return err
}

// PostTracks converts echo context to params.
func (w *ServerInterfaceWrapper) PostTracks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTracks(ctx)
	return err
}

// DeleteTracksId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTracksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTracksId(ctx, id)
	return err
}

// GetTracksId converts echo context to params.
func (w *ServerInterfaceWrapper) GetTracksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTracksId(ctx, id)
	return err
}

// PutTracksId converts echo context to params.
func (w *ServerInterfaceWrapper) PutTracksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutTracksId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/tracks", wrapper.GetTracks)
	router.POST(baseURL+"/tracks", wrapper.PostTracks)
	router.DELETE(baseURL+"/tracks/:id", wrapper.DeleteTracksId)
	router.GET(baseURL+"/tracks/:id", wrapper.GetTracksId)
	router.PUT(baseURL+"/tracks/:id", wrapper.PutTracksId)

}