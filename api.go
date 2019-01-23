package movizor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// API - это клиент к API Мовизора. Сервиса определения гео-координат по GSM.
type API struct {
	Endpoint string
	Project  string
	Token    string
	Client   *http.Client
	IsDebug  bool

	//Buffer          int
	//shutdownChannel chan interface{}
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
		IsDebug:  false,
		//Buffer:          100,
		//shutdownChannel: make(chan interface{}),
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
	bytes, err := api.decodeAPIResponse(resp.Body, &apiResp)
	if err != nil {
		return APIResponse{}, err
	}

	if apiResp.Result == "success" {
		if api.IsDebug {
			log.Printf("INFO: request: %s\nresponse: %s", req.URL, bytes)
		}
		return apiResp, nil
	}

	err = errors.New(fmt.Sprintf("movizor API returns error on request: %s - %s",
		apiResp.ErrorCode, apiResp.ErrorText))
	if api.IsDebug {
		log.Printf("ERROR: request: %s\nresponse: %s", req.URL, bytes)
	}

	return apiResp, err
}

// decodeAPIResponse checks if response
func (api *API) decodeAPIResponse(responseBody io.Reader, resp *APIResponse) (_ []byte, err error) {
	if !api.IsDebug {
		dec := json.NewDecoder(responseBody)
		err = dec.Decode(resp)
		return
	}

	// if logging is on, read response body
	data, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}

	return data, nil
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

func (api *API) AddObject(o Object, oo *ObjectOptions) (APIResponse, error) {
	v := o.values()
	if oo != nil {
		oo.addValuesTo(&v)
	}

	resp, err := api.MakeRequest("object_add", v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (api *API) GetObjectInfo(o Object) (ObjectInfo, error) {
	resp, err := api.MakeRequest("object_get", o.values())
	if err != nil {
		return ObjectInfo{}, err
	}

	var oi ObjectInfo
	err = json.Unmarshal(resp.Data, &oi)
	if err != nil {
		return ObjectInfo{}, err
	}

	return oi, nil
}

func (api *API) EditObject(o Object, oo *ObjectOptions) (APIResponse, error) {
	v := o.values()
	if oo != nil {
		oo.addValuesTo(&v)
	}

	resp, err := api.MakeRequest("object_edit", v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (api *API) DeleteObject(o Object) (APIResponse, error) {
	resp, err := api.MakeRequest("object_delete", o.values())
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (api *API) ReactivateObject(o Object) (APIResponse, error) {
	resp, err := api.MakeRequest("object_reactivate", o.values())
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (api *API) CancelTariffChangeObject(o Object) (APIResponse, error) {
	resp, err := api.MakeRequest("object_cancel_tariff", o.values())
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetObjectPositions returns slice of objects with its positions and ETA
func (api *API) GetObjects() (ObjectsWithStatus, error) {
	resp, err := api.MakeRequest("object_list", nil)
	if err != nil {
		return ObjectsWithStatus{}, err
	}

	var o ObjectsWithStatus
	err = json.Unmarshal(resp.Data, &o)
	if err != nil {
		return ObjectsWithStatus{}, err
	}

	return o, nil
}

// GetObjectPositions returns slice of objects with its positions and ETA
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

// GetOperatorInfo возвращает информацию по оператору объекта трекинга (номеру телефона)
func (api *API) GetOperatorInfo(o Object) (OperatorInfo, error) {
	resp, err := api.MakeRequest("get_operator", o.values())
	if err != nil {
		return OperatorInfo{}, err
	}

	var oi OperatorInfo
	err = json.Unmarshal(resp.Data, &oi)
	if err != nil {
		return OperatorInfo{}, err
	}

	return oi, nil
}

// GetEvents получает список событий, с возможностью определить с какого id события выводить данные.
func (api *API) GetEvents(o ObjectEventsOptions) (ObjectEvents, error) {
	resp, err := api.MakeRequest("events", o.values())
	if err != nil {
		return ObjectEvents{}, err
	}

	var oe ObjectEvents
	err = json.Unmarshal(resp.Data, &oe)
	if err != nil {
		return ObjectEvents{}, err
	}

	return oe, nil
}

// DeleteEventsSubscription удаляет подписку по ее id. Для получения id используйте GetEventSubscriptions.
func (api *API) DeleteEventsSubscription(id uint64) (APIResponse, error) {
	v := url.Values{}
	v.Add("id", strconv.FormatUint(id, 10))
	resp, err := api.MakeRequest("events_subscribe_delete", v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetEventSubscriptions получает список подписок активных на текущий момент.
func (api *API) GetEventSubscriptions() (SubscribedEvents, error) {
	resp, err := api.MakeRequest("events_subscribe_list", nil)
	if err != nil {
		return SubscribedEvents{}, err
	}

	var se SubscribedEvents
	err = json.Unmarshal(resp.Data, &se)
	if err != nil {
		return SubscribedEvents{}, err
	}

	return se, nil
}

// SubscribeEvent выполняет подписку на указанное тип события для всех объектов (телефонов) или по списку.
func (api *API) SubscribeEvent(o SubscribeEventOptions) (APIResponse, error) {
	v, err := o.values()
	if err != nil {
		return APIResponse{}, err
	}

	resp, err := api.MakeRequest("events_subscribe_add", v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
