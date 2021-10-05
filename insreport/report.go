package insreport

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"github.com/industry-netsecurity-solution/ins-security-channel/request"
	"net/http"
	"net/url"
)

type SecurityEventLog struct {
	EventType     string `json:"eventType"`
	EventSource   string `json:"eventSource"`
	GatewayType   string `json:"gatewayType"`
	GatewayId     string `json:"gatewayId"`
	EventTime     string `json:"eventTime"`
	Message       string `json:"message"`
	Content       string `json:"content"`
}

func (v *SecurityEventLog) EncodeLog() ([]byte, error) {
	data, _e := json.Marshal(v)
	if _e != nil {
		return nil, _e
	}

	return data, nil
}

func ReportSecurityLog(reportUrl string, evtType, eventSource, evtGwType, sourceId, message, content string) error {
	evt := SecurityEventLog {
		EventType: evtType,
		EventSource: eventSource,
		GatewayType: evtGwType,
		GatewayId: sourceId,
		EventTime: ins.TimeYYmmddHHMMSS(nil),
		Message: message,
		Content: content,
	}

	var u *url.URL = nil
	var data []byte = nil
	var err  error

	if data, err = evt.EncodeLog(); err != nil {
		return err
	}

	if len(reportUrl) == 0 {
		return nil
	}

	if u, err = url.Parse(reportUrl); err != nil {
		return err
	}

	// 기존의 Config file로 firmware 업데이트가 있는지 검사한다.
	reportParam := request.NewRequestParam()
	reportParam.SetHeader("Content-Type", "application/json")
	reportParam.SetHeader("Accept", "application/json")
	reportParam.SetData(data)


	r := &request.RequestURL{*u}

	if _, e := r.DoRequest(reportParam, func(resp *resty.Response) error {
		status := resp.StatusCode()
		if status != 200 && status != 201 && status != 202 && status != 205 {
			return errors.New(fmt.Sprintf("%d %s - %s", status, http.StatusText(status), resp.Request.URL))
		}
		return nil
	}); e != nil {
		return e
	}

	return nil
}
