package isc

import (
	"encoding/json"
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
