// Package imagebuilder provides access to Red Hat's Image Builder API.
//
// General information about how to access Red Hat APIs can be found here:
//
//   https://access.redhat.com/articles/3626371

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=imagebuilder --generate types -o imagebuilderapi.gen.go openapi.json

package imagebuilder

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

type Client struct {
	client *http.Client
}

func NewClient(ctx context.Context, refreshToken string) *Client {
	conf := &oauth2.Config{
		ClientID: "rhsm-api",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token",
		},
	}

	return &Client{
		client: conf.Client(ctx, &oauth2.Token{
			RefreshToken: refreshToken,
		}),
	}
}

func (c *Client) Compose(cr *ComposeRequest) (string, error) {
	var resp ComposeResponse
	err := c.postJSON("https://cloud.redhat.com/api/image-builder/v1.0/compose", cr, &resp)
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

func (c *Client) GetComposeStatus(id string) (*ImageStatus, error) {
	var resp ComposeStatus
	err := c.getJSON(fmt.Sprintf("https://cloud.redhat.com/api/image-builder/v1.0/composes/%s", id), &resp)
	if err != nil {
		return nil, err
	}
	return &resp.ImageStatus, nil
}

func responseError(r *http.Response) error {
	var errors HTTPErrorList
	err := json.NewDecoder(r.Body).Decode(&errors)
	if err != nil {
		return fmt.Errorf("server responded with code %v and unexpected or empty JSON: %v", r.StatusCode, err)
	}

	var msg strings.Builder
	fmt.Fprintf(&msg, "%v errors:\n", len(errors.Errors))
	for _, e := range errors.Errors {
		fmt.Fprintf(&msg, "  %s: %s\n", e.Title, e.Detail)
	}

	return fmt.Errorf("server responded with code %v:\n%v", r.StatusCode, msg.String())
}

func (c *Client) getJSON(url string, reply interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return responseError(resp)
	}

	if reply != nil {
		err = json.NewDecoder(resp.Body).Decode(reply)
		if err != nil {
			return fmt.Errorf("received invalid JSON: %v", err)
		}
	}

	return nil
}

func (c *Client) postJSON(url string, request, reply interface{}) error {
	var body bytes.Buffer
	if request != nil {
		err := json.NewEncoder(&body).Encode(request)
		if err != nil {
			return fmt.Errorf("error JSON-encoding %v: %v", request, err)
		}
	}

	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return responseError(resp)
	}

	if reply != nil {
		err = json.NewDecoder(resp.Body).Decode(reply)
		if err != nil {
			return fmt.Errorf("received invalid JSON: %v", err)
		}
	}

	return nil
}
