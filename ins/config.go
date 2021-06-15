package ins

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

type MQTTConfigurations struct {
	Prefix              string
	Broker              string
	Cleansess           bool
	ClientId            string
	Qos                 int
	User                string
	Password            string
}

type ServiceConfigurations struct {
	EnableTls           bool
	TlsCert             string
	TlsKey              string
	Address             string
	Port                int
	Timeout             int64
}

type CenterGWConfigurations struct {
	Service             ServiceConfigurations
	UWBServer           ServiceConfigurations
	MQTT                MQTTConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int
}

type PublishConfigurations struct {
	Service             ServiceConfigurations
	Remote              ServiceConfigurations
	MQTT                MQTTConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int
}

type SubscribeConfigurations struct {
	Remote              ServiceConfigurations
	MQTT                MQTTConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int
}

type RelayConfigurations struct {
	Service             ServiceConfigurations
	Remote              ServiceConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int
}

type ServerConfigurations struct {
	Service             ServiceConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int
}

type ClientConfigurations struct {
	Remote              ServiceConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int
}

func (v MQTTConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("Broker: %s", v.Broker))
	strings = append(strings, fmt.Sprintf("Cleansess: %t", v.Cleansess))
	strings = append(strings, fmt.Sprintf("ClientId: %s", v.ClientId))
	strings = append(strings, fmt.Sprintf("User: %s", v.User))
	strings = append(strings, fmt.Sprintf("Password: %s", v.Password))

	return strings
}

func (v ServiceConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("EnableTls: %t", v.EnableTls))
	strings = append(strings, fmt.Sprintf("TlsCert: %s", v.TlsCert))
	strings = append(strings, fmt.Sprintf("TlsKey: %s", v.TlsKey))
	strings = append(strings, fmt.Sprintf("Address: %s", v.Address))
	strings = append(strings, fmt.Sprintf("Port: %d", v.Port))
	strings = append(strings, fmt.Sprintf("Timeout: %d", v.Timeout))

	return strings
}

func (v CenterGWConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("Service: %s", v.Service.ToString()))
	strings = append(strings, fmt.Sprintf("Remote: %s", v.UWBServer.ToString()))
	strings = append(strings, fmt.Sprintf("MQTT: %s", v.MQTT.ToString()))
	strings = append(strings, fmt.Sprintf("SourceId: %s", v.SourceId))
	strings = append(strings, fmt.Sprintf("EventLogUrl: %s", v.EventLogUrl))
	strings = append(strings, fmt.Sprintf("DiagInterval: %d", v.DiagInterval))

	return strings
}

func (v PublishConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("Service: %s", v.Service.ToString()))
	strings = append(strings, fmt.Sprintf("Remote: %s", v.Remote.ToString()))
	strings = append(strings, fmt.Sprintf("MQTT: %s", v.MQTT.ToString()))
	strings = append(strings, fmt.Sprintf("SourceId: %s", v.SourceId))
	strings = append(strings, fmt.Sprintf("EventLogUrl: %s", v.EventLogUrl))
	strings = append(strings, fmt.Sprintf("DiagInterval: %d", v.DiagInterval))

	return strings
}

func (v SubscribeConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("Remote: %s", v.Remote.ToString()))
	strings = append(strings, fmt.Sprintf("MQTT: %s", v.MQTT.ToString()))
	strings = append(strings, fmt.Sprintf("SourceId: %s", v.SourceId))
	strings = append(strings, fmt.Sprintf("EventLogUrl: %s", v.EventLogUrl))
	strings = append(strings, fmt.Sprintf("DiagInterval: %d", v.DiagInterval))

	return strings
}

func (v RelayConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("Service: %s", v.Service.ToString()))
	strings = append(strings, fmt.Sprintf("Remote: %s", v.Remote.ToString()))
	strings = append(strings, fmt.Sprintf("SourceId: %s", v.SourceId))
	strings = append(strings, fmt.Sprintf("EventLogUrl: %s", v.EventLogUrl))
	strings = append(strings, fmt.Sprintf("DiagInterval: %d", v.DiagInterval))

	return strings
}

func (v ServerConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("Service: %s", v.Service.ToString()))
	strings = append(strings, fmt.Sprintf("SourceId: %s", v.SourceId))
	strings = append(strings, fmt.Sprintf("EventLogUrl: %s", v.EventLogUrl))
	strings = append(strings, fmt.Sprintf("DiagInterval: %d", v.DiagInterval))

	return strings
}

func (v ClientConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("Remote: %s", v.Remote.ToString()))
	strings = append(strings, fmt.Sprintf("SourceId: %s", v.SourceId))
	strings = append(strings, fmt.Sprintf("EventLogUrl: %s", v.EventLogUrl))
	strings = append(strings, fmt.Sprintf("DiagInterval: %d", v.DiagInterval))

	return strings
}

/**
 * PublishConfigurations JSON 형태로 출력한다.
 */
func (v *PublishConfigurations) PrintConfigurations() {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))
}

/**
 * PublishConfigurations JSON 형태로 출력한다.
 */
func (v *CenterGWConfigurations) PrintConfigurations() {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))
}

/**
 * RelayConfigurations JSON 형태로 출력한다.
 */
func (v *RelayConfigurations) PrintConfigurations() {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))
}

/**
 * SubscribeConfigurations을 JSON 형태로 출력한다.
 */
func (v *SubscribeConfigurations) PrintConfigurations() {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))
}

/**
 * ServerConfigurations을 JSON 형태로 출력한다.
 */
func (v *ServerConfigurations) PrintConfigurations() {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))
}

/**
 * ClientConfigurations을 JSON 형태로 출력한다.
 */
func (v *ClientConfigurations) PrintConfigurations() {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))
}

/**
 * 지정된 이름의 인자가 전달되었는지 확인한다.
 */
func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
