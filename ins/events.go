package ins

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (v *EventLog) SetEventGatewayType(gatewayType string) {
	v.GatewayType = gatewayType
}

func (v *EventLog) SetEventType(eventType string) {
	v.EventType = eventType
}

func (v *EventLog) SetEventGatewayId(gatewayId string) {
	v.GatewayId = gatewayId
}

func (v *EventLog) SetEventTime(eventTime string) {
	v.EventTime = eventTime
}

func (v *EventLog) SetEventStatus(status string) {
	v.Status = status
}

func (v *EventLog) SetEventMessage(message string) {
	v.Message = message
}

func (v *EventLog) SetEventContent(content string) {
	v.Content = content
}

func (v *EventLog) EventLog() ([]byte, error) {
	data, _e := json.Marshal(v)
	if _e != nil {
		return nil, _e
	}

	return data, nil
}

func ReportEvent(url string, log EventLog) error {

	//data, _e := json.Marshal(log)
	data, _e := log.EventLog()
	if _e != nil {
		return _e
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
/*
	if 0 < len(Config.Report.Authorization) {
		encoded := base64.StdEncoding.EncodeToString([]byte(Config.Report.Authorization))
		headers["Authorization"] = fmt.Sprintf("Bearer %s", encoded)
	}
*/

	statusCode, err := HttpPost(url, headers, data, func(resp *http.Response) {
	})
	if err != nil {
		return err
	}

	if statusCode == 200 || statusCode == 201 || statusCode == 202 || statusCode == 205 {
		return nil
	}

	return errors.New(fmt.Sprintf("%d %s - %s", statusCode, http.StatusText(statusCode), url))


/*
	//data, _e := json.Marshal(log)
	data, _e := log.EventLog()
	if _e != nil {
		return _e
	}


	fmt.Println(string(data))

	// HTTP를 통한 데이터 전송
	client := resty.New()
	client.SetCloseConnection(true)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(resp)

	return nil
*/
}

func ReportLog(url string, sourceId string, evtGwType string, evtType string, status string, message string, content string) error {
	evt := EventLog {
		GatewayType: evtGwType,
		EventType: evtType,
		GatewayId: sourceId,
		EventTime: TimeYYmmddHHMMSS(nil),
		Status: status,
		Message: message,
		Content: content,
	}

	if len(url) == 0 {
		return nil
	}

	return ReportEvent(url, evt)
}
