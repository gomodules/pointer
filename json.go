package aws

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSONClient is the underlying client for json APIs.
type JSONClient struct {
	Context      Context
	Client       *http.Client
	Endpoint     string
	TargetPrefix string
	JSONVersion  string
}

// Do sends an HTTP request and returns an HTTP response, following policy
// (e.g. redirects, cookies, auth) as configured on the client.
func (c *JSONClient) Do(op, method, uri string, req, resp interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(method, c.Endpoint+uri, bytes.NewReader(b))
	if err != nil {
		return err
	}
	httpReq.Header.Set("User-Agent", "aws-go")
	httpReq.Header.Set("X-Amz-Target", c.TargetPrefix+"."+op)
	httpReq.Header.Set("Content-Type", "application/x-amz-json-"+c.JSONVersion)
	if err := c.Context.sign(httpReq); err != nil {
		return err
	}

	httpResp, err := c.Client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		var err jsonErrorResponse
		if err := json.NewDecoder(httpResp.Body).Decode(&err); err != nil {
			return err
		}
		return err.Err()
	}

	return json.NewDecoder(httpResp.Body).Decode(resp)
}

type jsonErrorResponse struct {
	Type    string `json:"__type"`
	Message string `json:"message"`
}

func (e jsonErrorResponse) Err() error {
	return APIError{
		Type:    e.Type,
		Message: e.Message,
	}
}
