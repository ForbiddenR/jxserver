package manage

import (
	"encoding/json"
	"fmt"
	"strings"
)

var validate map[string]string = map[string]string{
	"login":                         "Login",
	"heartbeat":                     "Heartbeat",
	"notifyevent":                   "NotifyEvent",
	"notifyreport":                  "NotifyReport",
	"transactionevent":              "TransactionEvent",
	"metervalues":                   "MeterValues",
	"bmsinfo":                       "BMSInfo",
	"bmslimit":                      "BMSLimit",
	"statusnotification":            "StatusNotification",
	"logstatusnotification":         "LogStatusNotification",
	"firmwarestatusnotification":    "FirmwareStatusNotification",
	"reservationstatusnotification": "ReservationStatusNotification",
	"getbasereport":                 "GetBaseReport",
	"reset":                         "Reset",
	"getlog":                        "GetLog",
	"updatefirmware":                "UpdateFirmware",
	"requeststarttransaction":       "RequestStartTransaction",
	"requeststoptransaction":        "RequestStopTransaction",
	"reservenow":                    "ReserveNow",
	"cancelreservaion":              "CancelReservaion",
	"setchargingprofile":            "SetChargingProfile",
	"setpricescheme":                "SetPriceScheme",
	"setintellectcharging":          "SetIntellectCharging",
	"cancelintellectcharging":       "CancelIntellectCharging",
	"clearcache":                    "ClearCache",
	"setvariables":                  "SetVariables",
	"getvariables":                  "GetVariables",
	"getconnectorstatus":            "GetConnectorStatus",
	"authorize":                     "Authorize",
	"sendlocallist":                 "SendLocalList",
	"qrcode":                        "QRCode",
	"sendqrcode":                    "SendQRCode",
	"clearchargingprofile":          "ClearChargingProfile",
	"chargeencryinfonotification":   "ChargeEncryInfoNotification",
}

type SetLoggingSwitchRequest struct {
	Feature string `json:"feature"`
	Switch  uint8  `json:"switch"`
}

func (f *SetLoggingSwitchRequest) UnmarshalJSON(data []byte) error {
	var plain map[string]interface{}
	if err := json.Unmarshal(data, &plain); err != nil {
		return err
	}
	var feature string
	var ok bool
	var v interface{}
	if v, ok = plain["feature"]; !ok {
		return fmt.Errorf("feature is needed")
	} else if feature, ok = v.(string); !ok {
		return fmt.Errorf("invalid feature type")
	}
	request := &SetLoggingSwitchRequest{}
	if request.Feature, ok = validate[strings.ToLower(feature)]; !ok {
		return fmt.Errorf("invalid feature value")
	}

	var swh float64
	if v, ok = plain["switch"]; !ok {
		return fmt.Errorf("switch is needed")
	} else if swh, ok = v.(float64); !ok || int(swh) > 1 {
		return fmt.Errorf("invalid switch")
	}
	request.Switch = uint8(swh)
	*f = *request
	return nil
}

type GetConnectionsRequest struct {
	Type uint8 `json:"type"`
}

func (g *GetConnectionsRequest) UnmarshalJSON(data []byte) error {
	type plain GetConnectionsRequest
	request := &plain{}
	if err := json.Unmarshal(data, request); err != nil {
		return err
	}
	if request.Type != 1 {
		return fmt.Errorf("invalid value of type")
	}
	*g = (GetConnectionsRequest)(*request)
	return nil
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
			Status:  response.Status,
			Message: response.Message,
		},
		Data: data,
	}
}

type DisconnectConnectionRequest struct {
	Sn string `json:"sn"`
}

type GetConnectionStatusRequest struct {
	Sn string `json:"sn"`
}

type GetConnectionStatusResponse struct {
	Response
	Data *GetConnectionStatusResponseData `json:"data"`
}

type GetConnectionStatusResponseData struct {
	LocalAddress  string `json:"localAddress"`
	RemoteAddress string `json:"remoteAddress"`
}

type GetConnectionAlarmRulesResponse struct {
	Response
	Data *GetConnectionAlarmRulesResponseData `json:"data"`
}

type GetConnectionAlarmRulesResponseData struct {
	Rule  string `json:"rule"`
	Limit uint   `json:"limit"`
}

type SetConnectionAlarmRulesRequest struct {
	Rule  string `json:"rule"`
	Limit uint   `json:"limit"`
}

func (s *SetConnectionAlarmRulesRequest) UnmarshalJSON(data []byte) error {
	type plain SetConnectionAlarmRulesRequest
	request := &plain{}
	if err := json.Unmarshal(data, request); err != nil {
		return err
	}
	if request.Limit == 0 {
		return fmt.Errorf("invalid value of limit")
	}
	if request.Rule != "gte" && request.Rule != "gt" && request.Rule != "lte" && request.Rule != "e" && request.Rule != "lt" {
		return fmt.Errorf("invalid value of rule")
	}
	*s = (SetConnectionAlarmRulesRequest)(*request)
	return nil
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
