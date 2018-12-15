package movizor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// API gives a set of interfaces to Movizor service API
type API struct {
	Endpoint string
	Project  string
	Token    string
	Client   *http.Client
}

// NewMovizorAPIWithEndpoint creates new instance of Movizor API
// It can get non standard Movizor endpoint
// in case it will be moved to another address.
func NewMovizorAPIWithEndpoint(endp string, prj string, token string) (*API, error) {
	api := &API{
		Endpoint: endp,
		Project:  prj,
		Token:    token,
		Client:   &http.Client{},
	}
	return api, nil
}

// NewMovizorAPI creates new instance of Movizor API with default url of endpoint
func NewMovizorAPI(prj string, token string) (*API, error) {
	return NewMovizorAPIWithEndpoint(DefaultAPIMovizorEndpoint, prj, token)
}

// MakeRequest makes request to Movizor API with specific action.
// It gives one point of requests
func (api *API) MakeRequest(action string, params url.Values) (APIResponse, error) {
	// MakeRequest itself
	endpAction := fmt.Sprintf(fmt.Sprint(api.Endpoint, APIMovizorEndpointSuffix), api.Project, action)
	if params == nil {
		params = url.Values{}
	}
	params.Add("key", api.Token)
	uri, _ := url.Parse(endpAction)

	uri.RawQuery = params.Encode()
	endpAction = uri.String()

	req, err := http.NewRequest("GET", endpAction, nil)
	if err != nil {
		return APIResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := api.Client.Do(req)
	if err != nil {
		return APIResponse{}, err
	}
	defer resp.Body.Close()

	// Response handling
	var apiResp APIResponse
	err = api.decodeAPIResponse(resp.Body, &apiResp)
	if err != nil {
		return apiResp, err
	}

	return apiResp, nil
}

// decodeAPIResponse checks if response
func (api *API) decodeAPIResponse(responseBody io.Reader, resp *APIResponse) (err error) {
	dec := json.NewDecoder(responseBody)
	err = dec.Decode(resp)
	if err != nil {
		return
	}

	if resp.Result == "success" {
		return
	}

	err = errors.New(fmt.Sprintf("movizor API returns error on request: %s - %s",
		resp.ErrorCode, resp.ErrorText))
	return
}

// GetBalance returns current remain of money and collected credit
// with tariffs that set for the Project
func (api *API) GetBalance() (Balance, error) {
	resp, err := api.MakeRequest("balance", nil)
	if err != nil {
		return Balance{}, err
	}

	var b Balance
	err = json.Unmarshal(resp.Data, &b)
	if err != nil {
		return Balance{}, err
	}

	return b, nil
}

// GetObjectPositions returns array of objects with its positions and ETA
func (api *API) GetObjectPositions() (ObjectPositions, error) {
	resp, err := api.MakeRequest("pos_objects", nil)
	if err != nil {
		return ObjectPositions{}, err
	}

	var op ObjectPositions
	err = json.Unmarshal(resp.Data, &op)
	if err != nil {
		return ObjectPositions{}, err
	}

	return op, nil
}
