package manage

import (
	"encoding/json"
	"fmt"
)

type Interface interface {
	SwitchLogging(name string, instructment uint8) error
	GetConnections(mode uint8) (uint64, error)
}

type NoopInterface struct {}

func (*NoopInterface) SwitchLogging(name string, instructment uint8) error {
	return nil
}

func (*NoopInterface) GetConnections(mode uint8) (uint64, error) {
	return 0, nil
}

var validate map[string]struct{} = map[string]struct{}{
	"heartbeat": {},
}

type SetLoggingSwitchRequest struct {
	Feature string `json:"feature"`
	Switch  uint8  `json:"switch"`
}

func (f *SetLoggingSwitchRequest) UnmarshalJSON(data []byte) error {
	type plain SetLoggingSwitchRequest
	request := &plain{}
	if err := json.Unmarshal(data, request); err != nil {
		return err
	}
	if _, ok := validate[request.Feature]; !ok {
		return fmt.Errorf("invalid feature")
	}
	if request.Switch > 1 {
		return fmt.Errorf("invalid status of switch")
	}
	f = (*SetLoggingSwitchRequest)(request)
	return nil
}

type GetConnectionsRequest struct {
	Type uint8 `json:"type"`
}

type GetConnectionsResponse struct {
	Response
	Data *GetConnectionsResponseData `json:"data"`
}

type GetConnectionsResponseData struct {
	Count  uint64   `json:"count"`
	Handle []string `json:"handle"`
}

func NewGetConnectionsResponse(response *Response, data *GetConnectionsResponseData) *GetConnectionsResponse {
	return &GetConnectionsResponse{
		Response: Response{
			Status:  Succeeded,
			Message: "success",
		},
		Data: data,
	}
}

type responseStatus int

const (
	Succeeded responseStatus = 0
	Failed    responseStatus = 1
)

type Response struct {
	Status  responseStatus `json:"status"`
	Message string         `json:"msg"`
}

func NewResponse(status responseStatus, msg string) *Response {
	return &Response{
		Status:  status,
		Message: msg,
	}
}
