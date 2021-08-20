package ins

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"path"
)

type GatewayConfigurations struct {
	Date                string
	Model               string
	Manufacture         string
	Serial              string
}

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
	Port                int64
	Timeout             int64
}

type HttpConfigurations struct {
	EnableTls           bool
	Authorization       string
	TlsCert             string
	TlsKey              string
	Address             string
	Port                int64
	Path                string
	Timeout             int64
}

type PublishConfigurations struct {
	Service             ServiceConfigurations
	Remote              ServiceConfigurations
	MQTT                MQTTConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int64
}

type SubscribeConfigurations struct {
	Remote              ServiceConfigurations
	MQTT                MQTTConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int64
}

type RelayConfigurations struct {
	Service             ServiceConfigurations
	Remote              ServiceConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int64
}

type ServerConfigurations struct {
	Service             ServiceConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int64
}

type ClientConfigurations struct {
	Remote              ServiceConfigurations
	SourceId            string
	EventLogUrl         string
	DiagInterval        int64
}

type FirmwareConfigurations struct {
	Enable              bool
	ConfigFilepath      string
	DownlaodFilepath    string
	Http                HttpConfigurations
}

func (v GatewayConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("Date: %s", v.Date))
	strings = append(strings, fmt.Sprintf("Model: %t", v.Model))
	strings = append(strings, fmt.Sprintf("Manufacture: %s", v.Manufacture))
	strings = append(strings, fmt.Sprintf("Serial: %s", v.Serial))

	return strings
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

func (v HttpConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("EnableTls: %t", v.EnableTls))
	strings = append(strings, fmt.Sprintf("Authorization: %s", v.Authorization))
	strings = append(strings, fmt.Sprintf("TlsCert: %s", v.TlsCert))
	strings = append(strings, fmt.Sprintf("TlsKey: %s", v.TlsKey))
	strings = append(strings, fmt.Sprintf("Address: %s", v.Address))
	strings = append(strings, fmt.Sprintf("Port: %d", v.Port))
	strings = append(strings, fmt.Sprintf("Path: %s", v.Path))
	strings = append(strings, fmt.Sprintf("Timeout: %d", v.Timeout))

	return strings
}

func (v ServiceConfigurations) Url(args...string) (*url.URL,  error) {
	var s string = ""
	for _, v := range args {
		s += "/" + v
	}

	var u *url.URL = nil
	var err  error
	if len(s) == 0 {
		u = new(url.URL)
	} else {
		u, err = url.Parse(path.Clean(s))
		if err != nil {
			return nil, err
		}
	}

	if v.EnableTls {
		u.Scheme = "ssl"
	} else {
		u.Scheme = "tcp"
	}

	if v.Port == 0 {
		u.Host = v.Address
	} else {
		u.Host = fmt.Sprintf("%s:%d", v.Address, v.Port)
	}

	return u, nil
}

func (v HttpConfigurations) Url(args...string) (*url.URL,  error) {

	var s string = v.Path
	for _, v := range args {
		s += "/" + v
	}

	var u *url.URL = nil
	var err  error
	if len(s) == 0 {
		u = new(url.URL)
	} else {
		u, err = url.Parse(path.Clean(s))
		if err != nil {
			return nil, err
		}
	}

	if v.EnableTls {
		u.Scheme = "https"
	} else {
		u.Scheme = "http"
	}

	if v.Port == 0 {
		u.Host = v.Address
	} else {
		if v.EnableTls {
			if v.Port == 443 {
				u.Host = v.Address
			} else {
				u.Host = fmt.Sprintf("%s:%d", v.Address, v.Port)
			}
		} else {
			if v.Port == 80 {
				u.Host = v.Address
			} else {
				u.Host = fmt.Sprintf("%s:%d", v.Address, v.Port)
			}
		}
	}

	return u, nil
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

func (v *GatewayConfigurations) PrintConfigurations() {
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
func (v *ServiceConfigurations) PrintConfigurations() {
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
func (v *PublishConfigurations) PrintConfigurations() {
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

func GetFlagString(option *string, name string, param *string, deffunc func() string ) {
	if IsFlagPassed(name) {
		if param == nil {
			return
		}
		*option = *param
	} else if deffunc != nil {
		*option = deffunc()
	//} else {
	//	if len(*option) == 0 {
	//		if param != nil && 0 < len(*param) {
	//			*option = *param
	//		}
	//	}
	}
}

func GetFlagBoolean(option *bool, name string, param *bool, deffunc func() bool) {
	if IsFlagPassed(name) {
		if param == nil {
			return
		}
		*option = *param
	} else if deffunc != nil {
		*option = deffunc()
	}
}

func GetFlagInt(option *int, name string, param *int, deffunc func() int) {
	if IsFlagPassed(name) {
		if param == nil {
			return
		}
		*option = *param
	} else if deffunc != nil {
		*option = deffunc()
	}
}

func GetFlagInt32(option *int32, name string, param *int32, deffunc func() int32) {
	if IsFlagPassed(name) {
		if param == nil {
			return
		}
		*option = *param
	} else if deffunc != nil {
		*option = deffunc()
	}
}

func GetFlagInt64(option *int64, name string, param *int64, deffunc func() int64) {
	if IsFlagPassed(name) {
		if param == nil {
			return
		}
		*option = *param
	} else if deffunc != nil {
		*option = deffunc()
	}
}

func GetFlagUint(option *uint, name string, param *uint, deffunc func() uint) {
	if IsFlagPassed(name) {
		if param == nil {
			return
		}
		*option = *param
	} else if deffunc != nil {
		*option = deffunc()
	}
}

func GetFlagUint32(option *uint32, name string, param *uint32, deffunc func() uint32) {
	if IsFlagPassed(name) {
		if param == nil {
			return
		}
		*option = *param
	} else if deffunc != nil {
		*option = deffunc()
	}
}

func GetFlagUint64(option *uint64, name string, param *uint64, deffunc func() uint64) {
	if IsFlagPassed(name) {
		if param == nil {
			return
		}
		*option = *param
	} else if deffunc != nil {
		*option = deffunc()
	}
}
