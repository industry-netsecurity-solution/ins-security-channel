package ins

import (
	"encoding/json"
	"errors"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"net/http"
)

type EventLog struct {
	GatewayType string `json:"gatewayType"`
	EventType   string `json:"eventType"`
	GatewayId   string `json:"gatewayId"`
	EventTime   string `json:"eventTime"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	Content     string `json:"content"`
}

func (v *EventLog) EncodeLog() ([]byte, error) {
	data, _e := json.Marshal(v)
	if _e != nil {
		return nil, _e
	}

	return data, nil
}

func ReportEvent(reportUrl *HttpConfigurations, log EventLog) error {

	//data, _e := json.Marshal(log)
	data, _e := log.EncodeLog()
	if _e != nil {
		return _e
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	u, err := reportUrl.Url()
	if err != nil {
		return err
	}

	url := u.String()

	querypath := []string{}

	statusCode, err := HttpPost(reportUrl, querypath, headers, data, func(resp *resty.Response) {
	})
	if err != nil {
		return err
	}

	if statusCode == 200 || statusCode == 201 || statusCode == 202 || statusCode == 205 {
		return nil
	}

	return errors.New(fmt.Sprintf("%d %s - %s", statusCode, http.StatusText(statusCode), url))
}

func ReportLog(reportUrl *HttpConfigurations, sourceId string, evtGwType string, evtType string, status string, message string, content string) error {
	evt := EventLog{
		GatewayType: evtGwType,
		EventType:   evtType,
		GatewayId:   sourceId,
		EventTime:   TimeYYmmddHHMMSS(nil),
		Status:      status,
		Message:     message,
		Content:     content,
	}

	if reportUrl == nil {
		return nil
	}

	return ReportEvent(reportUrl, evt)
}
