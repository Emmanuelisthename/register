package handlers

import (
	"context"
	"dev/register.git/foundation/web"
	"net/http"
)

// create ...
func (r Register) create(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	// Get request data
	var request = struct {
		Value string `validate:"required"`
	}{}

	// Decode, sanitize & validate request
	err := web.Decode(req, &request)
	if err != nil {
		return err
	}

	// Log
	r.Service.Log.Printf("Creating...")

	// Create
	err = r.Service.Create(ctx)
	if err != nil {
		return err
	}

	// Send response data
	response := struct {
		Status string `json:"Status"`
	}{
		Status: "Success",
	}
	return web.Respond(ctx, w, response, http.StatusOK)

}
