// Code generated by sysl DO NOT EDIT.
package petstore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Service interface for Petstore
type Service interface {
	GetPetsList(ctx context.Context, req *GetPetsListRequest) (*Pet, error)
}

// Client for Petstore API
type Client struct {
	client *http.Client
	url    string
}

// NewClient for Petstore
func NewClient(client *http.Client, serviceURL string) *Client {
	return &Client{client, serviceURL}
}

// GetPetsList ...
func (s *Client) GetPetsList(ctx context.Context, req *GetPetsListRequest) (*Pet, error) {
	required := []string{}
	var okResponse Pet
	var errorResponse Error
	u, err := url.Parse(fmt.Sprintf("%s/pets", s.url))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	q := u.Query()
	if req.Limit != nil {
		q.Add("limit", fmt.Sprintf("%v", *req.Limit))
	}

	u.RawQuery = q.Encode()
	result, err := restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, required, &okResponse, &errorResponse)
	if err != nil {
		response, ok := err.(*restlib.HTTPResult)
		if !ok {
			return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Petstore <- GET "+u.String(), err)
		}
		return nil, common.CreateDownstreamError(ctx, common.DownstreamResponseError, response.HTTPResponse, response.Body, &errorResponse)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkPetResponse, ok := result.Response.(*Pet)
	if ok {
		valErr := validator.Validate(OkPetResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkPetResponse, nil
	}

	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}
